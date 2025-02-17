// Code generated by go-swagger; DO NOT EDIT.

package share

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UpdateShareOKCode is the HTTP code returned for type UpdateShareOK
const UpdateShareOKCode int = 200

/*
UpdateShareOK share updated

swagger:response updateShareOK
*/
type UpdateShareOK struct {
}

// NewUpdateShareOK creates UpdateShareOK with default headers values
func NewUpdateShareOK() *UpdateShareOK {

	return &UpdateShareOK{}
}

// WriteResponse to the client
func (o *UpdateShareOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UpdateShareUnauthorizedCode is the HTTP code returned for type UpdateShareUnauthorized
const UpdateShareUnauthorizedCode int = 401

/*
UpdateShareUnauthorized unauthorized

swagger:response updateShareUnauthorized
*/
type UpdateShareUnauthorized struct {
}

// NewUpdateShareUnauthorized creates UpdateShareUnauthorized with default headers values
func NewUpdateShareUnauthorized() *UpdateShareUnauthorized {

	return &UpdateShareUnauthorized{}
}

// WriteResponse to the client
func (o *UpdateShareUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// UpdateShareNotFoundCode is the HTTP code returned for type UpdateShareNotFound
const UpdateShareNotFoundCode int = 404

/*
UpdateShareNotFound not found

swagger:response updateShareNotFound
*/
type UpdateShareNotFound struct {
}

// NewUpdateShareNotFound creates UpdateShareNotFound with default headers values
func NewUpdateShareNotFound() *UpdateShareNotFound {

	return &UpdateShareNotFound{}
}

// WriteResponse to the client
func (o *UpdateShareNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// UpdateShareInternalServerErrorCode is the HTTP code returned for type UpdateShareInternalServerError
const UpdateShareInternalServerErrorCode int = 500

/*
UpdateShareInternalServerError internal server error

swagger:response updateShareInternalServerError
*/
type UpdateShareInternalServerError struct {
}

// NewUpdateShareInternalServerError creates UpdateShareInternalServerError with default headers values
func NewUpdateShareInternalServerError() *UpdateShareInternalServerError {

	return &UpdateShareInternalServerError{}
}

// WriteResponse to the client
func (o *UpdateShareInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
