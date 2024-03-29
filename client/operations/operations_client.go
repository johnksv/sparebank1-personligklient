// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAccount finds the account with the given id returns account with details

Will not accept account numbers, the id must be used. If unknown the id can be accessed by getting all accounts and following the links.
*/
func (a *Client) GetAccount(params *GetAccountParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccount",
		Method:             "GET",
		PathPattern:        "/{id}",
		ProducesMediaTypes: []string{"application/vnd.sparebank1.v1+json; charset=utf-8"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountOK), nil

}

/*
GetAccountTransactions reads account transactions data from a given account
*/
func (a *Client) GetAccountTransactions(params *GetAccountTransactionsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountTransactionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountTransactionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccountTransactions",
		Method:             "GET",
		PathPattern:        "/{accountId}/transactions",
		ProducesMediaTypes: []string{"application/vnd.sparebank1.v1+json; charset=utf-8"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountTransactionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountTransactionsOK), nil

}

/*
GetAccounts finds all banking accounts
*/
func (a *Client) GetAccounts(params *GetAccountsParams, authInfo runtime.ClientAuthInfoWriter) (*GetAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAccountsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAccounts",
		Method:             "GET",
		PathPattern:        "/all",
		ProducesMediaTypes: []string{"application/vnd.sparebank1.v1+json; charset=utf-8"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAccountsOK), nil

}

/*
GetDefaultBalanceAccount finds the default account

Returns the default payment account of the logged in user, if no default payment account is found the users first account is returned.
*/
func (a *Client) GetDefaultBalanceAccount(params *GetDefaultBalanceAccountParams, authInfo runtime.ClientAuthInfoWriter) (*GetDefaultBalanceAccountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDefaultBalanceAccountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDefaultBalanceAccount",
		Method:             "GET",
		PathPattern:        "/default",
		ProducesMediaTypes: []string{"application/vnd.sparebank1.v1+json; charset=utf-8"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDefaultBalanceAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDefaultBalanceAccountOK), nil

}

/*
GetRoot roots API
*/
func (a *Client) GetRoot(params *GetRootParams, authInfo runtime.ClientAuthInfoWriter) (*GetRootOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRootParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRoot",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/vnd.sparebank1.v1+json; charset=utf-8"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRootReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRootOK), nil

}

/*
GetTransactionDetails reads account transaction details for a given transaction
*/
func (a *Client) GetTransactionDetails(params *GetTransactionDetailsParams, authInfo runtime.ClientAuthInfoWriter) (*GetTransactionDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTransactionDetailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTransactionDetails",
		Method:             "GET",
		PathPattern:        "/{accountId}/transactions/{transactionId}",
		ProducesMediaTypes: []string{"application/vnd.sparebank1.v1+json; charset=utf-8"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTransactionDetailsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTransactionDetailsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
