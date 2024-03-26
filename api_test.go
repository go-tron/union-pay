package unionPay

import (
	"fmt"
	"github.com/go-tron/logger"
	"testing"
)

const (
	frontUrl = "http://47.96.113.100/front"
	backUrl  = "http://47.96.113.100/back"
)

//var api = unionPay.New(&unionPay.UnionPay{
//	Env:           "production",
//	MerId:         "898610153111077",
//	CertId:        "75565791313",
//	PrivateKeyPem: "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQDON43xbiFuQ8qx\n17SGkjJ73vn4cLvq8Zp7CkXeS1pJ5hltMRK6tfBA4t1iOcrzFqSyZrg5Wt+sZFxq\nx0yc09Tv+kCtXKcJozGKwNbcHDmhnOwjJ6a+09S6KtfhwrUPpvZC4dL9/S4rt15y\nex7HnyxmICkL+gwRODuNNEy4wrRcpHkO9201PrbSONE+KSHt7u5RGGHLknCI2Al1\nUB6uLEXFiDOPTcXv1WP05pyt12W7Tym+eBk7W2zD/jq6PpDXsHTTnGwcXfvY8QNU\nRFJZoBr/J4LitKBfb4TfVmwmb9P18Dp6sqC0cZVddHdxv+521x9VXa25d+YYUPlC\nvKHUyN5hAgMBAAECggEBAJnQYESLo83uLNs4DIQvug4hBAymcdRNf690uyaSx/bE\n9YQ3jUPM/zY2wXBJsTJeWlxjN9g4Cww4E6LloyR68KJK5EErEcCjhAl9ywWE3AvK\n02tpYPuHX225FStkos7Y92htroYaeXzMGluDsDKvbFuEc3lmeefCLNEnCG7rc8KA\n6Iybg0p5/k2LqqpKWNS8rO3B9WyqSc/8BdfQaQbgAjpMmeUoPrW/8Xg+tfZn9kE/\nIE9W1nWleOSDaAPy2UW9e/fnHmHJScSRryrhqIvGKKSeY2t9O+FemDODm3euRSj7\n924Wo/3FdoecphhiUQOlopAEPhvT1tLjncv07P+L7n0CgYEA7M1VeA+Q4RVei6JS\nmmIB10ZJX9+nbJ1I9T8JcOgOyulHwV1jOGmNZOszM7QRzEJZ8d16TqRIWG+Im0p/\nFihs2C0Pxoqsdnjv3XsQ8aN3egpXXID0kv9nhjgaZqM2jL7keGWHzrYFyYTao6ax\nnHKyUPCV3dTHTukMg3k9M1xWkmsCgYEA3u9yux24+aPEgCiXt2F/o3mmc4Y/nydZ\nvF9tVCLQpGdg7h0v/adi1wVidJyVzveg2O4CgB9uNfTstZOgmHQv1VQbsZqujdAE\nJRCurOL5JIl1s5vHJX0CoQqaKW5/QXf9Pody1snsTihoESJJUXoaimX7Fjskc06y\nkyXejSFwfWMCgYEA3LviqdHIuSeURBEBcstc7/CiwO3YQRS0R3eFM2v9LgQaURLe\nRF2P4byks9Nsq7xF34EWW2wjy9vbuPXOVTdFjpUimW32qVC+NvRarQQ1gNtW5/GF\nCelt9N0jLv2BAd2HVvs9huCBsIFY1SezdCQuN+irqmaHivUOJxfChrfSMxUCgYEA\noYLopBeMg4/aNN8zy2+fgVgI93kzbqKbqJWaJ3JxpiofV2Nd+W7jdNQ52DH5dtCA\nf6kvtpMOzZh8RbCGkOzcGrmstqdR/vvNhOzpH+fwiD/uPgmF5esDlJiRx6J+H3hG\nyJ5o3KO9x4IOoxsr8xN7Vvk7R4fhiIhjZ6fv7F9ll/ECgYAsBsPXHBGRfvTh+QxC\nxtPnjnbVabh96JzKCFc/zsqkl8gRx8iJ460SoH2y6dpxyh1UV+82Mv09k+C/l0yO\nNQVOotGkOG29e1dujt3Y34Qa7evlYZA0om/DOv/asuJyMy4AeRQIQ8ITsmd2B8WP\nBs/kDSj2Fop+Il49Lu9QCjRmAg==\n-----END PRIVATE KEY-----",
//	TrId:          "",
//	IsEncrypt:     true,
//	FrontUrl:      frontUrl,
//	BackUrl:       backUrl,
//	Logger:        zapLogger.New("unionPay", "info"),
//})

var api = New(&UnionPay{
	Env:           "production",
	MerId:         "802440355410544",
	CertId:        "79782041625",
	PrivateKeyPem: "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCzy3GCckhMdKdV\n6MQ0wf8p94sFe9vMJR173SWQjTRyHnXYGhK+oowjaJKx5cky0XLi8mNGHYpw/oN0\nGwLkQkV1gWh7RnpOOC2RtzwCEmmNewufSFZljlVd4ZfD+BeA5LbTA5fwkRlXaDdA\nuM6VL99WPayL/M0FROczNuCjFyFRnF8ujTlYbw4+0lH5lQ3PqVL8fRdGH1j0fLuE\nVnJY8pN5eR1JLms63LsU+Hb4TgLJjmiWnyRCCwhAJyiW6g/venBTVic95Z9u3+Hp\nzesRUKs5Gn8dbMvYDNP3e8oGuG6yNERePAN3lJfZbY8K89ZKfcOlShos1gN0mm+E\nOPW2KnmBAgMBAAECggEAdppd5o884jKRsDo8NBFFIYoKSzPJaoGRUYQwo4qCbkVp\nt4R7mQXhK/pvyqqqbrt77fm3Qyl5idBbJtwqrCiLaVhUc/2p8K9eYCFflH5Q4uOo\nZwpw881Qwv1t//hIiN5XPBvCvw74iJBW4nsGy6Mo0NseV7oMBM3PNhmdVfiMIijH\ncD6mPH/YzGzPJGQ2GmJPHidi9BOoT2ezhVDuJ6HN8mi39cmrg+WJS0pHy5dPlFzf\nNk7OpOC5JtTweDQ5FaYENnPgyyeHb2YiKpXcryORvrdB20IsdjDajt7iBzMW9Nj0\nKdsso217kPtZtErw2WvTFO75gCdjQQsQCnGGlgQjgQKBgQD16Z0TAoM0ZPAr6mxz\nibxhImtYhh/J+jsxPwPIVVbzhDZ+UdroELwlrdbk4vHl0A7tdhzksHXW9mEhsydR\n8qmo3PYLLMccuew1RQliVQ5BPPXfU1vt8X5QjTD0Gm7HO8i2g4ZhlFdcbr65cg0l\nE6CmguOVZiH2SgPHjBOOjscj7wKBgQC7K4KuausoCZn6wPlz0K04Vs86Uuh2U2kg\nROrl7854PC87TblCqGpoA+mO7A5zgIez3DWqUvXX/nZ+MKpimGQ893bmcnemYnGU\nq0Hu1ypX+EzK4htutzNFEyqFu6dF1B+D+/gCiRO9aluV1hQjo9b54ofIhhUMvMoY\nlozDY/YJjwKBgQDoWRzCc2NIw/5xYAOsfv1wwbavY7rxee8nrqSCSq/nUoqye+Db\n18QNzdz3Ur7AJv+Tuj++jEQrIYvQ7Zc/RIqtNGl9UYeoSrs18c7WDAt19IVTx01a\n2wxAS8dvPPJiaMhSOp9j32dTvIeUbICjAZKPAajK5j88l95OUrm/voPGHQKBgQCH\nHhb98OOICL5m/g5W5EtYQ5rf4OA+I2Ldz8K/cYnc1J3IMwjNrST7p+9gpNqd2Sid\n2BfJuJAPeHx0Bo/KE3cxZ/gmznW/4ItcJvG3CQ9haDePswDYwVo6wGZPGrPbvhFu\nR2S399sPP4uUNHJfdDSKOlWJlfn2Mwe2DmDZ4PzJ0wKBgQCfL5KCocyZ3xo44+cL\n/UykiSScSuGy1xUp3/pJUR2+fNJc4CrLp2MSMQZg5YYAy8H0vWQyKwJO7wWX+Yhq\nsD5CBWwA/CaJ/WPi3EXXykps2XIDo54vNTsWbpZ57MhTvhl+XayqOBBzzabSncW8\nNS/ER7hdTSjBGXhd5tDnEQuJPQ==\n-----END PRIVATE KEY-----",
	TrId:          "62000004446",
	IsEncrypt:     true,
	FrontUrl:      frontUrl,
	BackUrl:       backUrl,
	Logger:        logger.NewZap("unionPay", "info"),
})

func TestContractConsume(t *testing.T) {
	result, err := api.ContractConsume(&ContractConsumeReq{
		TransactionId: "20201021001",
		TxnAmount:     0.01,
		ContractId:    "6218400000084676152",
		Description:   "测试",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestRevoke(t *testing.T) {
	result, err := api.Revoke(&RevokeReq{
		TransactionId:            "12854923295488778241",
		TxnAmount:                0.01,
		OrigChannelTransactionId: "752007211630292867678",
		Description:              "撤销",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestRefund(t *testing.T) {
	result, err := api.Refund(&RefundReq{
		TransactionId:            "20200122006",
		TxnAmount:                0.01,
		OrigChannelTransactionId: "952003011613043235058",
		Description:              "撤销",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestQuery(t *testing.T) {
	result, err := api.Query(&QueryReq{
		TransactionId: "1285492329548877824",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestTokenQuery(t *testing.T) {
	result, err := api.TokenQuery(&TokenQueryReq{
		TransactionId: "1285551248044462080",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestTokenCode(t *testing.T) {
	result, err := api.TokenCode(&TokenCodeReq{
		TransactionId: "20200722001",
		TxnAmount:     0.01,
		Card:          "",
		Token:         "6235240000760877254",
		Phone:         "13002937208",
		Description:   "cs",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestTokenConsume(t *testing.T) {
	result, err := api.TokenConsume(&TokenConsumeReq{
		TransactionId: "20200722001",
		TxnAmount:     0.01,
		Token:         "6235240000760877254",
		Code:          "",
		Description:   "cs",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestGetCustomerInfo(t *testing.T) {
	a := "accNo=Wu4BGLH4HLwkTE55FaT%2FE4lhrb14KMqTRDY4X9hI1CxjMhLCxP6ldGiqw5uRWSemlorPht9WI4Mg7zAxXDcLS2%2Bn%2FDjqjXZg947VA9k3YnCdHAOtugHP%2BbljsrU0VbDpaW7KruJwYJdSrROhAWx1vnerWFSU3DgPJxhjb3im%2BesQKzQU3eMQRCxYHFkZ4wrJdQnotPaEWYyZ9t6j02YwSf1vWFIueMb6QONqjkftL5mzhk%2F6zxV1Y43eC8nz7Am400QGJcbX%2B73ASaOgaZ%2Fcaj2yO5pem1qmi9nH8fl5ycRZK0wMv4AMj0A80ngl%2FPaPW59N8fc3jar%2Fl2qMtQAdlg%3D%3D&accessType=0&activateStatus=1&bizType=000902&customerInfo=e2VuY3J5cHRlZEluZm89cG5yakhWZGkxSllsZWtUOVdzVnd2elJKMG1MeXBNTFBQU3RQcCtOMlROWFlPL25peG5jT0xTdkFVTTBQS1Back5XcmpXSTRUR2U0MSszdE5FdHZLakx5YkxsR3V5Sng4QXFGT1B4QXluN0R2bFM2VlppTk5PdXJ6SUtJSTdiN3lGQTFhcDJtSzZ4Mnc5WUxVMEpkZzJtOW5uWUJwNUhLc3N2NWRKdjVzc3BvdHROMzE5SEd0VlR5b2tLQ01OeWtGcnErTXh5RWRHSTA0OHNpd1BybkdDbnBPdTdnYjhSUHNjenU1c0RCZjhsazZJcm1LU0E1cjlENTQ3aFA4TmlYc0FCVi8veGFlNWZnZWRPYjFVQ3ZEdzFaS05wdGVLUHpSSkQzeVR1dnBVRGora2NnRzR5eXRONzVnWUZYVElwUkI4Q3ZPRWl1b0ZCNE15QjNxOGRTU0JRPT19&encoding=UTF-8&merId=898450155411621&orderId=1285551248044462080&reqReserved=%E7%BB%91%E5%AE%9A%E9%93%B6%E8%A1%8C%E5%8D%A1&respCode=00&respMsg=success&signMethod=01&signPubKeyCert=-----BEGIN+CERTIFICATE-----%0D%0AMIIEIDCCAwigAwIBAgIFEDRVM3AwDQYJKoZIhvcNAQEFBQAwITELMAkGA1UEBhMC%0D%0AQ04xEjAQBgNVBAoTCUNGQ0EgT0NBMTAeFw0xNTEwMjcwOTA2MjlaFw0yMDEwMjIw%0D%0AOTU4MjJaMIGWMQswCQYDVQQGEwJjbjESMBAGA1UEChMJQ0ZDQSBPQ0ExMRYwFAYD%0D%0AVQQLEw1Mb2NhbCBSQSBPQ0ExMRQwEgYDVQQLEwtFbnRlcnByaXNlczFFMEMGA1UE%0D%0AAww8MDQxQDgzMTAwMDAwMDAwODMwNDBA5Lit5Zu96ZO26IGU6IKh5Lu95pyJ6ZmQ%0D%0A5YWs5Y%2B4QDAwMDE2NDkzMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA%0D%0AtXclo3H4pB%2BWi4wSd0DGwnyZWni7%2B22Tkk6lbXQErMNHPk84c8DnjT8CW8jIfv3z%0D%0Ad5NBpvG3O3jQ%2FYHFlad39DdgUvqDd0WY8%2FC4Lf2xyo0%2BgQRZckMKEAId8Fl6%2FrPN%0D%0AHsbPRGNIZgE6AByvCRbriiFNFtuXzP4ogG7vilqBckGWfAYaJ5zJpaGlMBOW1Ti3%0D%0AMVjKg5x8t1%2FoFBkpFVsBnAeSGPJYrBn0irfnXDhOz7hcIWPbNDoq2bJ9VwbkKhJq%0D%0AVz7j7116pziUcLSFJasnWMnp8CrISj52cXzS%2FY1kuaIMPP%2F1B0pcjVqMNJjowooD%0D%0AOxID3TZGfk5V7S%2B%2B4FowVwIDAQABo4HoMIHlMB8GA1UdIwQYMBaAFNHb6YiC5d0a%0D%0Aj0yqAIy%2BfPKrG%2FbZMEgGA1UdIARBMD8wPQYIYIEchu8qAQEwMTAvBggrBgEFBQcC%0D%0AARYjaHR0cDovL3d3dy5jZmNhLmNvbS5jbi91cy91cy0xNC5odG0wNwYDVR0fBDAw%0D%0ALjAsoCqgKIYmaHR0cDovL2NybC5jZmNhLmNvbS5jbi9SU0EvY3JsMjI3Mi5jcmww%0D%0ACwYDVR0PBAQDAgPoMB0GA1UdDgQWBBTEIzenf3VR6CZRS61ARrWMto0GODATBgNV%0D%0AHSUEDDAKBggrBgEFBQcDAjANBgkqhkiG9w0BAQUFAAOCAQEAHMgTi%2B4Y9g0yvsUA%0D%0Ap7MkdnPtWLS6XwL3IQuXoPInmBSbg2NP8jNhlq8tGL%2FWJXjycme%2F8BKu%2BHht6lgN%0D%0AZhv9STnA59UFo9vxwSQy88bbyui5fKXVliZEiTUhjKM6SOod2Pnp5oWMVjLxujkk%0D%0AWKjSakPvV6N6H66xhJSCk%2BRef59HuFZY4%2FLqyZysiMua4qyYfEfdKk5h27%2Bz1MWy%0D%0AnadnxA5QexHHck9Y4ZyisbUubW7wTaaWFd%2BcZ3P%2FzmIUskE%2FdAG0%2FHEvmOR6CGlM%0D%0A55BFCVmJEufHtike3shu7lZGVm2adKNFFTqLoEFkfBO6Y%2FN6ViraBilcXjmWBJNE%0D%0AMFF%2FyA%3D%3D%0D%0A-----END+CERTIFICATE-----&signature=EdTy7DKasNojtTU59zYyn7n3e28vKiZZ9mIizrYD2N5OyaCwsLWoMrMG%2FbUKNgqm5gnH%2Bk4IIoxx90dg3UyVcZJVDBFs1u8irIslTJJmg0sny9Eilbunx7oYAB9uvOJ233E2L%2BKo1IGvWnR7uk4bKQ13mUoyfFt33wFJEIVnsuQhZG1ShP%2BOYxfOIKU%2F0gI1hW7NfYaPNneTuR9ZRcxo1Kg0%2FpLSixW%2FzMcUc30huy1fU5Pug89fVA6Ue8pXjMiU7eEs12%2BAsMehmXQpYj7Bb5fWZrOqBg3mcqPLvosigrX3SGhL9a8HqzUkHZtOMBnbV7qrT2Km5uz%2Frc1SWgZgXw%3D%3D&tokenPayData=%7Btoken%3D6235240000760877254%26trId%3D62000000606%26tokenType%3D01%26tokenBegin%3D20200721202459%26tokenEnd%3D20250720202459%26tokenLevel%3D10%7D&txnSubType=00&txnTime=20200721202436&txnType=79&version=5.1.0"
	body, err := UnmarshalBody(a)
	fmt.Println("body err", err)
	fmt.Println("body", body)

	fmt.Println("orderId", body["orderId"])

	accNo, err := api.GetCardNo(body["accNo"].(string))
	fmt.Println("accNo err", err)
	fmt.Println("accNo", accNo)

	tokenPayData, err := GetTokenPayData(body["tokenPayData"].(string))
	fmt.Println("tokenPayData err", err)
	fmt.Println("tokenPayData", tokenPayData)

	customerInfo, err := api.GetCustomerInfo(body["customerInfo"].(string))
	fmt.Println("customerInfo err", err)
	fmt.Println("customerInfo", customerInfo)
}
