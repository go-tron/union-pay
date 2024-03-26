package unionPay

import (
	"fmt"
	"github.com/go-tron/local-time"
	"github.com/go-tron/types/fieldUtil"
	"strconv"
)

type WebReq struct {
	TransactionId string  `json:"transactionId" validate:"required"`
	TxnAmount     float64 `json:"txnAmount" validate:"required"`
	Card          string  `json:"card" validate:"required"`
	Description   string  `json:"description" validate:"required"`
	FrontUrl      string  `json:"frontUrl"`
	BackUrl       string  `json:"backUrl"`
}

func (u *UnionPay) Web(data *WebReq) (map[string]interface{}, error) {

	if fieldUtil.IsEmpty(data.TransactionId) {
		return nil, ErrorParam("订单号")
	}
	if fieldUtil.IsEmpty(data.TxnAmount) {
		return nil, ErrorParam("支付金额")
	}
	if fieldUtil.IsEmpty(data.Card) {
		return nil, ErrorParam("银行卡号")
	}
	if fieldUtil.IsEmpty(data.Description) {
		return nil, ErrorParam("订单详情")
	}

	txnAmt, err := strconv.ParseFloat(fmt.Sprintf("%.2f", data.TxnAmount*100), 64)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"accNo":        data.Card,
		"accType":      "01",
		"accessType":   "0",
		"backUrl":      u.BackUrl,
		"bizType":      "000201",
		"certId":       u.CertId,
		"channelType":  "07",
		"currencyCode": "156",
		"encoding":     "UTF-8",
		"frontUrl":     u.FrontUrl,
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
	return u.Execute("Web", params)
}
