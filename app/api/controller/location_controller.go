package controller

import (
	"drone-navigation-service-master/app/api/preprocessor"
	"drone-navigation-service-master/app/config/locales/local_config"
	"drone-navigation-service-master/app/model/request"
	"drone-navigation-service-master/app/processor"
	"drone-navigation-service-master/app/response_handler"
	"github.com/gorilla/mux"
	"net/http"
)

// swagger:operation GET /location getLocation
// ---
// summary: Get Location
// description: Get Location
// parameters:
// - name: sector_id
//   in: query
//   description: sector id
//   type: string
//   required: true
// - name: company_id
//   in: query
//   description: company id
//   type: string
//   required: true
// - name: z
//   in: query
//   description: z coordinate
//   type: string
//   required: true
// - name: x
//   in: query
//   description: x coordinate
//   type: string
//   required: true
// - name: y
//   in: query
//   description: y coordinate
//   type: string
//   required: true
// - name: vel
//   in: query
//   description: velocity
//   type: string
//   required: true
//   Responses:
//     default: body:genericError
//     200: body:genericModel
// swagger:route GET /location getLocation
//
// Get Location
//
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//
//     Responses:
//       default: body:genericError
//       200: body:genericModel
func GetLocation(w http.ResponseWriter, r *http.Request) {

	var locationRequest request.LocationRequest
	var params request.LocationStringToFloatRequest

	sectorId := mux.Vars(r)["sectorId"]
	err := preprocessor.DecodeAndValidateRequestParams(r, &locationRequest)

	if err != nil || sectorId == "" {
		msg := local_config.GetTranslationMessage(r.Context(), "invalid_request_params")
		response_handler.WriteErrorResponseAsJson(w, r, http.StatusBadRequest, msg)
		return
	}

	locationRequest.SectorId = sectorId

	_, err = preprocessor.BindAndUnMarshallRequest(r, &locationRequest, &params)
	if err != nil {
		msg := local_config.GetTranslationMessage(r.Context(), "invalid_request_params")
		response_handler.WriteErrorResponseAsJson(w, r, http.StatusBadRequest, msg)
		return
	}

	locationService := processor.GetNewLocationService()

	response, locationErr := locationService.GetLocation(r.Context(), params, locationRequest)

	if locationErr != nil {
		msg := local_config.GetTranslationMessage(r.Context(), "invalid_companyId_or_sectorId")
		response_handler.WriteErrorResponseAsJson(w, r, http.StatusBadRequest, msg)
		return
	}

	response_handler.WriteResponseMapAsJson(w, r, http.StatusOK, response)

}
