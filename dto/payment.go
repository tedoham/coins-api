package dto

type PaymentRequest struct {
	StatusID string `json:"delivery_status_id"`
}

type PaymentResponse struct {
	StatusID string `json:"delivery_status_id"`
}
