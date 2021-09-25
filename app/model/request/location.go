package request

// Get location Request
// swagger:model locationRequest
type LocationRequest struct {
	// a location request
	// in: body
	SectorId  string `json:"sectorId" conform:"trim,omitempty"`
	CompanyId string `json:"companyId" conform:"trim,omitempty"`
	X         string `json:"x" validate:"required" conform:"trim,omitempty"`
	Y         string `json:"y" validate:"required" conform:"trim,omitempty"`
	Z         string `json:"z" validate:"required" conform:"trim,omitempty"`
	Vel       string `json:"vel" validate:"required" conform:"trim,omitempty"`
}

type LocationStringToFloatRequest struct {
	SectorId  float64 `json:",string"`
	CompanyId string  `json:"companyId" conform:"trim,omitempty"`
	X         float64 `json:",string"`
	Y         float64 `json:",string"`
	Z         float64 `json:",string"`
	Vel       float64 `json:",string"`
}

// Location Generic Error
// swagger:model genericError
type GenericError struct {
	Code      int    `json:"code"`
	Error     string `json:"error"`
	RequestId string `json:"request_id"`
}

// Location Generic Model
// swagger:model genericModel
type GenericResponse struct {
	Data      map[string]interface{} `json:"data"`
	Code      int                    `json:"code"`
	RequestId string                 `json:"request_id"`
	Error     string                 `json:"error"`
}
