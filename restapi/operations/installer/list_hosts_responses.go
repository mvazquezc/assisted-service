// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// ListHostsOKCode is the HTTP code returned for type ListHostsOK
const ListHostsOKCode int = 200

/*ListHostsOK Success.

swagger:response listHostsOK
*/
type ListHostsOK struct {

	/*
	  In: Body
	*/
	Payload models.HostList `json:"body,omitempty"`
}

// NewListHostsOK creates ListHostsOK with default headers values
func NewListHostsOK() *ListHostsOK {

	return &ListHostsOK{}
}

// WithPayload adds the payload to the list hosts o k response
func (o *ListHostsOK) WithPayload(payload models.HostList) *ListHostsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list hosts o k response
func (o *ListHostsOK) SetPayload(payload models.HostList) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListHostsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.HostList{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ListHostsUnauthorizedCode is the HTTP code returned for type ListHostsUnauthorized
const ListHostsUnauthorizedCode int = 401

/*ListHostsUnauthorized Unauthorized.

swagger:response listHostsUnauthorized
*/
type ListHostsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewListHostsUnauthorized creates ListHostsUnauthorized with default headers values
func NewListHostsUnauthorized() *ListHostsUnauthorized {

	return &ListHostsUnauthorized{}
}

// WithPayload adds the payload to the list hosts unauthorized response
func (o *ListHostsUnauthorized) WithPayload(payload *models.InfraError) *ListHostsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list hosts unauthorized response
func (o *ListHostsUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListHostsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListHostsForbiddenCode is the HTTP code returned for type ListHostsForbidden
const ListHostsForbiddenCode int = 403

/*ListHostsForbidden Forbidden.

swagger:response listHostsForbidden
*/
type ListHostsForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewListHostsForbidden creates ListHostsForbidden with default headers values
func NewListHostsForbidden() *ListHostsForbidden {

	return &ListHostsForbidden{}
}

// WithPayload adds the payload to the list hosts forbidden response
func (o *ListHostsForbidden) WithPayload(payload *models.InfraError) *ListHostsForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list hosts forbidden response
func (o *ListHostsForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListHostsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListHostsMethodNotAllowedCode is the HTTP code returned for type ListHostsMethodNotAllowed
const ListHostsMethodNotAllowedCode int = 405

/*ListHostsMethodNotAllowed Method Not Allowed.

swagger:response listHostsMethodNotAllowed
*/
type ListHostsMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListHostsMethodNotAllowed creates ListHostsMethodNotAllowed with default headers values
func NewListHostsMethodNotAllowed() *ListHostsMethodNotAllowed {

	return &ListHostsMethodNotAllowed{}
}

// WithPayload adds the payload to the list hosts method not allowed response
func (o *ListHostsMethodNotAllowed) WithPayload(payload *models.Error) *ListHostsMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list hosts method not allowed response
func (o *ListHostsMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListHostsMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListHostsInternalServerErrorCode is the HTTP code returned for type ListHostsInternalServerError
const ListHostsInternalServerErrorCode int = 500

/*ListHostsInternalServerError Error.

swagger:response listHostsInternalServerError
*/
type ListHostsInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListHostsInternalServerError creates ListHostsInternalServerError with default headers values
func NewListHostsInternalServerError() *ListHostsInternalServerError {

	return &ListHostsInternalServerError{}
}

// WithPayload adds the payload to the list hosts internal server error response
func (o *ListHostsInternalServerError) WithPayload(payload *models.Error) *ListHostsInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list hosts internal server error response
func (o *ListHostsInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListHostsInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListHostsServiceUnavailableCode is the HTTP code returned for type ListHostsServiceUnavailable
const ListHostsServiceUnavailableCode int = 503

/*ListHostsServiceUnavailable Unavailable.

swagger:response listHostsServiceUnavailable
*/
type ListHostsServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewListHostsServiceUnavailable creates ListHostsServiceUnavailable with default headers values
func NewListHostsServiceUnavailable() *ListHostsServiceUnavailable {

	return &ListHostsServiceUnavailable{}
}

// WithPayload adds the payload to the list hosts service unavailable response
func (o *ListHostsServiceUnavailable) WithPayload(payload *models.Error) *ListHostsServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list hosts service unavailable response
func (o *ListHostsServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListHostsServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
