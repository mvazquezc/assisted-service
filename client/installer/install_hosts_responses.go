// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// InstallHostsReader is a Reader for the InstallHosts structure.
type InstallHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InstallHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewInstallHostsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInstallHostsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInstallHostsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInstallHostsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInstallHostsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewInstallHostsMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewInstallHostsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInstallHostsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInstallHostsAccepted creates a InstallHostsAccepted with default headers values
func NewInstallHostsAccepted() *InstallHostsAccepted {
	return &InstallHostsAccepted{}
}

/*InstallHostsAccepted handles this case with default header values.

Success.
*/
type InstallHostsAccepted struct {
	Payload *models.Cluster
}

func (o *InstallHostsAccepted) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/actions/install_hosts][%d] installHostsAccepted  %+v", 202, o.Payload)
}

func (o *InstallHostsAccepted) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *InstallHostsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallHostsBadRequest creates a InstallHostsBadRequest with default headers values
func NewInstallHostsBadRequest() *InstallHostsBadRequest {
	return &InstallHostsBadRequest{}
}

/*InstallHostsBadRequest handles this case with default header values.

Error.
*/
type InstallHostsBadRequest struct {
	Payload *models.Error
}

func (o *InstallHostsBadRequest) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/actions/install_hosts][%d] installHostsBadRequest  %+v", 400, o.Payload)
}

func (o *InstallHostsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *InstallHostsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallHostsUnauthorized creates a InstallHostsUnauthorized with default headers values
func NewInstallHostsUnauthorized() *InstallHostsUnauthorized {
	return &InstallHostsUnauthorized{}
}

/*InstallHostsUnauthorized handles this case with default header values.

Unauthorized.
*/
type InstallHostsUnauthorized struct {
	Payload *models.InfraError
}

func (o *InstallHostsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/actions/install_hosts][%d] installHostsUnauthorized  %+v", 401, o.Payload)
}

func (o *InstallHostsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *InstallHostsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallHostsForbidden creates a InstallHostsForbidden with default headers values
func NewInstallHostsForbidden() *InstallHostsForbidden {
	return &InstallHostsForbidden{}
}

/*InstallHostsForbidden handles this case with default header values.

Forbidden.
*/
type InstallHostsForbidden struct {
	Payload *models.InfraError
}

func (o *InstallHostsForbidden) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/actions/install_hosts][%d] installHostsForbidden  %+v", 403, o.Payload)
}

func (o *InstallHostsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *InstallHostsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallHostsNotFound creates a InstallHostsNotFound with default headers values
func NewInstallHostsNotFound() *InstallHostsNotFound {
	return &InstallHostsNotFound{}
}

/*InstallHostsNotFound handles this case with default header values.

Error.
*/
type InstallHostsNotFound struct {
	Payload *models.Error
}

func (o *InstallHostsNotFound) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/actions/install_hosts][%d] installHostsNotFound  %+v", 404, o.Payload)
}

func (o *InstallHostsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *InstallHostsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallHostsMethodNotAllowed creates a InstallHostsMethodNotAllowed with default headers values
func NewInstallHostsMethodNotAllowed() *InstallHostsMethodNotAllowed {
	return &InstallHostsMethodNotAllowed{}
}

/*InstallHostsMethodNotAllowed handles this case with default header values.

Method Not Allowed.
*/
type InstallHostsMethodNotAllowed struct {
	Payload *models.Error
}

func (o *InstallHostsMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/actions/install_hosts][%d] installHostsMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *InstallHostsMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *InstallHostsMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallHostsConflict creates a InstallHostsConflict with default headers values
func NewInstallHostsConflict() *InstallHostsConflict {
	return &InstallHostsConflict{}
}

/*InstallHostsConflict handles this case with default header values.

Error.
*/
type InstallHostsConflict struct {
	Payload *models.Error
}

func (o *InstallHostsConflict) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/actions/install_hosts][%d] installHostsConflict  %+v", 409, o.Payload)
}

func (o *InstallHostsConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *InstallHostsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInstallHostsInternalServerError creates a InstallHostsInternalServerError with default headers values
func NewInstallHostsInternalServerError() *InstallHostsInternalServerError {
	return &InstallHostsInternalServerError{}
}

/*InstallHostsInternalServerError handles this case with default header values.

Error.
*/
type InstallHostsInternalServerError struct {
	Payload *models.Error
}

func (o *InstallHostsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /clusters/{cluster_id}/actions/install_hosts][%d] installHostsInternalServerError  %+v", 500, o.Payload)
}

func (o *InstallHostsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InstallHostsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
