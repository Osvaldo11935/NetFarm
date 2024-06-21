package externalService

import (
	"NetFarm/application/interfaces/iinfrastructure/iexternalServices"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type PaymentService struct {
}

func NewPaymentService() iexternalServices.IPaymentService {
	return &PaymentService{}
}
func (s *PaymentService) SimulatePayment(reference string, expirationTime time.Time) bool {
	if len(reference) != 9 {
		fmt.Println("Referência inválida. A referência deve ter nove dígitos.")
		return false
	}

	if time.Now().After(expirationTime) {
		fmt.Println("A referência expirou. O pagamento não pode ser processado.")
		return false
	}

	fmt.Println("Pagamento com referência", reference, "foi processado com sucesso.")

	return true
}
func (s *PaymentService) GetPaymentReference(amount float64, phoneNumber string, expirationMinutes int) (string, time.Time) {

	rand.Seed(time.Now().UnixNano())

	amountStr := strconv.FormatFloat(amount, 'f', 0, 64)

	phoneNumber = strings.TrimPrefix(phoneNumber, "+244")

	phoneNumber = strings.ReplaceAll(phoneNumber, "-", "")
	phoneNumber = strings.ReplaceAll(phoneNumber, " ", "")

	reference := fmt.Sprintf("%s%s", amountStr, phoneNumber)

	referenceBytes := []byte(reference)
	rand.Shuffle(len(referenceBytes), func(i, j int) {
		referenceBytes[i], referenceBytes[j] = referenceBytes[j], referenceBytes[i]
	})

	reference = string(referenceBytes)

	reference = reference[:9]

	expirationTime := time.Now().Add(time.Duration(expirationMinutes) * time.Minute)

	return reference, expirationTime
}
