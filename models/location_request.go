// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LocationRequest Get location Request
// swagger:model LocationRequest
type LocationRequest struct {

	// company Id
	CompanyID string `json:"companyId,omitempty"`

	// a location request
	// in: body
	SectorID string `json:"sectorId,omitempty"`

	// vel
	Vel string `json:"vel,omitempty"`

	// x
	X string `json:"x,omitempty"`

	// y
	Y string `json:"y,omitempty"`

	// z
	Z string `json:"z,omitempty"`
}

// Validate validates this location request
func (m *LocationRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LocationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocationRequest) UnmarshalBinary(b []byte) error {
	var res LocationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
