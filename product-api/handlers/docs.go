// Package classification of Product API
//
// Documentation for Product API
//
// 		Schemes: http
// 		BasePath: /
// 		Version: 1.0.0
//
// 		Consumes:
// 		- application/json
//
// 		Produces:
// 		- application/json
//
// swagger:meta
package handlers

import "working/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of products returns on the the response
// swagger:response productsResponse
type productsResponseWrapper struct {
	// All products in the system
	// in: body
	Body []data.Product
}

// Data structure representing a single product
// swagger:response productResponse
type productResponseWrapper struct {
	// Newly created product
	// in: body
	Body data.Product
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int `json:"id"`
}

// swagger:parameters updateProduct
type productIDParamsWrapper struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}

// No response
// swagger:response noContent
type productsNoContent struct{}
