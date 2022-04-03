package dto

type AccountRequest struct {
	StatusID string `json:"delivery_status_id"`
}

type AccountResponse struct {
	StatusID string `json:"delivery_status_id"`
}
