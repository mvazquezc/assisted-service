// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openshift/assisted-service/models"
)

// DownloadClusterFilesOKCode is the HTTP code returned for type DownloadClusterFilesOK
const DownloadClusterFilesOKCode int = 200

/*DownloadClusterFilesOK Success.

swagger:response downloadClusterFilesOK
*/
type DownloadClusterFilesOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewDownloadClusterFilesOK creates DownloadClusterFilesOK with default headers values
func NewDownloadClusterFilesOK() *DownloadClusterFilesOK {

	return &DownloadClusterFilesOK{}
}

// WithPayload adds the payload to the download cluster files o k response
func (o *DownloadClusterFilesOK) WithPayload(payload io.ReadCloser) *DownloadClusterFilesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster files o k response
func (o *DownloadClusterFilesOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterFilesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DownloadClusterFilesUnauthorizedCode is the HTTP code returned for type DownloadClusterFilesUnauthorized
const DownloadClusterFilesUnauthorizedCode int = 401

/*DownloadClusterFilesUnauthorized Unauthorized.

swagger:response downloadClusterFilesUnauthorized
*/
type DownloadClusterFilesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadClusterFilesUnauthorized creates DownloadClusterFilesUnauthorized with default headers values
func NewDownloadClusterFilesUnauthorized() *DownloadClusterFilesUnauthorized {

	return &DownloadClusterFilesUnauthorized{}
}

// WithPayload adds the payload to the download cluster files unauthorized response
func (o *DownloadClusterFilesUnauthorized) WithPayload(payload *models.InfraError) *DownloadClusterFilesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster files unauthorized response
func (o *DownloadClusterFilesUnauthorized) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterFilesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterFilesForbiddenCode is the HTTP code returned for type DownloadClusterFilesForbidden
const DownloadClusterFilesForbiddenCode int = 403

/*DownloadClusterFilesForbidden Forbidden.

swagger:response downloadClusterFilesForbidden
*/
type DownloadClusterFilesForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.InfraError `json:"body,omitempty"`
}

// NewDownloadClusterFilesForbidden creates DownloadClusterFilesForbidden with default headers values
func NewDownloadClusterFilesForbidden() *DownloadClusterFilesForbidden {

	return &DownloadClusterFilesForbidden{}
}

// WithPayload adds the payload to the download cluster files forbidden response
func (o *DownloadClusterFilesForbidden) WithPayload(payload *models.InfraError) *DownloadClusterFilesForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster files forbidden response
func (o *DownloadClusterFilesForbidden) SetPayload(payload *models.InfraError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterFilesForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterFilesNotFoundCode is the HTTP code returned for type DownloadClusterFilesNotFound
const DownloadClusterFilesNotFoundCode int = 404

/*DownloadClusterFilesNotFound Error.

swagger:response downloadClusterFilesNotFound
*/
type DownloadClusterFilesNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterFilesNotFound creates DownloadClusterFilesNotFound with default headers values
func NewDownloadClusterFilesNotFound() *DownloadClusterFilesNotFound {

	return &DownloadClusterFilesNotFound{}
}

// WithPayload adds the payload to the download cluster files not found response
func (o *DownloadClusterFilesNotFound) WithPayload(payload *models.Error) *DownloadClusterFilesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster files not found response
func (o *DownloadClusterFilesNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterFilesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterFilesMethodNotAllowedCode is the HTTP code returned for type DownloadClusterFilesMethodNotAllowed
const DownloadClusterFilesMethodNotAllowedCode int = 405

/*DownloadClusterFilesMethodNotAllowed Method Not Allowed.

swagger:response downloadClusterFilesMethodNotAllowed
*/
type DownloadClusterFilesMethodNotAllowed struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterFilesMethodNotAllowed creates DownloadClusterFilesMethodNotAllowed with default headers values
func NewDownloadClusterFilesMethodNotAllowed() *DownloadClusterFilesMethodNotAllowed {

	return &DownloadClusterFilesMethodNotAllowed{}
}

// WithPayload adds the payload to the download cluster files method not allowed response
func (o *DownloadClusterFilesMethodNotAllowed) WithPayload(payload *models.Error) *DownloadClusterFilesMethodNotAllowed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster files method not allowed response
func (o *DownloadClusterFilesMethodNotAllowed) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterFilesMethodNotAllowed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(405)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterFilesConflictCode is the HTTP code returned for type DownloadClusterFilesConflict
const DownloadClusterFilesConflictCode int = 409

/*DownloadClusterFilesConflict Error.

swagger:response downloadClusterFilesConflict
*/
type DownloadClusterFilesConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterFilesConflict creates DownloadClusterFilesConflict with default headers values
func NewDownloadClusterFilesConflict() *DownloadClusterFilesConflict {

	return &DownloadClusterFilesConflict{}
}

// WithPayload adds the payload to the download cluster files conflict response
func (o *DownloadClusterFilesConflict) WithPayload(payload *models.Error) *DownloadClusterFilesConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster files conflict response
func (o *DownloadClusterFilesConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterFilesConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterFilesInternalServerErrorCode is the HTTP code returned for type DownloadClusterFilesInternalServerError
const DownloadClusterFilesInternalServerErrorCode int = 500

/*DownloadClusterFilesInternalServerError Error.

swagger:response downloadClusterFilesInternalServerError
*/
type DownloadClusterFilesInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterFilesInternalServerError creates DownloadClusterFilesInternalServerError with default headers values
func NewDownloadClusterFilesInternalServerError() *DownloadClusterFilesInternalServerError {

	return &DownloadClusterFilesInternalServerError{}
}

// WithPayload adds the payload to the download cluster files internal server error response
func (o *DownloadClusterFilesInternalServerError) WithPayload(payload *models.Error) *DownloadClusterFilesInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster files internal server error response
func (o *DownloadClusterFilesInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterFilesInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DownloadClusterFilesServiceUnavailableCode is the HTTP code returned for type DownloadClusterFilesServiceUnavailable
const DownloadClusterFilesServiceUnavailableCode int = 503

/*DownloadClusterFilesServiceUnavailable Unavailable.

swagger:response downloadClusterFilesServiceUnavailable
*/
type DownloadClusterFilesServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDownloadClusterFilesServiceUnavailable creates DownloadClusterFilesServiceUnavailable with default headers values
func NewDownloadClusterFilesServiceUnavailable() *DownloadClusterFilesServiceUnavailable {

	return &DownloadClusterFilesServiceUnavailable{}
}

// WithPayload adds the payload to the download cluster files service unavailable response
func (o *DownloadClusterFilesServiceUnavailable) WithPayload(payload *models.Error) *DownloadClusterFilesServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster files service unavailable response
func (o *DownloadClusterFilesServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterFilesServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
