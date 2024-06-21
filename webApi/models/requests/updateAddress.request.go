package requests

type UpdateAddressRequest struct {
	Country      *string `json:"country"`
	State        *string `json:"state"`
	City         *string `json:"city"`
	Neighborhood *string `json:"neighborhood"`
	Street       *string `json:"street"`
	Number       *string `json:"number"`
	Complement   *string `json:"complement"`
}
