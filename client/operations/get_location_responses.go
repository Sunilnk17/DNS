// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	models "drone-navigation-service/models"
	"fmt"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"io"
)

// GetLocationReader is a Reader for the GetLocation structure.
type GetLocationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLocationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLocationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetLocationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetLocationOK creates a GetLocationOK with default headers values
func NewGetLocationOK() *GetLocationOK {
	return &GetLocationOK{}
}

/*GetLocationOK handles this case with default header values.

genericModel
*/
type GetLocationOK struct {
	Payload *models.GenericResponse
}

func (o *GetLocationOK) Error() string {
	return fmt.Sprintf("[GET /location][%d] getLocationOK  %+v", 200, o.Payload)
}

func (o *GetLocationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLocationDefault creates a GetLocationDefault with default headers values
func NewGetLocationDefault(code int) *GetLocationDefault {
	return &GetLocationDefault{
		_statusCode: code,
	}
}

/*GetLocationDefault handles this case with default header values.

genericError
*/
type GetLocationDefault struct {
	_statusCode int

	Payload *models.GenericError
}

// Code gets the status code for the get location default response
func (o *GetLocationDefault) Code() int {
	return o._statusCode
}

func (o *GetLocationDefault) Error() string {
	return fmt.Sprintf("[GET /location][%d] getLocation default  %+v", o._statusCode, o.Payload)
}

func (o *GetLocationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
