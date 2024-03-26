package unionPay

import (
	"github.com/go-tron/local-time"
	"github.com/go-tron/types/fieldUtil"
)

type QueryReq struct {
	TransactionId string `json:"transactionId" validate:"required"`
}

func (u *UnionPay) Query(data *QueryReq) (map[string]interface{}, error) {

	if fieldUtil.IsEmpty(data.TransactionId) {
		return nil, ErrorParam("订单号")
	}
	params := map[string]interface{}{
		"accessType":  "0",
		"bizType":     "000000",
		"certId":      u.CertId,
		"channelType": "07",
		"encoding":    "utf-8",
		"merId":       u.MerId,
		"orderId":     data.TransactionId,
		"signMethod":  "01",
		"txnSubType":  "00",
		"txnTime":     localTime.Now().Compact(),
		"txnType":     "00",
		"version":     "5.1.0",
	}

	return u.Execute("Query", params)
}
