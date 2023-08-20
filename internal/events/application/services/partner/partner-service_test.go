package partner_service

import (
	"testing"

	"github.com/gabrielsc1998/go-ddd/internal/common/tests"
	partner_dto "github.com/gabrielsc1998/go-ddd/internal/events/application/dto/partner"
	partner_repository "github.com/gabrielsc1998/go-ddd/internal/events/infra/db/repositories/partner"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var test *tests.Tests
var partnerService PartnerService

func Setup() {
	test = tests.Setup()
	test.UOW.Register("PartnerRepository", func(db *gorm.DB) interface{} {
		repo := partner_repository.NewPartnerRepository(db)
		return repo
	})
	partnerService = NewPartnerService(PartnerServiceProps{
		UOW:               test.UOW,
		PartnerRepository: partner_repository.NewPartnerRepository(test.DB.DB),
	})
}

func TestShouldRegisterAPartner(t *testing.T) {
	Setup()

	err := partnerService.Register(partner_dto.PartnerRegisterDto{
		Name: "Jhon Doe",
	})
	assert.Nil(t, err)

	repo := partner_repository.NewPartnerRepository(test.DB.DB)
	partner, err := repo.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(partner))
	assert.Equal(t, "Jhon Doe", partner[0].Name)
}

func TestShouldUpdateAPartner(t *testing.T) {
	Setup()

	err := partnerService.Register(partner_dto.PartnerRegisterDto{
		Name: "Jhon Doe",
	})
	assert.Nil(t, err)

	repo := partner_repository.NewPartnerRepository(test.DB.DB)
	partners, err := repo.FindAll()

	assert.Nil(t, err)
	assert.Equal(t, 1, len(partners))
	assert.NotNil(t, partners[0].Id.Value)
	assert.Equal(t, "Jhon Doe", partners[0].Name)

	partnerID := partners[0].Id.Value

	err = partnerService.Update(partner_dto.PartnerUpdateDto{
		Id:   partnerID,
		Name: "Jhon Doe 2",
	})
	assert.Nil(t, err)

	partner, err := repo.FindById(partnerID)

	assert.Nil(t, err)
	assert.Equal(t, "Jhon Doe 2", partner.Name)
}
