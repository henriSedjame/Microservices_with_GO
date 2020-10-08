// Code generated by go-swagger; DO NOT EDIT.

package update_product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new update product API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for update product API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	UpdateProduct(params *UpdateProductParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateProductOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  UpdateProduct Update a product
*/
func (a *Client) UpdateProduct(params *UpdateProductParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateProductOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateProductParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateProduct",
		Method:             "PUT",
		PathPattern:        "/products",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateProductReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateProductOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateProduct: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
