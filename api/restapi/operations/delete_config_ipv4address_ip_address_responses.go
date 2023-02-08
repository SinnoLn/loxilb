// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/loxilb-io/loxilb/api/models"
)

// DeleteConfigIpv4addressIPAddressNoContentCode is the HTTP code returned for type DeleteConfigIpv4addressIPAddressNoContent
const DeleteConfigIpv4addressIPAddressNoContentCode int = 204

/*
DeleteConfigIpv4addressIPAddressNoContent OK

swagger:response deleteConfigIpv4addressIpAddressNoContent
*/
type DeleteConfigIpv4addressIPAddressNoContent struct {
}

// NewDeleteConfigIpv4addressIPAddressNoContent creates DeleteConfigIpv4addressIPAddressNoContent with default headers values
func NewDeleteConfigIpv4addressIPAddressNoContent() *DeleteConfigIpv4addressIPAddressNoContent {

	return &DeleteConfigIpv4addressIPAddressNoContent{}
}

// WriteResponse to the client
func (o *DeleteConfigIpv4addressIPAddressNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteConfigIpv4addressIPAddressBadRequestCode is the HTTP code returned for type DeleteConfigIpv4addressIPAddressBadRequest
const DeleteConfigIpv4addressIPAddressBadRequestCode int = 400

/*
DeleteConfigIpv4addressIPAddressBadRequest Malformed arguments for API call

swagger:response deleteConfigIpv4addressIpAddressBadRequest
*/
type DeleteConfigIpv4addressIPAddressBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigIpv4addressIPAddressBadRequest creates DeleteConfigIpv4addressIPAddressBadRequest with default headers values
func NewDeleteConfigIpv4addressIPAddressBadRequest() *DeleteConfigIpv4addressIPAddressBadRequest {

	return &DeleteConfigIpv4addressIPAddressBadRequest{}
}

// WithPayload adds the payload to the delete config ipv4address Ip address bad request response
func (o *DeleteConfigIpv4addressIPAddressBadRequest) WithPayload(payload *models.Error) *DeleteConfigIpv4addressIPAddressBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config ipv4address Ip address bad request response
func (o *DeleteConfigIpv4addressIPAddressBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigIpv4addressIPAddressBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigIpv4addressIPAddressUnauthorizedCode is the HTTP code returned for type DeleteConfigIpv4addressIPAddressUnauthorized
const DeleteConfigIpv4addressIPAddressUnauthorizedCode int = 401

/*
DeleteConfigIpv4addressIPAddressUnauthorized Invalid authentication credentials

swagger:response deleteConfigIpv4addressIpAddressUnauthorized
*/
type DeleteConfigIpv4addressIPAddressUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigIpv4addressIPAddressUnauthorized creates DeleteConfigIpv4addressIPAddressUnauthorized with default headers values
func NewDeleteConfigIpv4addressIPAddressUnauthorized() *DeleteConfigIpv4addressIPAddressUnauthorized {

	return &DeleteConfigIpv4addressIPAddressUnauthorized{}
}

// WithPayload adds the payload to the delete config ipv4address Ip address unauthorized response
func (o *DeleteConfigIpv4addressIPAddressUnauthorized) WithPayload(payload *models.Error) *DeleteConfigIpv4addressIPAddressUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config ipv4address Ip address unauthorized response
func (o *DeleteConfigIpv4addressIPAddressUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigIpv4addressIPAddressUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigIpv4addressIPAddressForbiddenCode is the HTTP code returned for type DeleteConfigIpv4addressIPAddressForbidden
const DeleteConfigIpv4addressIPAddressForbiddenCode int = 403

/*
DeleteConfigIpv4addressIPAddressForbidden Capacity insufficient

swagger:response deleteConfigIpv4addressIpAddressForbidden
*/
type DeleteConfigIpv4addressIPAddressForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigIpv4addressIPAddressForbidden creates DeleteConfigIpv4addressIPAddressForbidden with default headers values
func NewDeleteConfigIpv4addressIPAddressForbidden() *DeleteConfigIpv4addressIPAddressForbidden {

	return &DeleteConfigIpv4addressIPAddressForbidden{}
}

// WithPayload adds the payload to the delete config ipv4address Ip address forbidden response
func (o *DeleteConfigIpv4addressIPAddressForbidden) WithPayload(payload *models.Error) *DeleteConfigIpv4addressIPAddressForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config ipv4address Ip address forbidden response
func (o *DeleteConfigIpv4addressIPAddressForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigIpv4addressIPAddressForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigIpv4addressIPAddressNotFoundCode is the HTTP code returned for type DeleteConfigIpv4addressIPAddressNotFound
const DeleteConfigIpv4addressIPAddressNotFoundCode int = 404

/*
DeleteConfigIpv4addressIPAddressNotFound Resource not found

swagger:response deleteConfigIpv4addressIpAddressNotFound
*/
type DeleteConfigIpv4addressIPAddressNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigIpv4addressIPAddressNotFound creates DeleteConfigIpv4addressIPAddressNotFound with default headers values
func NewDeleteConfigIpv4addressIPAddressNotFound() *DeleteConfigIpv4addressIPAddressNotFound {

	return &DeleteConfigIpv4addressIPAddressNotFound{}
}

// WithPayload adds the payload to the delete config ipv4address Ip address not found response
func (o *DeleteConfigIpv4addressIPAddressNotFound) WithPayload(payload *models.Error) *DeleteConfigIpv4addressIPAddressNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config ipv4address Ip address not found response
func (o *DeleteConfigIpv4addressIPAddressNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigIpv4addressIPAddressNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigIpv4addressIPAddressConflictCode is the HTTP code returned for type DeleteConfigIpv4addressIPAddressConflict
const DeleteConfigIpv4addressIPAddressConflictCode int = 409

/*
DeleteConfigIpv4addressIPAddressConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response deleteConfigIpv4addressIpAddressConflict
*/
type DeleteConfigIpv4addressIPAddressConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigIpv4addressIPAddressConflict creates DeleteConfigIpv4addressIPAddressConflict with default headers values
func NewDeleteConfigIpv4addressIPAddressConflict() *DeleteConfigIpv4addressIPAddressConflict {

	return &DeleteConfigIpv4addressIPAddressConflict{}
}

// WithPayload adds the payload to the delete config ipv4address Ip address conflict response
func (o *DeleteConfigIpv4addressIPAddressConflict) WithPayload(payload *models.Error) *DeleteConfigIpv4addressIPAddressConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config ipv4address Ip address conflict response
func (o *DeleteConfigIpv4addressIPAddressConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigIpv4addressIPAddressConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigIpv4addressIPAddressInternalServerErrorCode is the HTTP code returned for type DeleteConfigIpv4addressIPAddressInternalServerError
const DeleteConfigIpv4addressIPAddressInternalServerErrorCode int = 500

/*
DeleteConfigIpv4addressIPAddressInternalServerError Internal service error

swagger:response deleteConfigIpv4addressIpAddressInternalServerError
*/
type DeleteConfigIpv4addressIPAddressInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigIpv4addressIPAddressInternalServerError creates DeleteConfigIpv4addressIPAddressInternalServerError with default headers values
func NewDeleteConfigIpv4addressIPAddressInternalServerError() *DeleteConfigIpv4addressIPAddressInternalServerError {

	return &DeleteConfigIpv4addressIPAddressInternalServerError{}
}

// WithPayload adds the payload to the delete config ipv4address Ip address internal server error response
func (o *DeleteConfigIpv4addressIPAddressInternalServerError) WithPayload(payload *models.Error) *DeleteConfigIpv4addressIPAddressInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config ipv4address Ip address internal server error response
func (o *DeleteConfigIpv4addressIPAddressInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigIpv4addressIPAddressInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigIpv4addressIPAddressServiceUnavailableCode is the HTTP code returned for type DeleteConfigIpv4addressIPAddressServiceUnavailable
const DeleteConfigIpv4addressIPAddressServiceUnavailableCode int = 503

/*
DeleteConfigIpv4addressIPAddressServiceUnavailable Maintanence mode

swagger:response deleteConfigIpv4addressIpAddressServiceUnavailable
*/
type DeleteConfigIpv4addressIPAddressServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigIpv4addressIPAddressServiceUnavailable creates DeleteConfigIpv4addressIPAddressServiceUnavailable with default headers values
func NewDeleteConfigIpv4addressIPAddressServiceUnavailable() *DeleteConfigIpv4addressIPAddressServiceUnavailable {

	return &DeleteConfigIpv4addressIPAddressServiceUnavailable{}
}

// WithPayload adds the payload to the delete config ipv4address Ip address service unavailable response
func (o *DeleteConfigIpv4addressIPAddressServiceUnavailable) WithPayload(payload *models.Error) *DeleteConfigIpv4addressIPAddressServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config ipv4address Ip address service unavailable response
func (o *DeleteConfigIpv4addressIPAddressServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigIpv4addressIPAddressServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
