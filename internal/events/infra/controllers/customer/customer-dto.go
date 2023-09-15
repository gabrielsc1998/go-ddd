package customer_controller

type CreateCustomerInputDto struct {
	Name string `json:"name"`
	CPF  string `json:"cpf"`
}

type UpdateCustomerInputDto struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
