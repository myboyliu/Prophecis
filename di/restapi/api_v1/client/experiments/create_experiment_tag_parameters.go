// Code generated by go-swagger; DO NOT EDIT.

package experiments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	restmodels "webank/DI/restapi/api_v1/restmodels"
)

// NewCreateExperimentTagParams creates a new CreateExperimentTagParams object
// with the default values initialized.
func NewCreateExperimentTagParams() *CreateExperimentTagParams {
	var ()
	return &CreateExperimentTagParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateExperimentTagParamsWithTimeout creates a new CreateExperimentTagParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateExperimentTagParamsWithTimeout(timeout time.Duration) *CreateExperimentTagParams {
	var ()
	return &CreateExperimentTagParams{

		timeout: timeout,
	}
}

// NewCreateExperimentTagParamsWithContext creates a new CreateExperimentTagParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateExperimentTagParamsWithContext(ctx context.Context) *CreateExperimentTagParams {
	var ()
	return &CreateExperimentTagParams{

		Context: ctx,
	}
}

// NewCreateExperimentTagParamsWithHTTPClient creates a new CreateExperimentTagParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateExperimentTagParamsWithHTTPClient(client *http.Client) *CreateExperimentTagParams {
	var ()
	return &CreateExperimentTagParams{
		HTTPClient: client,
	}
}

/*CreateExperimentTagParams contains all the parameters to send to the API endpoint
for the create experiment tag operation typically these are written to a http.Request
*/
type CreateExperimentTagParams struct {

	/*ExperimentTag
	  The Experiment Tag Request

	*/
	ExperimentTag *restmodels.ProphecisExperimentTagPostRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create experiment tag params
func (o *CreateExperimentTagParams) WithTimeout(timeout time.Duration) *CreateExperimentTagParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create experiment tag params
func (o *CreateExperimentTagParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create experiment tag params
func (o *CreateExperimentTagParams) WithContext(ctx context.Context) *CreateExperimentTagParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create experiment tag params
func (o *CreateExperimentTagParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create experiment tag params
func (o *CreateExperimentTagParams) WithHTTPClient(client *http.Client) *CreateExperimentTagParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create experiment tag params
func (o *CreateExperimentTagParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExperimentTag adds the experimentTag to the create experiment tag params
func (o *CreateExperimentTagParams) WithExperimentTag(experimentTag *restmodels.ProphecisExperimentTagPostRequest) *CreateExperimentTagParams {
	o.SetExperimentTag(experimentTag)
	return o
}

// SetExperimentTag adds the experimentTag to the create experiment tag params
func (o *CreateExperimentTagParams) SetExperimentTag(experimentTag *restmodels.ProphecisExperimentTagPostRequest) {
	o.ExperimentTag = experimentTag
}

// WriteToRequest writes these params to a swagger request
func (o *CreateExperimentTagParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExperimentTag != nil {
		if err := r.SetBodyParam(o.ExperimentTag); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}