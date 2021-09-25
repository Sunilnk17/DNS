package processor

import (
	"context"
	"drone-navigation-service/app/config"
	"drone-navigation-service/app/model/request"
	"errors"
)

//go:generate sh -c "$GOPATH/bin/mockery -case=underscore -dir=. -name=LocationService"
type LocationService interface {
	GetLocation(ctx context.Context, request request.LocationStringtoFloatRequest) (map[string]interface{}, error)
}

type locationService struct {
}

func GetNewLocationService() *locationService {
	return &locationService{}
}

func (ls locationService) GetLocation(ctx context.Context, floatRequest request.LocationStringtoFloatRequest, locationRequest request.LocationRequest) (map[string]interface{}, error) {

	sectors := config.GetConfig().Sectors

	response := make(map[string]interface{})

	responseName := ""

	for _, sector := range sectors {

		if sector.Id == locationRequest.SectorId {
			companies := sector.Companies
			for _, company := range companies {
				if company.Name == locationRequest.CompanyId {
					responseName = company.Response
				}
			}
		}

	}

	if responseName != "" {
		response[responseName] = ls.doCalculation(ctx, floatRequest, floatRequest.SectorId)
		return response, nil
	} else {
		return nil, errors.New("Invalid SectorId or CompanyId")
	}
}

func (ls locationService) doCalculation(ctx context.Context, floatRequest request.LocationStringtoFloatRequest, sectorId float64) float64 {

	result := sectorId*floatRequest.X + sectorId*floatRequest.Y + sectorId*floatRequest.Z + floatRequest.Vel
	return result
}
