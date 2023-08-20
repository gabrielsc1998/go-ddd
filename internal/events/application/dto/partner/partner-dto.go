package partner_dto

type PartnerRegisterDto struct {
	Name string `json:"name"`
}

type PartnerUpdateDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
