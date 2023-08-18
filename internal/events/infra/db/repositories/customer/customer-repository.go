package customer_repository

import (
	"gorm.io/gorm"

	entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/customer"
	mapper "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/mappers/customer"
	model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/customer"
)

type CustomerRepository struct {
	db     *gorm.DB
	mapper *mapper.CustomerMapper
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
	mapper := mapper.NewCustomerMapper()
	return &CustomerRepository{db: db, mapper: mapper}
}

func (r *CustomerRepository) Add(customer *entity.Customer) error {
	return r.db.Create(r.mapper.ToModel(customer)).Error
}

func (r *CustomerRepository) Update(event *entity.Customer) error {
	return r.db.Updates(r.mapper.ToModel(event)).Error
}

func (r *CustomerRepository) FindById(id string) (*entity.Customer, error) {
	var customer model.Customer
	err := r.db.Where("id = ?", id).First(&customer).Error
	return r.mapper.ToEntity(&customer), err
}

func (r *CustomerRepository) FindAll() ([]*entity.Customer, error) {
	var customers []*model.Customer
	err := r.db.Find(&customers).Error
	var customersEntity []*entity.Customer
	for _, customer := range customers {
		customersEntity = append(customersEntity, r.mapper.ToEntity(customer))
	}
	return customersEntity, err
}

func (r *CustomerRepository) Delete(customer *entity.Customer) error {
	return r.db.Delete(r.mapper.ToModel(customer)).Error
}
