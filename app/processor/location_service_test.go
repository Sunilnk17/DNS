package processor

import (
	"context"
	"drone-navigation-service/app/config"
	"drone-navigation-service/app/model/request"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getSectorConfig() []config.SectorConfig {
	companyConfig := []config.CompanyConfig{{Name: "atlas", Response: "loc"}}
	sectorConfig := config.SectorConfig{}
	sectorConfig.Id = "1234"
	sectorConfig.Companies = companyConfig
	return []config.SectorConfig{sectorConfig}
}

func getLocationRequest() request.LocationRequest {

	locationRequest := request.LocationRequest{}
	locationRequest.SectorId = "1234"
	locationRequest.CompanyId = "atlas"
	locationRequest.Vel = "1.2"
	locationRequest.X = "2"
	locationRequest.Y = "3"
	locationRequest.Z = "5.7"

	return locationRequest
}

func getLocationtoFloatRequest() request.LocationStringtoFloatRequest {

	locationRequest := request.LocationStringtoFloatRequest{}
	locationRequest.SectorId = 1234
	locationRequest.CompanyId = "atlas"
	locationRequest.Vel = 1.2
	locationRequest.X = 2
	locationRequest.Y = 3
	locationRequest.Z = 5.7

	return locationRequest
}

func TestGetNewLocationService(t *testing.T) {
	config.SetDummyConfig()
	config.GetConfig().Sectors = getSectorConfig()

	service := GetNewLocationService()

	_, err := service.GetLocation(context.TODO(), getLocationtoFloatRequest(), getLocationRequest())

	assert.Nil(t, err)

}

func TestGetNewLocationServiceForInvalidCompany(t *testing.T) {
	config.SetDummyConfig()
	config.GetConfig().Sectors = getSectorConfig()

	service := GetNewLocationService()

	_, err := service.GetLocation(context.TODO(), request.LocationStringtoFloatRequest{SectorId: 12}, request.LocationRequest{SectorId: "12"})

	assert.NotNil(t, err)

}
