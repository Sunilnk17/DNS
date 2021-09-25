package processor

import (
	"context"
	"drone-navigation-service-master/app/config"
	"drone-navigation-service-master/app/model/request"
	"errors"
)

// go:generate sh -c "$GOPATH/bin/mockery -case=underscore -dir=. -name=LocationService"
type LocationService interface {
	GetLocation(ctx context.Context, request request.LocationStringToFloatRequest) (map[string]interface{}, error)
}

type locationService struct {
}

func GetNewLocationService() *locationService {
	return &locationService{}
}

func (ls locationService) GetLocation(ctx context.Context, floatRequest request.LocationStringToFloatRequest, locationRequest request.LocationRequest) (map[string]interface{}, error) {

	sectors := config.GetConfig().Sectors
	response := make(map[string]interface{})
	fieldName := ""

	for _, sector := range sectors {
		if sector.Id == locationRequest.SectorId {
			companies := sector.Companies
			for _, company := range companies {
				if company.Name == locationRequest.CompanyId {
					fieldName = company.Response
				}
			}
		}
	}

	if fieldName != "" {
		response[fieldName] = doCalculation(floatRequest.X, floatRequest.Y, floatRequest.Z, floatRequest.Vel, floatRequest.SectorId)
		return response, nil
	} else {
		return nil, errors.New("invalid sectorId or companyId")
	}
}

func doCalculation(xValue float64, yValue float64, zValue float64, velocity float64, sectorId float64) float64 {
	result := sectorId*xValue + sectorId*yValue + sectorId*zValue + velocity
	return result
}
