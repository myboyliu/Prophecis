// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new models API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for models API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteModel deletes an existing model

Deletes an existing model. It does not delete any data in the user's data store.

*/
func (a *Client) DeleteModel(params *DeleteModelParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteModelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteModelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteModel",
		Method:             "DELETE",
		PathPattern:        "/v1/models/{model_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteModelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteModelOK), nil

}

/*
DownloadModelDefinition downloads the model definition

Downloads the model definition that was initial used for training as ZIP archive.
*/
func (a *Client) DownloadModelDefinition(params *DownloadModelDefinitionParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*DownloadModelDefinitionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadModelDefinitionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "downloadModelDefinition",
		Method:             "GET",
		PathPattern:        "/v1/models/{model_id}/definition",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadModelDefinitionReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DownloadModelDefinitionOK), nil

}

/*
DownloadTrainedModel downloads the trained model

Downloads the trained model as ZIP archive.
*/
func (a *Client) DownloadTrainedModel(params *DownloadTrainedModelParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*DownloadTrainedModelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDownloadTrainedModelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "downloadTrainedModel",
		Method:             "GET",
		PathPattern:        "/v1/models/{model_id}/trained_model",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DownloadTrainedModelReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DownloadTrainedModelOK), nil

}

/*
GetLogs gets training logs as websocket stream

Get training logs for the given model as websocket stream. Each message can contain one or more log lines.

*/
func (a *Client) GetLogs(params *GetLogsParams, writer io.Writer) (*GetLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLogs",
		Method:             "GET",
		PathPattern:        "/v1/models/{model_id}/logs",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLogsReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLogsOK), nil

}

/*
GetMetrics gets training metrics from a model training as websocket

Get training metrics such as learning rate, accurarcy, loss from a model training via a websocket. It returns one or more JSON payloads at a time with the following schema.

*/
func (a *Client) GetMetrics(params *GetMetricsParams, writer io.Writer) (*GetMetricsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMetricsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMetrics",
		Method:             "GET",
		PathPattern:        "/v1/models/{model_id}/metrics",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMetricsReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMetricsOK), nil

}

/*
GetModel gets detailed information about a model

Get detailed information about a model such as training status.

*/
func (a *Client) GetModel(params *GetModelParams, authInfo runtime.ClientAuthInfoWriter) (*GetModelOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetModelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getModel",
		Method:             "GET",
		PathPattern:        "/v1/models/{model_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetModelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetModelOK), nil

}

/*
ListModels gets a list of available deep learning models

Get a list of all available deep learning models and their configuration that a user can see.

*/
func (a *Client) ListModels(params *ListModelsParams, authInfo runtime.ClientAuthInfoWriter) (*ListModelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListModelsParams()
	}
	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listModels",
		Method:             "GET",
		PathPattern:        "/v1/models",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListModelsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListModelsOK), nil

}

/*
PatchModel changes the status of the training progress

Changes the status of the training progress to the given `status` value (currently `halt` only). Halt means the training will be stopped and the last snapshot will be stored and can be retrieved.
*/
func (a *Client) PatchModel(params *PatchModelParams, authInfo runtime.ClientAuthInfoWriter) (*PatchModelAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchModelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchModel",
		Method:             "PATCH",
		PathPattern:        "/v1/models/{model_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchModelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchModelAccepted), nil

}

/*
PostModel trains a new deep learning model

Trains a deep neural network written in a DL framework supported by the DLaaS platform (such as Caffe, Tensorflow, etc.). The model code has to be uploaded and configuration parameters have to be provided.

*/
func (a *Client) PostModel(params *PostModelParams, authInfo runtime.ClientAuthInfoWriter) (*PostModelCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostModelParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postModel",
		Method:             "POST",
		PathPattern:        "/v1/models",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostModelReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostModelCreated), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}