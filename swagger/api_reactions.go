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

type ReactionsApiService service

/* 
ReactionsApiService
Adds a reaction to an item.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ReactionsAddOpts - Optional Parameters:
     * @param "Name" (optional.String) -  Reaction (emoji) name.
     * @param "FileComment" (optional.String) -  File comment to add reaction to. Now that [file threads](/changelog/2018-05-file-threads-soon-tread#whats_changed) work the way you&#39;d expect, this argument is deprecated. Specify the timestamp and channel of the message associated with a file instead.
     * @param "Timestamp" (optional.Float32) -  Timestamp of the message to add reaction to.
     * @param "Token" (optional.String) -  Authentication token. Requires scope: &#x60;reactions:write&#x60;
     * @param "File" (optional.String) -  File to add reaction to. Now that [file threads](/changelog/2018-05-file-threads-soon-tread#whats_changed) work the way you&#39;d expect, this argument is deprecated. Specify the timestamp and channel of the message associated with a file instead.
     * @param "Channel" (optional.String) -  Channel where the message to add reaction to was posted.

@return map[string]interface{}
*/

type ReactionsAddOpts struct { 
	Name optional.String
	FileComment optional.String
	Timestamp optional.Float32
	Token optional.String
	File optional.String
	Channel optional.String
}

func (a *ReactionsApiService) ReactionsAdd(ctx context.Context, localVarOptionals *ReactionsAddOpts) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue map[string]interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/reactions.add"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded", "application/json"}

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
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarHeaderParams["token"] = parameterToString(localVarOptionals.Token.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.Name.IsSet() {
		localVarFormParams.Add("name", parameterToString(localVarOptionals.Name.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileComment.IsSet() {
		localVarFormParams.Add("file_comment", parameterToString(localVarOptionals.FileComment.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Timestamp.IsSet() {
		localVarFormParams.Add("timestamp", parameterToString(localVarOptionals.Timestamp.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.File.IsSet() {
		localVarFormParams.Add("file", parameterToString(localVarOptionals.File.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Channel.IsSet() {
		localVarFormParams.Add("channel", parameterToString(localVarOptionals.Channel.Value(), ""))
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
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v map[string]interface{}
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
ReactionsApiService
Gets reactions for an item.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ReactionsGetOpts - Optional Parameters:
     * @param "Full" (optional.Bool) -  If true always return the complete reaction list.
     * @param "FileComment" (optional.String) -  File comment to get reactions for.
     * @param "Timestamp" (optional.Float32) -  Timestamp of the message to get reactions for.
     * @param "Token" (optional.String) -  Authentication token. Requires scope: &#x60;reactions:read&#x60;
     * @param "File" (optional.String) -  File to get reactions for.
     * @param "Channel" (optional.String) -  Channel where the message to get reactions for was posted.

@return interface{}
*/

type ReactionsGetOpts struct { 
	Full optional.Bool
	FileComment optional.String
	Timestamp optional.Float32
	Token optional.String
	File optional.String
	Channel optional.String
}

func (a *ReactionsApiService) ReactionsGet(ctx context.Context, localVarOptionals *ReactionsGetOpts) (interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/reactions.get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Full.IsSet() {
		localVarQueryParams.Add("full", parameterToString(localVarOptionals.Full.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileComment.IsSet() {
		localVarQueryParams.Add("file_comment", parameterToString(localVarOptionals.FileComment.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Timestamp.IsSet() {
		localVarQueryParams.Add("timestamp", parameterToString(localVarOptionals.Timestamp.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.File.IsSet() {
		localVarQueryParams.Add("file", parameterToString(localVarOptionals.File.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Channel.IsSet() {
		localVarQueryParams.Add("channel", parameterToString(localVarOptionals.Channel.Value(), ""))
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
			var v interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v map[string]interface{}
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
ReactionsApiService
Lists reactions made by a user.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ReactionsListOpts - Optional Parameters:
     * @param "Count" (optional.String) - 
     * @param "Full" (optional.Bool) -  If true always return the complete reaction list.
     * @param "Cursor" (optional.String) -  Parameter for pagination. Set &#x60;cursor&#x60; equal to the &#x60;next_cursor&#x60; attribute returned by the previous request&#39;s &#x60;response_metadata&#x60;. This parameter is optional, but pagination is mandatory: the default value simply fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more details.
     * @param "Token" (optional.String) -  Authentication token. Requires scope: &#x60;reactions:read&#x60;
     * @param "Limit" (optional.Int32) -  The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the list hasn&#39;t been reached.
     * @param "User" (optional.String) -  Show reactions made by this user. Defaults to the authed user.
     * @param "Page" (optional.String) - 

@return map[string]interface{}
*/

type ReactionsListOpts struct { 
	Count optional.String
	Full optional.Bool
	Cursor optional.String
	Token optional.String
	Limit optional.Int32
	User optional.String
	Page optional.String
}

func (a *ReactionsApiService) ReactionsList(ctx context.Context, localVarOptionals *ReactionsListOpts) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue map[string]interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/reactions.list"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Full.IsSet() {
		localVarQueryParams.Add("full", parameterToString(localVarOptionals.Full.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Cursor.IsSet() {
		localVarQueryParams.Add("cursor", parameterToString(localVarOptionals.Cursor.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarQueryParams.Add("token", parameterToString(localVarOptionals.Token.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.User.IsSet() {
		localVarQueryParams.Add("user", parameterToString(localVarOptionals.User.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
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
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v map[string]interface{}
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
ReactionsApiService
Removes a reaction from an item.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ReactionsRemoveOpts - Optional Parameters:
     * @param "Name" (optional.String) -  Reaction (emoji) name.
     * @param "FileComment" (optional.String) -  File comment to remove reaction from.
     * @param "Timestamp" (optional.Float32) -  Timestamp of the message to remove reaction from.
     * @param "Token" (optional.String) -  Authentication token. Requires scope: &#x60;reactions:write&#x60;
     * @param "File" (optional.String) -  File to remove reaction from.
     * @param "Channel" (optional.String) -  Channel where the message to remove reaction from was posted.

@return map[string]interface{}
*/

type ReactionsRemoveOpts struct { 
	Name optional.String
	FileComment optional.String
	Timestamp optional.Float32
	Token optional.String
	File optional.String
	Channel optional.String
}

func (a *ReactionsApiService) ReactionsRemove(ctx context.Context, localVarOptionals *ReactionsRemoveOpts) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue map[string]interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/reactions.remove"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded", "application/json"}

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
	if localVarOptionals != nil && localVarOptionals.Token.IsSet() {
		localVarHeaderParams["token"] = parameterToString(localVarOptionals.Token.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.Name.IsSet() {
		localVarFormParams.Add("name", parameterToString(localVarOptionals.Name.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.FileComment.IsSet() {
		localVarFormParams.Add("file_comment", parameterToString(localVarOptionals.FileComment.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Timestamp.IsSet() {
		localVarFormParams.Add("timestamp", parameterToString(localVarOptionals.Timestamp.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.File.IsSet() {
		localVarFormParams.Add("file", parameterToString(localVarOptionals.File.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Channel.IsSet() {
		localVarFormParams.Add("channel", parameterToString(localVarOptionals.Channel.Value(), ""))
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
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		
		if localVarHttpResponse.StatusCode == 0 {
			var v map[string]interface{}
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
