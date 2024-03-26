package unionPay

import (
	"crypto"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	baseError "github.com/go-tron/base-error"
	"github.com/go-tron/crypto/encoding"
	"github.com/go-tron/crypto/rsaUtil"
	localTime "github.com/go-tron/local-time"
	"github.com/go-tron/logger"
	"github.com/go-tron/types/fieldUtil"
	"github.com/go-tron/types/mapUtil"
	"github.com/go-tron/union-pay/sdkConfig"
	"github.com/parnurzeal/gorequest"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

var (
	ErrorParam         = baseError.SystemFactory("3001", "支付参数错误:{}")
	ErrorMethod        = baseError.SystemFactory("3002", "支付方式无效:{}")
	ErrorVersion       = baseError.System("3003", "支付版本错误")
	ErrorSign          = baseError.System("3004", "支付签名失败")
	ErrorRequest       = baseError.System("3005", "支付服务连接失败")
	ErrorUnmarshalBody = baseError.System("3306", "支付解析失败")
	ErrorVerify        = baseError.System("3007", "支付验证失败")
	ErrorEncrypt       = baseError.System("3008", "支付加密失败")
	ErrorDecrypt       = baseError.System("3009", "支付解密失败")
	ErrorCode          = baseError.Factory("3010")
)

func New(config *UnionPay) *UnionPay {

	if config.Env == "" || config.MerId == "" || config.CertId == "" || config.PrivateKeyPem == "" || config.Logger == nil {
		panic("invalid unionPay config")
	}

	privateKey, err := rsaUtil.GetPrivateKeyPem([]byte(config.PrivateKeyPem))
	if err != nil {
		panic(err)
	}
	config.PrivateKey = privateKey

	if config.Env == "production" {
		config.SDKConfig = sdkConfig.Production
	} else {
		config.SDKConfig = sdkConfig.Testing
	}
	return config
}

type UnionPay struct {
	Env           string
	MerId         string
	CertId        string
	PrivateKeyPem string
	PrivateKey    *rsa.PrivateKey
	TrId          string
	IsEncrypt     bool
	SDKConfig     *sdkConfig.SDKConfig
	FrontUrl      string
	BackUrl       string
	Logger        logger.Logger
}

func (u *UnionPay) HashType(version string) (crypto.Hash, error) {
	switch version {
	case "5.0.0":
		return crypto.SHA1, nil
	case "5.1.0":
		return crypto.SHA256, nil
	default:
		return 0, ErrorVersion
	}
}

func (u *UnionPay) Sign(obj map[string]interface{}) error {
	signStr := mapUtil.ToSortString(obj)

	hashType, err := u.HashType(obj["version"].(string))
	if err != nil {
		return err
	}
	var hashMethod = hashType.New()
	hashMethod.Write([]byte(signStr))
	signHash := (&encoding.Hex{}).EncodeToString(hashMethod.Sum(nil))

	sign, err := rsaUtil.Sign(signHash, u.PrivateKey, hashType, &encoding.Base64{})
	if err != nil {
		return ErrorSign
	}
	obj["signature"] = sign
	return nil
}

func (u *UnionPay) Verify(obj map[string]interface{}) error {
	sign := obj["signature"]
	if sign == nil {
		return ErrorVerify
	}
	delete(obj, "signature")
	signStr := mapUtil.ToSortString(obj)

	hashType, err := u.HashType(obj["version"].(string))
	if err != nil {
		return err
	}
	var hashMethod = hashType.New()
	hashMethod.Write([]byte(signStr))
	signHash := (&encoding.Hex{}).EncodeToString(hashMethod.Sum(nil))

	publicKey := u.SDKConfig.SignPublicKey["default"]
	signPubKeyCert := obj["signPubKeyCert"]
	if signPubKeyCert != nil {
		cert := signPubKeyCert.(string)
		cert = strings.ReplaceAll(cert, "\r", "")
		cert = strings.ReplaceAll(cert, "\n", "")
		for key, val := range u.SDKConfig.SignCertificate {
			if val == cert {
				publicKey = u.SDKConfig.SignPublicKey[key]
			}
		}
	}

	if err := rsaUtil.Verify(signHash, sign.(string), publicKey, hashType, &encoding.Base64{}); err != nil {
		return ErrorVerify
	}
	return nil
}

func (u *UnionPay) Encrypt(text string) (string, error) {
	cipherText, err := rsaUtil.Encrypt(text, u.SDKConfig.EncryptPublicKey, &encoding.Base64{})
	if err != nil {
		return "", ErrorEncrypt
	} else {
		return cipherText, nil
	}
}

func (u *UnionPay) Decrypt(cipherText string) (string, error) {
	text, err := rsaUtil.Decrypt(cipherText, u.PrivateKey, &encoding.Base64{})
	if err != nil {
		return "", ErrorDecrypt
	} else {
		return text, nil
	}
}

func (u *UnionPay) GetUrl(name string) string {
	switch name {
	case "Query":
		return u.SDKConfig.SingleQueryPath
	case "Revoke":
		return u.SDKConfig.BackTransPath
	case "Refund":
		return u.SDKConfig.BackTransPath
	case "ContractConsume":
		return u.SDKConfig.BackTransPath
	case "App":
		return u.SDKConfig.AppRequestPath
	case "Web":
		return u.SDKConfig.FrontTransPath
	case "TokenOpen":
		return u.SDKConfig.FrontTransPath
	case "TokenCode":
		return u.SDKConfig.BackTransPath
	case "TokenConsume":
		return u.SDKConfig.BackTransPath
	case "TokenQuery":
		return u.SDKConfig.BackTransPath
	default:
		return ""
	}
}

func (u *UnionPay) Authorize(data map[string]interface{}) error {
	if u.IsEncrypt {
		data["encryptCertId"] = u.SDKConfig.EncryptCertId
		if !fieldUtil.IsEmpty(data["accNo"]) {
			accNo, err := u.Encrypt(data["accNo"].(string))
			if err != nil {
				return err
			}
			data["accNo"] = accNo
		}
	}
	if err := u.Sign(data); err != nil {
		return err
	}
	return nil
}

func (u *UnionPay) Execute(name string, data map[string]interface{}) (m map[string]interface{}, err error) {
	url := u.GetUrl(name)
	if url == "" {
		return nil, ErrorMethod(name)
	}
	if err := u.Authorize(data); err != nil {
		return nil, err
	}

	request, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	u.Logger.Info(string(request),
		u.Logger.Field("orderId", data["orderId"]),
		u.Logger.Field("name", name),
		u.Logger.Field("type", "request"),
	)

	if url == u.SDKConfig.FrontTransPath {
		m, err = u.Html(name, url, data)
	} else {
		m, err = u.Request(name, url, data)
	}

	return m, err
}

func (u *UnionPay) Request(name string, url string, data map[string]interface{}) (result map[string]interface{}, err error) {

	var (
		response = ""
		errs     []error
	)
	defer func() {
		u.Logger.Info(response,
			u.Logger.Field("orderId", data["orderId"]),
			u.Logger.Field("name", name),
			u.Logger.Field("type", "response"),
			u.Logger.Field("error", err))
	}()

	_, response, errs = gorequest.New().CustomMethod("POST", url).
		Type("form").
		Send(data).
		End()

	if errs != nil {
		return nil, ErrorRequest
	}

	res, err := UnmarshalBody(response)
	if err != nil {
		return nil, err
	}

	resStr, err := json.Marshal(res)
	if err == nil {
		response = string(resStr)
	}

	if err := u.Verify(res); err != nil {
		return nil, err
	}

	if res["respCode"] != "00" {
		return nil, ErrorCode(res["respMsg"])
	}

	return res, nil
}

func (u *UnionPay) Html(name string, url string, data map[string]interface{}) (map[string]interface{}, error) {

	html := "<html>\n" +
		"<head>\n" +
		"<meta http-equiv='Content-Type' content='text/html; charset=UTF-8'/>\n" +
		"</head>\n" +
		"<body>\n" +
		"<form id = 'form' action='" + url + "' method='post'>\n"

	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		html += "<input type='hidden' name='" + key + "' id='" + key + "' value='" + fmt.Sprint(data[key]) + "'/>\n"
	}

	html += "</form>\n" +
		"</body>\n" +
		"<script type='text/javascript'>document.getElementById('form').submit()</script>\n" +
		"</html>\n"

	return map[string]interface{}{
		"html": html,
	}, nil
}

func (u *UnionPay) CreateCustomerInfo(data map[string]interface{}) (string, error) {
	str := ""
	encryptedStr := ""
	for key, value := range data {
		if !fieldUtil.IsEmpty(value) {
			if u.IsEncrypt && (key == "phoneNo" || key == "cvn2" || key == "expired") {
				if encryptedStr != "" {
					encryptedStr += "&"
				}
				encryptedStr += key + "=" + fmt.Sprint(value)
			} else {
				if str != "" {
					str += "&"
				}
				str += key + "=" + fmt.Sprint(value)
			}
		}
	}
	if encryptedStr != "" {
		encryptedInfo, err := u.Encrypt(encryptedStr)
		if err != nil {
			return "", err
		}
		if str != "" {
			str += "&"
		}
		str += "encryptedInfo=" + encryptedInfo
	}
	str = "{" + str + "}"
	info := (&encoding.Base64{}).EncodeToString([]byte(str))
	return info, nil
}

func (u *UnionPay) GetCardNo(body string) (string, error) {
	unescapeBody, err := url.QueryUnescape(body)
	if err != nil {
		return "", ErrorDecrypt
	}
	return u.Decrypt(unescapeBody)
}

func (u *UnionPay) GetCustomerInfo(body string) (map[string]interface{}, error) {
	infoStr, err := (&encoding.Base64{}).DecodeString(body)
	if err != nil {
		return nil, ErrorDecrypt
	}
	if !u.IsEncrypt {
		return UnmarshalElement(string(infoStr))
	}

	data, err := UnmarshalElement(string(infoStr))
	if err != nil {
		return nil, ErrorDecrypt
	}
	info, err := u.Decrypt(data["encryptedInfo"].(string))
	if err != nil {
		return nil, err
	}
	return UnmarshalElement(info)
}

func GetReserved(body string) (map[string]interface{}, error) {
	infoStr, err := (&encoding.Base64{}).DecodeString(body)
	if err != nil {
		return nil, ErrorDecrypt
	}
	data, err := UnmarshalElement(string(infoStr))
	if err != nil {
		return nil, ErrorDecrypt
	}
	return data, nil
}

type TokenPayData struct {
	Token      string         `json:"token"`
	TokenBegin localTime.Time `json:"tokenBegin"`
	TokenEnd   localTime.Time `json:"tokenEnd"`
	TokenLevel string         `json:"tokenLevel"`
	TokenType  string         `json:"tokenType"`
	TrId       string         `json:"trId"`
}

func GetTokenPayData(body string) (*TokenPayData, error) {
	unescapeBody, err := url.QueryUnescape(body)
	if err != nil {
		return nil, ErrorDecrypt
	}
	m, err := UnmarshalElement(unescapeBody)
	if err != nil {
		return nil, err
	}

	tokenBegin, err := localTime.ParseCompact(m["tokenBegin"].(string))
	if err != nil {
		return nil, err
	}
	tokenEnd, err := localTime.ParseCompact(m["tokenEnd"].(string))
	if err != nil {
		return nil, err
	}
	tokenPayData := &TokenPayData{
		Token:      m["token"].(string),
		TokenBegin: tokenBegin,
		TokenEnd:   tokenEnd,
		TokenLevel: m["tokenLevel"].(string),
		TokenType:  m["tokenType"].(string),
		TrId:       m["trId"].(string),
	}

	return tokenPayData, nil
}

func UnmarshalElement(body string) (map[string]interface{}, error) {
	var m = map[string]interface{}{}
	body = strings.Replace(body, "{", "", 1)
	body = strings.Replace(body, "}", "", 1)
	bodyArr := strings.Split(body, "&")
	if len(bodyArr) > 0 {
		for _, temp := range bodyArr {
			index := strings.Index(temp, "=")
			if index == -1 {
				continue
			}
			key := temp[0:index]
			value := temp[index+1:]
			m[key] = value
		}
	}
	return m, nil
}

func UnmarshalBody(body string) (map[string]interface{}, error) {
	var m = map[string]interface{}{}
	if body == "" {
		return nil, ErrorUnmarshalBody
	}

	itemNum := 0
	var itemArr []string

	for {
		sIndex := strings.Index(body, "{")
		if sIndex == -1 {
			break
		}
		eIndex := strings.Index(body, "}")
		if eIndex == -1 {
			return nil, ErrorUnmarshalBody
		}

		item := body[sIndex : eIndex+1]
		itemArr = append(itemArr, item)
		body = strings.Replace(body, item, "<@|"+strconv.Itoa(itemNum)+"|@>", 1)
		itemNum++
	}

	bodyArr := strings.Split(body, "&")
	if len(bodyArr) > 0 {
		for _, temp := range bodyArr {
			index := strings.Index(temp, "=")
			if index == -1 {
				continue
			}
			key := temp[0:index]
			value := temp[index+1:]
			if strings.Index(value, "<@|") != -1 {
				vIndex := strings.Replace(value, "<@|", "", 1)
				vIndex = strings.Replace(vIndex, "|@>", "", 1)
				intIndex, err := strconv.Atoi(vIndex)
				if err == nil {
					value = itemArr[intIndex]
				}
			}
			m[key] = value
		}
	}

	return m, nil
}
