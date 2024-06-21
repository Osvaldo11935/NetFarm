package requests

type GetPaymentReferenceRequest struct {
	Amount      float64 `json:"amount"`
	PhoneNumber string  `json:"phoneNumber"`
}
