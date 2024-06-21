package responses

import "time"

type GetPaymentReferenceResponse struct {
	Reference      string    `json:"reference"`
	ExpirationTime time.Time `json:"expirationTime"`
}
