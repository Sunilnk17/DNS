// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LocationStringtoFloatRequest location stringto float request
// swagger:model LocationStringtoFloatRequest
type LocationStringtoFloatRequest struct {

	// company Id
	CompanyID string `json:"companyId,omitempty"`

	// sector Id
	SectorID float64 `json:"SectorId,omitempty"`

	// vel
	Vel float64 `json:"Vel,omitempty"`

	// x
	X float64 `json:"X,omitempty"`

	// y
	Y float64 `json:"Y,omitempty"`

	// z
	Z float64 `json:"Z,omitempty"`
}

// Validate validates this location stringto float request
func (m *LocationStringtoFloatRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LocationStringtoFloatRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocationStringtoFloatRequest) UnmarshalBinary(b []byte) error {
	var res LocationStringtoFloatRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
