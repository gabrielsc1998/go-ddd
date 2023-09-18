package order_repository

import (
	"gorm.io/gorm"

	entity "github.com/gabrielsc1998/go-ddd/internal/events/domain/entities/order"
	mapper "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/mappers/order"
	model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/order"
)

type OrderRepository struct {
	db     *gorm.DB
	mapper *mapper.OrderMapper
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	mapper := mapper.NewOrderMapper()
	return &OrderRepository{db: db, mapper: mapper}
}

func (r *OrderRepository) Add(order *entity.Order) error {
	orderExists, _ := r.FindById(order.Id.Value)
	if orderExists != nil {
		return r.db.Updates(r.mapper.ToModel(order)).Error
	}
	return r.db.Create(r.mapper.ToModel(order)).Error
}

func (r *OrderRepository) FindById(id string) (*entity.Order, error) {
	var order model.Order
	err := r.db.Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return r.mapper.ToEntity(&order), nil
}

func (r *OrderRepository) FindAll(eventId string) ([]*entity.Order, error) {
	var orders []*model.Order
	err := r.db.Where("event_id = ?", eventId).Find(&orders).Error
	var ordersEntity []*entity.Order
	for _, order := range orders {
		ordersEntity = append(ordersEntity, r.mapper.ToEntity(order))
	}
	return ordersEntity, err
}

func (r *OrderRepository) Delete(order *entity.Order) error {
	return r.db.Delete(r.mapper.ToModel(order)).Error
}
