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
