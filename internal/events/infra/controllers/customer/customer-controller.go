package customer_controller

import (
	"encoding/json"
	"net/http"

	customer_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/customer"
	customer_service "github.com/gabrielsc1998/go-ddd/internal/events/application/services/customer"
	customer_presenter "github.com/gabrielsc1998/go-ddd/internal/events/infra/presenter/customer"
)

type CustomerController struct {
	customerService customer_service.CustomerService
}

func NewCustomerController(customerService customer_service.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: customerService,
	}
}

func (c *CustomerController) RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	var dto CreateCustomerInputDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.customerService.Register(customer_dto.CustomerRegisterDto{
		Name: dto.Name,
		CPF:  dto.CPF,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (c *CustomerController) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var dto UpdateCustomerInputDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.customerService.Update(customer_dto.CustomerUpdateDto{
		Name: dto.Name,
		Id:   dto.Id,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (c *CustomerController) FindCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := c.customerService.FindCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	customersPresenter := make([]customer_presenter.CustomerPresenter, len(customers))
	for i, customer := range customers {
		customersPresenter[i] = customer_presenter.ToPresent(customer)
	}
	json.NewEncoder(w).Encode(customersPresenter)
}
