package unionPay

import (
	"fmt"
	"github.com/go-tron/local-time"
	"github.com/go-tron/types/fieldUtil"
	"strconv"
)

type TokenOpenReq struct {
	TransactionId string `json:"transactionId" validate:"required"`
	Card          string `json:"card" validate:"required"`
	Phone         string `json:"phone" validate:"required"`
	CertifTp      string `json:"certifTp" validate:"required"`
	CertifId      string `json:"certifId" validate:"required"`
	CustomerNm    string `json:"customerNm" validate:"required"`
	Description   string `json:"description" validate:"required"`
	FrontUrl      string `json:"frontUrl"`
	BackUrl       string `json:"backUrl"`
}

func (u *UnionPay) TokenOpen(data *TokenOpenReq) (map[string]interface{}, error) {
	if fieldUtil.IsEmpty(data.TransactionId) {
		return nil, ErrorParam("订单号")
	}
	if fieldUtil.IsEmpty(data.Card) {
		return nil, ErrorParam("银行卡号")
	}
	if fieldUtil.IsEmpty(data.Description) {
		return nil, ErrorParam("订单详情")
	}

	params := map[string]interface{}{
		"accNo":       data.Card,
		"accType":     "01",
		"accessType":  "0",
		"backUrl":     u.BackUrl,
		"bizType":     "000301",
		"certId":      u.CertId,
		"channelType": "07",
		"encoding":    "UTF-8",
		"frontUrl":    u.FrontUrl,
		"merId":       u.MerId,
		"orderId":     data.TransactionId,
		"reqReserved": data.Description,
		"signMethod":  "01",
		"txnSubType":  "00",
		"txnTime":     localTime.Now().Compact(),
		"txnType":     "79",
		"version":     "5.1.0",
	}

	if data.FrontUrl != "" {
		params["frontUrl"] = data.FrontUrl
	}
	if data.BackUrl != "" {
		params["backUrl"] = data.BackUrl
	}

	customerInfo, err := u.CreateCustomerInfo(map[string]interface{}{
		"phoneNo":    data.Phone,
		"certifTp":   data.CertifTp,
		"certifId":   data.CertifId,
		"customerNm": data.CustomerNm,
	})
	if err != nil {
		return nil, err
	}
	params["customerInfo"] = customerInfo

	if u.TrId != "" {
		params["bizType"] = "000902"
		params["tokenPayData"] = "{trId=" + u.TrId + "&tokenType=01}"
	}

	return u.Execute("TokenOpen", params)
}

type TokenCodeReq struct {
	TransactionId string  `json:"transactionId" validate:"required"`
	TxnAmount     float64 `json:"txnAmount" validate:"required"`
	Card          string  `json:"card"`
	Token         string  `json:"token"`
	Phone         string  `json:"phone" validate:"required"`
	Description   string  `json:"description" validate:"required"`
	FrontUrl      string  `json:"frontUrl"`
	BackUrl       string  `json:"backUrl"`
}

func (u *UnionPay) TokenCode(data *TokenCodeReq) (map[string]interface{}, error) {

	if fieldUtil.IsEmpty(data.TransactionId) {
		return nil, ErrorParam("订单号")
	}
	if fieldUtil.IsEmpty(data.TxnAmount) {
		return nil, ErrorParam("支付金额")
	}
	if fieldUtil.IsEmpty(data.Phone) {
		return nil, ErrorParam("手机号")
	}
	if fieldUtil.IsEmpty(data.Description) {
		return nil, ErrorParam("订单详情")
	}

	txnAmt, err := strconv.ParseFloat(fmt.Sprintf("%.2f", data.TxnAmount*100), 64)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"accessType":   "0",
		"backUrl":      u.BackUrl,
		"bizType":      "000301",
		"certId":       u.CertId,
		"channelType":  "07",
		"currencyCode": "156",
		"encoding":     "UTF-8",
		"merId":        u.MerId,
		"orderId":      data.TransactionId,
		"reqReserved":  data.Description,
		"signMethod":   "01",
		"txnAmt":       txnAmt,
		"txnSubType":   "02",
		"txnTime":      localTime.Now().Compact(),
		"txnType":      "77",
		"version":      "5.1.0",
	}

	if data.FrontUrl != "" {
		params["frontUrl"] = data.FrontUrl
	}
	if data.BackUrl != "" {
		params["backUrl"] = data.BackUrl
	}

	customerInfo, err := u.CreateCustomerInfo(map[string]interface{}{
		"phoneNo": data.Phone,
	})
	if err != nil {
		return nil, err
	}
	params["customerInfo"] = customerInfo

	if u.TrId != "" {
		if fieldUtil.IsEmpty(data.Token) {
			return nil, ErrorParam("银行卡标识")
		}
		params["bizType"] = "000902"
		params["tokenPayData"] = "{token=" + data.Token + "&trId=" + u.TrId + "}"
	} else {
		if fieldUtil.IsEmpty(data.Card) {
			return nil, ErrorParam("银行卡号")
		}
		params["accNo"] = data.Card
	}

	return u.Execute("TokenCode", params)
}

type TokenConsumeReq struct {
	TransactionId string  `json:"transactionId" validate:"required"`
	TxnAmount     float64 `json:"txnAmount" validate:"required"`
	Card          string  `json:"card"`
	Token         string  `json:"token"`
	Code          string  `json:"code"`
	Description   string  `json:"description" validate:"required"`
	FrontUrl      string  `json:"frontUrl"`
	BackUrl       string  `json:"backUrl"`
}

func (u *UnionPay) TokenConsume(data *TokenConsumeReq) (map[string]interface{}, error) {

	if fieldUtil.IsEmpty(data.TransactionId) {
		return nil, ErrorParam("订单号")
	}
	if fieldUtil.IsEmpty(data.TxnAmount) {
		return nil, ErrorParam("支付金额")
	}
	if fieldUtil.IsEmpty(data.Description) {
		return nil, ErrorParam("订单详情")
	}

	txnAmt, err := strconv.ParseFloat(fmt.Sprintf("%.2f", data.TxnAmount*100), 64)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"accessType":   "0",
		"backUrl":      u.BackUrl,
		"bizType":      "000301",
		"certId":       u.CertId,
		"channelType":  "07",
		"currencyCode": "156",
		"encoding":     "UTF-8",
		"merId":        u.MerId,
		"orderId":      data.TransactionId,
		"reqReserved":  data.Description,
		"signMethod":   "01",
		"txnAmt":       txnAmt,
		"txnSubType":   "01",
		"txnTime":      localTime.Now().Compact(),
		"txnType":      "01",
		"version":      "5.1.0",
	}

	if data.FrontUrl != "" {
		params["frontUrl"] = data.FrontUrl
	}
	if data.BackUrl != "" {
		params["backUrl"] = data.BackUrl
	}

	customerInfo, err := u.CreateCustomerInfo(map[string]interface{}{
		"smsCode": data.Code,
	})
	if err != nil {
		return nil, err
	}
	params["customerInfo"] = customerInfo

	if u.TrId != "" {
		if fieldUtil.IsEmpty(data.Token) {
			return nil, ErrorParam("银行卡标识")
		}
		params["bizType"] = "000902"
		params["tokenPayData"] = "{token=" + data.Token + "&trId=" + u.TrId + "}"
	} else {
		if fieldUtil.IsEmpty(data.Card) {
			return nil, ErrorParam("银行卡号")
		}
		params["accNo"] = data.Card
	}

	return u.Execute("TokenConsume", params)
}

type TokenQueryReq struct {
	TransactionId string `json:"transactionId" validate:"required"`
}

func (u *UnionPay) TokenQuery(data *TokenQueryReq) (map[string]interface{}, error) {

	if fieldUtil.IsEmpty(data.TransactionId) {
		return nil, ErrorParam("订单号")
	}

	params := map[string]interface{}{
		"accessType":  "0",
		"backUrl":     u.BackUrl,
		"bizType":     "000301",
		"certId":      u.CertId,
		"channelType": "07",
		"encoding":    "UTF-8",
		"merId":       u.MerId,
		"orderId":     data.TransactionId,
		"signMethod":  "01",
		"txnSubType":  "02",
		"txnTime":     localTime.Now().Compact(),
		"txnType":     "78",
		"version":     "5.1.0",
	}

	if u.TrId != "" {
		params["bizType"] = "000902"
	}

	return u.Execute("TokenQuery", params)
}
