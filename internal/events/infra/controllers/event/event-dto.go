package event_controller

type CreateEventInputDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
	PartnerId   string `json:"partner_id"`
}
