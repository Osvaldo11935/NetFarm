package iexternalServices

import "time"

type IPaymentService interface {
	SimulatePayment(reference string, expirationTime time.Time) bool
	GetPaymentReference(amount float64, phoneNumber string, expirationMinutes int) (string, time.Time)
}
