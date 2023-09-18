package event_controller

type CreateEventInputDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	PartnerId   string `json:"partner_id"`
}

type AddSectionInputDto struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	TotalSpots  int     `json:"total_spots"`
	Price       float64 `json:"price"`
}

type UpdateSectionInputDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateLocationInputDto struct {
	Location string `json:"location"`
}

type CreateOrderInputDto struct {
	SectionId  string `json:"section_id"`
	SpotId     string `json:"spot_id"`
	CustomerId string `json:"customer_id"`
	CardToken  string `json:"card_token"`
}
