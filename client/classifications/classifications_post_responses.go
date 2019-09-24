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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// ClassificationsPostReader is a Reader for the ClassificationsPost structure.
type ClassificationsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClassificationsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewClassificationsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewClassificationsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewClassificationsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewClassificationsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewClassificationsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewClassificationsPostCreated creates a ClassificationsPostCreated with default headers values
func NewClassificationsPostCreated() *ClassificationsPostCreated {
	return &ClassificationsPostCreated{}
}

/*ClassificationsPostCreated handles this case with default header values.

Successfully started classification.
*/
type ClassificationsPostCreated struct {
	Payload *models.Classification
}

func (o *ClassificationsPostCreated) Error() string {
	return fmt.Sprintf("[POST /classifications/][%d] classificationsPostCreated  %+v", 201, o.Payload)
}

func (o *ClassificationsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Classification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClassificationsPostBadRequest creates a ClassificationsPostBadRequest with default headers values
func NewClassificationsPostBadRequest() *ClassificationsPostBadRequest {
	return &ClassificationsPostBadRequest{}
}

/*ClassificationsPostBadRequest handles this case with default header values.

Incorrect request
*/
type ClassificationsPostBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ClassificationsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /classifications/][%d] classificationsPostBadRequest  %+v", 400, o.Payload)
}

func (o *ClassificationsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClassificationsPostUnauthorized creates a ClassificationsPostUnauthorized with default headers values
func NewClassificationsPostUnauthorized() *ClassificationsPostUnauthorized {
	return &ClassificationsPostUnauthorized{}
}

/*ClassificationsPostUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ClassificationsPostUnauthorized struct {
}

func (o *ClassificationsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /classifications/][%d] classificationsPostUnauthorized ", 401)
}

func (o *ClassificationsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewClassificationsPostForbidden creates a ClassificationsPostForbidden with default headers values
func NewClassificationsPostForbidden() *ClassificationsPostForbidden {
	return &ClassificationsPostForbidden{}
}

/*ClassificationsPostForbidden handles this case with default header values.

Forbidden
*/
type ClassificationsPostForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ClassificationsPostForbidden) Error() string {
	return fmt.Sprintf("[POST /classifications/][%d] classificationsPostForbidden  %+v", 403, o.Payload)
}

func (o *ClassificationsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClassificationsPostInternalServerError creates a ClassificationsPostInternalServerError with default headers values
func NewClassificationsPostInternalServerError() *ClassificationsPostInternalServerError {
	return &ClassificationsPostInternalServerError{}
}

/*ClassificationsPostInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ClassificationsPostInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ClassificationsPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /classifications/][%d] classificationsPostInternalServerError  %+v", 500, o.Payload)
}

func (o *ClassificationsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
