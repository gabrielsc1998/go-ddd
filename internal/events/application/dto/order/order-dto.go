package order_dto

type OrderCreateDto struct {
	EventId    string `json:"eventId"`
	SectionId  string `json:"sectionId"`
	SpotId     string `json:"spotId"`
	CustomerId string `json:"customerId"`
	CardToken  string `json:"cardToken"`
}
