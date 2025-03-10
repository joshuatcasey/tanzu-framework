// Code generated by go-swagger; DO NOT EDIT.

package azure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vmware-tanzu/tanzu-framework/tkg/web/server/models"
)

// GetAzureVnetsOKCode is the HTTP code returned for type GetAzureVnetsOK
const GetAzureVnetsOKCode int = 200

/*
GetAzureVnetsOK Successful retrieval of Azure virtual networks

swagger:response getAzureVnetsOK
*/
type GetAzureVnetsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.AzureVirtualNetwork `json:"body,omitempty"`
}

// NewGetAzureVnetsOK creates GetAzureVnetsOK with default headers values
func NewGetAzureVnetsOK() *GetAzureVnetsOK {

	return &GetAzureVnetsOK{}
}

// WithPayload adds the payload to the get azure vnets o k response
func (o *GetAzureVnetsOK) WithPayload(payload []*models.AzureVirtualNetwork) *GetAzureVnetsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure vnets o k response
func (o *GetAzureVnetsOK) SetPayload(payload []*models.AzureVirtualNetwork) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureVnetsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.AzureVirtualNetwork, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetAzureVnetsBadRequestCode is the HTTP code returned for type GetAzureVnetsBadRequest
const GetAzureVnetsBadRequestCode int = 400

/*
GetAzureVnetsBadRequest Bad Request

swagger:response getAzureVnetsBadRequest
*/
type GetAzureVnetsBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureVnetsBadRequest creates GetAzureVnetsBadRequest with default headers values
func NewGetAzureVnetsBadRequest() *GetAzureVnetsBadRequest {

	return &GetAzureVnetsBadRequest{}
}

// WithPayload adds the payload to the get azure vnets bad request response
func (o *GetAzureVnetsBadRequest) WithPayload(payload *models.Error) *GetAzureVnetsBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure vnets bad request response
func (o *GetAzureVnetsBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureVnetsBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAzureVnetsUnauthorizedCode is the HTTP code returned for type GetAzureVnetsUnauthorized
const GetAzureVnetsUnauthorizedCode int = 401

/*
GetAzureVnetsUnauthorized Incorrect credentials

swagger:response getAzureVnetsUnauthorized
*/
type GetAzureVnetsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureVnetsUnauthorized creates GetAzureVnetsUnauthorized with default headers values
func NewGetAzureVnetsUnauthorized() *GetAzureVnetsUnauthorized {

	return &GetAzureVnetsUnauthorized{}
}

// WithPayload adds the payload to the get azure vnets unauthorized response
func (o *GetAzureVnetsUnauthorized) WithPayload(payload *models.Error) *GetAzureVnetsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure vnets unauthorized response
func (o *GetAzureVnetsUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureVnetsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetAzureVnetsInternalServerErrorCode is the HTTP code returned for type GetAzureVnetsInternalServerError
const GetAzureVnetsInternalServerErrorCode int = 500

/*
GetAzureVnetsInternalServerError Internal server error

swagger:response getAzureVnetsInternalServerError
*/
type GetAzureVnetsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetAzureVnetsInternalServerError creates GetAzureVnetsInternalServerError with default headers values
func NewGetAzureVnetsInternalServerError() *GetAzureVnetsInternalServerError {

	return &GetAzureVnetsInternalServerError{}
}

// WithPayload adds the payload to the get azure vnets internal server error response
func (o *GetAzureVnetsInternalServerError) WithPayload(payload *models.Error) *GetAzureVnetsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get azure vnets internal server error response
func (o *GetAzureVnetsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetAzureVnetsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
