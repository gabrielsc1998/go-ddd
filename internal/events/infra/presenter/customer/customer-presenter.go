package customer_presenter

import customer_entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/customer"

type CustomerPresenter struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	CPF  string `json:"cpf"`
}

func ToPresent(event *customer_entity.Customer) CustomerPresenter {
	return CustomerPresenter{
		Id:   event.Id.Value,
		Name: event.Name,
		CPF:  event.CPF.Formatted,
	}
}
