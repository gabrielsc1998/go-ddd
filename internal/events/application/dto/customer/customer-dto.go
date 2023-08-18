package customer_dto

type CustomerRegisterDto struct {
	Name string `json:"name"`
	CPF  string `json:"cpf"`
}

type CustomerUpdateDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
