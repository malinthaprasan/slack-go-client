/*
 * Slack Web API
 *
 * One way to interact with the Slack platform is its HTTP RPC-based Web API, a collection of methods requiring OAuth 2.0-based user, bot, or workspace tokens blessed with related OAuth scopes.
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type OauthApiService service

/* 
OauthApiService
Exchanges a temporary OAuth verifier code for an access token.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OauthAccessOpts - Optional Parameters:
     * @param "Code" (optional.String) -  The &#x60;code&#x60; param returned via the OAuth callback.
     * @param "RedirectUri" (optional.String) -  This must match the originally submitted URI (if one was sent).
     * @param "ClientId" (optional.String) -  Issued when you created your application.
     * @param "ClientSecret" (optional.String) -  Issued when you created your application.
     * @param "SingleChannel" (optional.Bool) -  Request the user to add your app only to a single channel.

@return DefaultSuccessTemplate
*/

type OauthAccessOpts struct { 
	Code optional.String
	RedirectUri optional.String
	ClientId optional.String
	ClientSecret optional.String
	SingleChannel optional.Bool
}

func (a *OauthApiService) OauthAccess(ctx context.Context, localVarOptionals *OauthAccessOpts) (DefaultSuccessTemplate, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue DefaultSuccessTemplate
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth.access"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Code.IsSet() {
		localVarQueryParams.Add("code", parameterToString(localVarOptionals.Code.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RedirectUri.IsSet() {
		localVarQueryParams.Add("redirect_uri", parameterToString(localVarOptionals.RedirectUri.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientId.IsSet() {
		localVarQueryParams.Add("client_id", parameterToString(localVarOptionals.ClientId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientSecret.IsSet() {
		localVarQueryParams.Add("client_secret", parameterToString(localVarOptionals.ClientSecret.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SingleChannel.IsSet() {
		localVarQueryParams.Add("single_channel", parameterToString(localVarOptionals.SingleChannel.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v DefaultSuccessTemplate
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v DefaultSuccessTemplate
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/* 
OauthApiService
Exchanges a temporary OAuth verifier code for a workspace token.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *OauthTokenOpts - Optional Parameters:
     * @param "ClientSecret" (optional.String) -  Issued when you created your application.
     * @param "Code" (optional.String) -  The &#x60;code&#x60; param returned via the OAuth callback.
     * @param "SingleChannel" (optional.Bool) -  Request the user to add your app only to a single channel.
     * @param "ClientId" (optional.String) -  Issued when you created your application.
     * @param "RedirectUri" (optional.String) -  This must match the originally submitted URI (if one was sent).

@return DefaultSuccessTemplate
*/

type OauthTokenOpts struct { 
	ClientSecret optional.String
	Code optional.String
	SingleChannel optional.Bool
	ClientId optional.String
	RedirectUri optional.String
}

func (a *OauthApiService) OauthToken(ctx context.Context, localVarOptionals *OauthTokenOpts) (DefaultSuccessTemplate, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue DefaultSuccessTemplate
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth.token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ClientSecret.IsSet() {
		localVarQueryParams.Add("client_secret", parameterToString(localVarOptionals.ClientSecret.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Code.IsSet() {
		localVarQueryParams.Add("code", parameterToString(localVarOptionals.Code.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SingleChannel.IsSet() {
		localVarQueryParams.Add("single_channel", parameterToString(localVarOptionals.SingleChannel.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientId.IsSet() {
		localVarQueryParams.Add("client_id", parameterToString(localVarOptionals.ClientId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RedirectUri.IsSet() {
		localVarQueryParams.Add("redirect_uri", parameterToString(localVarOptionals.RedirectUri.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		
		if localVarHttpResponse.StatusCode == 200 {
			var v DefaultSuccessTemplate
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v DefaultSuccessTemplate
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
