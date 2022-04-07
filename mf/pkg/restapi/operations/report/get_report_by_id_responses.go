// Code generated by go-swagger; DO NOT EDIT.

package report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "mlss-mf/pkg/models"
)

// GetReportByIDOKCode is the HTTP code returned for type GetReportByIDOK
const GetReportByIDOKCode int = 200

/*GetReportByIDOK OK

swagger:response getReportByIdOK
*/
type GetReportByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Report `json:"body,omitempty"`
}

// NewGetReportByIDOK creates GetReportByIDOK with default headers values
func NewGetReportByIDOK() *GetReportByIDOK {

	return &GetReportByIDOK{}
}

// WithPayload adds the payload to the get report by Id o k response
func (o *GetReportByIDOK) WithPayload(payload *models.Report) *GetReportByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get report by Id o k response
func (o *GetReportByIDOK) SetPayload(payload *models.Report) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReportByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReportByIDUnauthorizedCode is the HTTP code returned for type GetReportByIDUnauthorized
const GetReportByIDUnauthorizedCode int = 401

/*GetReportByIDUnauthorized Unauthorized

swagger:response getReportByIdUnauthorized
*/
type GetReportByIDUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetReportByIDUnauthorized creates GetReportByIDUnauthorized with default headers values
func NewGetReportByIDUnauthorized() *GetReportByIDUnauthorized {

	return &GetReportByIDUnauthorized{}
}

// WithPayload adds the payload to the get report by Id unauthorized response
func (o *GetReportByIDUnauthorized) WithPayload(payload *models.Error) *GetReportByIDUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get report by Id unauthorized response
func (o *GetReportByIDUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReportByIDUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReportByIDNotFoundCode is the HTTP code returned for type GetReportByIDNotFound
const GetReportByIDNotFoundCode int = 404

/*GetReportByIDNotFound Report get fail

swagger:response getReportByIdNotFound
*/
type GetReportByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetReportByIDNotFound creates GetReportByIDNotFound with default headers values
func NewGetReportByIDNotFound() *GetReportByIDNotFound {

	return &GetReportByIDNotFound{}
}

// WithPayload adds the payload to the get report by Id not found response
func (o *GetReportByIDNotFound) WithPayload(payload *models.Error) *GetReportByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get report by Id not found response
func (o *GetReportByIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReportByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}