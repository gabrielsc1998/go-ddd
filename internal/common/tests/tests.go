package tests

import (
	"context"

	unit_of_work "github.com/gabrielsc1998/go-ddd/internal/common/infra/db/unit-of-work"
	"github.com/gabrielsc1998/go-ddd/internal/database"
	customer_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/customer"
	event_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/event"
	order_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/order"
	partner_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/partner"
	section_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/section"
	spot_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/spot"
	spot_reservation_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/spot-reservation"
)

type Tests struct {
	DB  *database.Database
	UOW *unit_of_work.Uow
}

func Setup() *Tests {
	db := database.NewDatabase()
	err := db.ConnectMySQL(database.DatabaseMySQLOptions{
		Host:     "localhost",
		Port:     "3307",
		User:     "root",
		Password: "root",
		Database: "events",
	})
	panicIfHasError(err)

	err = db.DB.AutoMigrate(
		&event_model.Event{},
		&section_model.Section{},
		&spot_model.Spot{},
		&partner_model.Partner{},
		&customer_model.Customer{},
		&order_model.Order{},
		&spot_reservation_model.SpotReservation{},
	)
	panicIfHasError(err)

	db.DB.Exec("SET FOREIGN_KEY_CHECKS=0")
	db.DB.Exec("DELETE FROM spots")
	db.DB.Exec("DELETE FROM sections")
	db.DB.Exec("DELETE FROM events")
	db.DB.Exec("DELETE FROM partners")
	db.DB.Exec("DELETE FROM customers")
	db.DB.Exec("DELETE FROM orders")
	db.DB.Exec("DELETE FROM spot_reservations")
	db.DB.Exec("SET FOREIGN_KEY_CHECKS=1")

	uow := unit_of_work.NewUow(context.Background(), db.DB)
	return &Tests{
		DB:  db,
		UOW: uow,
	}
}

func panicIfHasError(err error) {
	if err != nil {
		panic(err)
	}
}
