package dto

type AccountRequest struct {
	AccountId string `json:"account_id"`
}

type AccountResponse struct {
	AccountId string `json:"account_id"`
	Name      string `json:"account_name"`
	Balance   string `json:"balance"`
}
