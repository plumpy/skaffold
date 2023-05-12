package containerregistry

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// RepositoryClient is the metadata API definition for the Azure Container Registry runtime
type RepositoryClient struct {
	BaseClient
}

// NewRepositoryClient creates an instance of the RepositoryClient client.
func NewRepositoryClient(loginURI string) RepositoryClient {
	return RepositoryClient{New(loginURI)}
}

// Delete delete the repository identified by `name`
// Parameters:
// name - name of the image (including the namespace)
func (client RepositoryClient) Delete(ctx context.Context, name string) (result DeletedRepository, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RepositoryClient.Delete")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RepositoryClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.RepositoryClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RepositoryClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client RepositoryClient) DeletePreparer(ctx context.Context, name string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"url": client.LoginURI,
	}

	pathParameters := map[string]interface{}{
		"name": autorest.Encode("path", name),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPathParameters("/acr/v1/{name}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client RepositoryClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client RepositoryClient) DeleteResponder(resp *http.Response) (result DeletedRepository, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetAttributes get repository attributes
// Parameters:
// name - name of the image (including the namespace)
func (client RepositoryClient) GetAttributes(ctx context.Context, name string) (result RepositoryAttributes, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RepositoryClient.GetAttributes")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetAttributesPreparer(ctx, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RepositoryClient", "GetAttributes", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAttributesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.RepositoryClient", "GetAttributes", resp, "Failure sending request")
		return
	}

	result, err = client.GetAttributesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RepositoryClient", "GetAttributes", resp, "Failure responding to request")
		return
	}

	return
}

// GetAttributesPreparer prepares the GetAttributes request.
func (client RepositoryClient) GetAttributesPreparer(ctx context.Context, name string) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"url": client.LoginURI,
	}

	pathParameters := map[string]interface{}{
		"name": autorest.Encode("path", name),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPathParameters("/acr/v1/{name}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetAttributesSender sends the GetAttributes request. The method will close the
// http.Response Body if it receives an error.
func (client RepositoryClient) GetAttributesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetAttributesResponder handles the response to the GetAttributes request. The method always
// closes the http.Response Body.
func (client RepositoryClient) GetAttributesResponder(resp *http.Response) (result RepositoryAttributes, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetList list repositories
// Parameters:
// last - query parameter for the last item in previous query. Result set will include values lexically after
// last.
// n - query parameter for max number of items
func (client RepositoryClient) GetList(ctx context.Context, last string, n *int32) (result Repositories, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RepositoryClient.GetList")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetListPreparer(ctx, last, n)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RepositoryClient", "GetList", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerregistry.RepositoryClient", "GetList", resp, "Failure sending request")
		return
	}

	result, err = client.GetListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RepositoryClient", "GetList", resp, "Failure responding to request")
		return
	}

	return
}

// GetListPreparer prepares the GetList request.
func (client RepositoryClient) GetListPreparer(ctx context.Context, last string, n *int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"url": client.LoginURI,
	}

	queryParameters := map[string]interface{}{}
	if len(last) > 0 {
		queryParameters["last"] = autorest.Encode("query", last)
	}
	if n != nil {
		queryParameters["n"] = autorest.Encode("query", *n)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPath("/acr/v1/_catalog"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetListSender sends the GetList request. The method will close the
// http.Response Body if it receives an error.
func (client RepositoryClient) GetListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetListResponder handles the response to the GetList request. The method always
// closes the http.Response Body.
func (client RepositoryClient) GetListResponder(resp *http.Response) (result Repositories, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// UpdateAttributes update the attribute identified by `name` where `reference` is the name of the repository.
// Parameters:
// name - name of the image (including the namespace)
// value - repository attribute value
func (client RepositoryClient) UpdateAttributes(ctx context.Context, name string, value *ChangeableAttributes) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/RepositoryClient.UpdateAttributes")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdateAttributesPreparer(ctx, name, value)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RepositoryClient", "UpdateAttributes", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateAttributesSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "containerregistry.RepositoryClient", "UpdateAttributes", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateAttributesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerregistry.RepositoryClient", "UpdateAttributes", resp, "Failure responding to request")
		return
	}

	return
}

// UpdateAttributesPreparer prepares the UpdateAttributes request.
func (client RepositoryClient) UpdateAttributesPreparer(ctx context.Context, name string, value *ChangeableAttributes) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"url": client.LoginURI,
	}

	pathParameters := map[string]interface{}{
		"name": autorest.Encode("path", name),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithCustomBaseURL("{url}", urlParameters),
		autorest.WithPathParameters("/acr/v1/{name}", pathParameters))
	if value != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(value))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateAttributesSender sends the UpdateAttributes request. The method will close the
// http.Response Body if it receives an error.
func (client RepositoryClient) UpdateAttributesSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// UpdateAttributesResponder handles the response to the UpdateAttributes request. The method always
// closes the http.Response Body.
func (client RepositoryClient) UpdateAttributesResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}