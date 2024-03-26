package unionPay

type IMerchants interface {
	GetMerchantById(string) (*UnionPay, error)
}

type Merchants struct {
	Merchants IMerchants
}

func (u *Merchants) Verify(merId string, data map[string]interface{}) error {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return err
	}
	return merchant.Verify(data)
}
func (u *Merchants) Query(merId string, data *QueryReq) (map[string]interface{}, error) {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return nil, err
	}
	return merchant.Query(data)
}
func (u *Merchants) Refund(merId string, data *RefundReq) (map[string]interface{}, error) {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return nil, err
	}
	return merchant.Refund(data)
}
func (u *Merchants) Revoke(merId string, data *RevokeReq) (map[string]interface{}, error) {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return nil, err
	}
	return merchant.Revoke(data)
}
func (u *Merchants) App(merId string, data *AppReq) (map[string]interface{}, error) {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return nil, err
	}
	return merchant.App(data)
}
func (u *Merchants) Web(merId string, data *WebReq) (map[string]interface{}, error) {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return nil, err
	}
	return merchant.Web(data)
}
func (u *Merchants) TokenOpen(merId string, data *TokenOpenReq) (map[string]interface{}, error) {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return nil, err
	}
	return merchant.TokenOpen(data)
}
func (u *Merchants) TokenCode(merId string, data *TokenCodeReq) (map[string]interface{}, error) {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return nil, err
	}
	return merchant.TokenCode(data)
}
func (u *Merchants) TokenConsume(merId string, data *TokenConsumeReq) (map[string]interface{}, error) {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return nil, err
	}
	return merchant.TokenConsume(data)
}
func (u *Merchants) TokenQuery(merId string, data *TokenQueryReq) (map[string]interface{}, error) {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return nil, err
	}
	return merchant.TokenQuery(data)
}
func (u *Merchants) ContractConsume(merId string, data *ContractConsumeReq) (map[string]interface{}, error) {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return nil, err
	}
	return merchant.ContractConsume(data)
}
