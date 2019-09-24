//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 SeMI Holding B.V. (registered @ Dutch Chamber of Commerce no 75221632). All rights reserved.
//  LICENSE WEAVIATE OPEN SOURCE: https://www.semi.technology/playbook/playbook/contract-weaviate-OSS.html
//  LICENSE WEAVIATE ENTERPRISE: https://www.semi.technology/playbook/contract-weaviate-enterprise.html
//  CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package classifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ClassificationsGetOKCode is the HTTP code returned for type ClassificationsGetOK
const ClassificationsGetOKCode int = 200

/*ClassificationsGetOK Found the classification, returned as body

swagger:response classificationsGetOK
*/
type ClassificationsGetOK struct {

	/*
	  In: Body
	*/
	Payload *models.Classification `json:"body,omitempty"`
}

// NewClassificationsGetOK creates ClassificationsGetOK with default headers values
func NewClassificationsGetOK() *ClassificationsGetOK {

	return &ClassificationsGetOK{}
}

// WithPayload adds the payload to the classifications get o k response
func (o *ClassificationsGetOK) WithPayload(payload *models.Classification) *ClassificationsGetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the classifications get o k response
func (o *ClassificationsGetOK) SetPayload(payload *models.Classification) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClassificationsGetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ClassificationsGetUnauthorizedCode is the HTTP code returned for type ClassificationsGetUnauthorized
const ClassificationsGetUnauthorizedCode int = 401

/*ClassificationsGetUnauthorized Unauthorized or invalid credentials.

swagger:response classificationsGetUnauthorized
*/
type ClassificationsGetUnauthorized struct {
}

// NewClassificationsGetUnauthorized creates ClassificationsGetUnauthorized with default headers values
func NewClassificationsGetUnauthorized() *ClassificationsGetUnauthorized {

	return &ClassificationsGetUnauthorized{}
}

// WriteResponse to the client
func (o *ClassificationsGetUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// ClassificationsGetForbiddenCode is the HTTP code returned for type ClassificationsGetForbidden
const ClassificationsGetForbiddenCode int = 403

/*ClassificationsGetForbidden Forbidden

swagger:response classificationsGetForbidden
*/
type ClassificationsGetForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewClassificationsGetForbidden creates ClassificationsGetForbidden with default headers values
func NewClassificationsGetForbidden() *ClassificationsGetForbidden {

	return &ClassificationsGetForbidden{}
}

// WithPayload adds the payload to the classifications get forbidden response
func (o *ClassificationsGetForbidden) WithPayload(payload *models.ErrorResponse) *ClassificationsGetForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the classifications get forbidden response
func (o *ClassificationsGetForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClassificationsGetForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ClassificationsGetNotFoundCode is the HTTP code returned for type ClassificationsGetNotFound
const ClassificationsGetNotFoundCode int = 404

/*ClassificationsGetNotFound Not Found - Classification does not exist

swagger:response classificationsGetNotFound
*/
type ClassificationsGetNotFound struct {
}

// NewClassificationsGetNotFound creates ClassificationsGetNotFound with default headers values
func NewClassificationsGetNotFound() *ClassificationsGetNotFound {

	return &ClassificationsGetNotFound{}
}

// WriteResponse to the client
func (o *ClassificationsGetNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ClassificationsGetInternalServerErrorCode is the HTTP code returned for type ClassificationsGetInternalServerError
const ClassificationsGetInternalServerErrorCode int = 500

/*ClassificationsGetInternalServerError An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.

swagger:response classificationsGetInternalServerError
*/
type ClassificationsGetInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewClassificationsGetInternalServerError creates ClassificationsGetInternalServerError with default headers values
func NewClassificationsGetInternalServerError() *ClassificationsGetInternalServerError {

	return &ClassificationsGetInternalServerError{}
}

// WithPayload adds the payload to the classifications get internal server error response
func (o *ClassificationsGetInternalServerError) WithPayload(payload *models.ErrorResponse) *ClassificationsGetInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the classifications get internal server error response
func (o *ClassificationsGetInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ClassificationsGetInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
