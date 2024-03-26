package unionPay

import (
	"fmt"
	"github.com/go-tron/local-time"
	"github.com/go-tron/types/fieldUtil"
	"strconv"
)

type RevokeReq struct {
	TransactionId            string  `json:"transactionId" validate:"required"`
	TxnAmount                float64 `json:"txnAmount" validate:"required"`
	OrigChannelTransactionId string  `json:"origChannelTransactionId" validate:"required"`
	Description              string  `json:"description" validate:"required"`
	BackUrl                  string  `json:"backUrl"`
}

func (u *UnionPay) Revoke(data *RevokeReq) (map[string]interface{}, error) {

	if fieldUtil.IsEmpty(data.TransactionId) {
		return nil, ErrorParam("订单号")
	}
	if fieldUtil.IsEmpty(data.TxnAmount) {
		return nil, ErrorParam("支付金额")
	}
	if fieldUtil.IsEmpty(data.OrigChannelTransactionId) {
		return nil, ErrorParam("原交易渠道单号")
	}
	if fieldUtil.IsEmpty(data.Description) {
		return nil, ErrorParam("订单详情")
	}

	txnAmt, err := strconv.ParseFloat(fmt.Sprintf("%.2f", data.TxnAmount*100), 64)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{
		"accessType":  "0",
		"backUrl":     u.BackUrl,
		"bizType":     "000201",
		"certId":      u.CertId,
		"channelType": "07",
		"encoding":    "utf-8",
		"merId":       u.MerId,
		"orderId":     data.TransactionId,
		"origQryId":   data.OrigChannelTransactionId,
		"reqReserved": data.Description,
		"signMethod":  "01",
		"txnAmt":      txnAmt,
		"txnSubType":  "00",
		"txnTime":     localTime.Now().Compact(),
		"txnType":     "31",
		"version":     "5.1.0",
	}
	if data.BackUrl != "" {
		params["backUrl"] = data.BackUrl
	}
	return u.Execute("Revoke", params)
}
