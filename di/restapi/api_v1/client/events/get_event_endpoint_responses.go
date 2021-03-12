// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"webank/DI/restapi/api_v1/restmodels"
)

// GetEventEndpointReader is a Reader for the GetEventEndpoint structure.
type GetEventEndpointReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEventEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEventEndpointOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetEventEndpointUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetEventEndpointNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 410:
		result := NewGetEventEndpointGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventEndpointOK creates a GetEventEndpointOK with default headers values
func NewGetEventEndpointOK() *GetEventEndpointOK {
	return &GetEventEndpointOK{}
}

/*GetEventEndpointOK handles this case with default header values.

Event description
*/
type GetEventEndpointOK struct {
	Payload *restmodels.Endpoint
}

func (o *GetEventEndpointOK) Error() string {
	return fmt.Sprintf("[GET /v1/models/{model_id}/events/{event_type}/{endpoint_id}][%d] getEventEndpointOK  %+v", 200, o.Payload)
}

func (o *GetEventEndpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Endpoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventEndpointUnauthorized creates a GetEventEndpointUnauthorized with default headers values
func NewGetEventEndpointUnauthorized() *GetEventEndpointUnauthorized {
	return &GetEventEndpointUnauthorized{}
}

/*GetEventEndpointUnauthorized handles this case with default header values.

Unauthorized
*/
type GetEventEndpointUnauthorized struct {
	Payload *restmodels.Error
}

func (o *GetEventEndpointUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/models/{model_id}/events/{event_type}/{endpoint_id}][%d] getEventEndpointUnauthorized  %+v", 401, o.Payload)
}

func (o *GetEventEndpointUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventEndpointNotFound creates a GetEventEndpointNotFound with default headers values
func NewGetEventEndpointNotFound() *GetEventEndpointNotFound {
	return &GetEventEndpointNotFound{}
}

/*GetEventEndpointNotFound handles this case with default header values.

The model or event type cannot be found.
*/
type GetEventEndpointNotFound struct {
	Payload *restmodels.Error
}

func (o *GetEventEndpointNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/models/{model_id}/events/{event_type}/{endpoint_id}][%d] getEventEndpointNotFound  %+v", 404, o.Payload)
}

func (o *GetEventEndpointNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEventEndpointGone creates a GetEventEndpointGone with default headers values
func NewGetEventEndpointGone() *GetEventEndpointGone {
	return &GetEventEndpointGone{}
}

/*GetEventEndpointGone handles this case with default header values.

If the trained model storage time has expired and it has been deleted. It only gets deleted if it not stored on an external data store.
*/
type GetEventEndpointGone struct {
	Payload *restmodels.Error
}

func (o *GetEventEndpointGone) Error() string {
	return fmt.Sprintf("[GET /v1/models/{model_id}/events/{event_type}/{endpoint_id}][%d] getEventEndpointGone  %+v", 410, o.Payload)
}

func (o *GetEventEndpointGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}