package main

import (
	"context"
	"fmt"
	"time"

	unit_of_work "github.com/gabrielsc1998/go-ddd/internal/common/infra/db/unit-of-work"
	"github.com/gabrielsc1998/go-ddd/internal/database"
	event_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/event"
	event_service "github.com/gabrielsc1998/go-ddd/internal/events/application/services/event"
	event_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/event"
	partner_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/partner"
	section_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/section"
	spot_model "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/models/spot"
	event_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/event"
	partner_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/partner"
	"gorm.io/gorm"
)

func main() {
	db := database.NewDatabase()
	err := db.ConnectMySQL(database.DatabaseMySQLOptions{
		Host:     "localhost",
		Port:     "3306",
		User:     "root",
		Password: "root",
		Database: "events",
	})
	panicIfHasError(err)
	fmt.Println("Connected to MySQL")
	err = db.DB.AutoMigrate(&event_model.Event{}, &section_model.Section{}, &spot_model.Spot{}, &partner_model.Partner{})
	panicIfHasError(err)

	uow := unit_of_work.NewUow(context.Background(), db.DB)
	uow.Register("EventRepository", func(db *gorm.DB) interface{} {
		repo := event_repository.NewEventRepository(db)
		return repo
	})

	panicIfHasError(err)

	eventService := event_service.NewEventService(event_service.EventServiceProps{
		UOW:               uow,
		EventRepository:   event_repository.NewEventRepository(db.DB),
		PartnerRepository: partner_repository.NewPartnerRepository(db.DB),
	})

	err = eventService.Create(event_dto.EventCreateDto{
		Name:               "Teste",
		Description:        "Teste",
		Date:               time.Now(),
		IsPublished:        true,
		TotalSpots:         0,
		TotalSpotsReserved: 0,
		PartnerId:          "ffd86f59-3d60-47aa-80ce-410474f954d3",
	})
	panicIfHasError(err)

	// events, err := eventService.FindEvents()
	// panicIfHasError(err)

	// for _, event := range events {
	// 	// fmt.Println(event.Id.Value)
	// 	// fmt.Println(event.Name)
	// 	// fmt.Println(event.Description)
	// 	// fmt.Println(event.Date)
	// 	// fmt.Println(event.IsPublished)
	// 	// fmt.Println(event.TotalSpots)
	// 	// fmt.Println(event.TotalSpotsReserved)
	// 	// fmt.Println(event.PartnerId.Value)
	// }

	// eventService.AddSection(event_dto.EventAddSectionDto{
	// 	EventId:            "b8f7c3ce-ab0d-456e-a3c3-f90de050ee29",
	// 	Name:               "Teste",
	// 	Description:        "Teste",
	// 	Date:               time.Now(),
	// 	IsPublished:        true,
	// 	TotalSpots:         1,
	// 	TotalSpotsReserved: 0,
	// 	Price:              0,
	// })

	// eventService.PublishAll("b8f7c3ce-ab0d-456e-a3c3-f90de050ee29")

	// spots, err := eventService.FindSpots(event_dto.EventFindSpotsDto{
	// 	EventId:   "b8f7c3ce-ab0d-456e-a3c3-f90de050ee29",
	// 	SectionId: "d10cfc62-8531-4ccc-a273-9721f4fed032",
	// })
	// panicIfHasError(err)
	// for _, spot := range spots {
	// 	fmt.Println("v", spot.Id.Value)
	// 	fmt.Println("l", spot.Location)
	// 	fmt.Println("r", spot.IsReserved)
	// 	fmt.Println("p", spot.IsPublished)
	// }

	// err = eventService.UpdateLocation(event_dto.EventUpdateLocationDto{
	// 	EventId:   "b8f7c3ce-ab0d-456e-a3c3-f90de050ee29",
	// 	SectionId: "d10cfc62-8531-4ccc-a273-9721f4fed032",
	// 	SpotId:    "d4e460fc-18e3-44e6-8c3d-3265848b94c0",
	// 	Location:  "outra location",
	// })
	// panicIfHasError(err)

	// section, _ := eventService.FindSections("b8f7c3ce-ab0d-456e-a3c3-f90de050ee29")
	// fmt.Println("here", section[0].Spots[0].Id.Value)

	eventService.Update(event_dto.EventUpdateDto{
		Id:          "b8f7c3ce-ab0d-456e-a3c3-f90de050ee29",
		Name:        "updated",
		Description: "updated",
	})
	for {
	}
}

func panicIfHasError(err error) {
	if err != nil {
		panic(err)
	}
}
