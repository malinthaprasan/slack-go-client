# \StarsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StarsAdd**](StarsApi.md#StarsAdd) | **Post** /stars.add | 
[**StarsList**](StarsApi.md#StarsList) | **Get** /stars.list | 
[**StarsRemove**](StarsApi.md#StarsRemove) | **Post** /stars.remove | 


# **StarsAdd**
> map[string]interface{} StarsAdd(ctx, optional)


Adds a star to an item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StarsAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StarsAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileComment** | **optional.String**| File comment to add star to. | 
 **timestamp** | **optional.Float32**| Timestamp of the message to add star to. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;stars:write&#x60; | 
 **channel** | **optional.String**| Channel to add star to, or channel where the message to add star to was posted (used with &#x60;timestamp&#x60;). | 
 **file** | **optional.String**| File to add star to. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StarsList**
> map[string]interface{} StarsList(ctx, optional)


Lists stars for a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StarsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StarsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.String**|  | 
 **cursor** | **optional.String**| Parameter for pagination. Set &#x60;cursor&#x60; equal to the &#x60;next_cursor&#x60; attribute returned by the previous request&#39;s &#x60;response_metadata&#x60;. This parameter is optional, but pagination is mandatory: the default value simply fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more details. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;stars:read&#x60; | 
 **limit** | **optional.Int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the list hasn&#39;t been reached. | 
 **page** | **optional.String**|  | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StarsRemove**
> map[string]interface{} StarsRemove(ctx, optional)


Removes a star from an item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StarsRemoveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StarsRemoveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileComment** | **optional.String**| File comment to remove star from. | 
 **timestamp** | **optional.Float32**| Timestamp of the message to remove star from. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;stars:write&#x60; | 
 **channel** | **optional.String**| Channel to remove star from, or channel where the message to remove star from was posted (used with &#x60;timestamp&#x60;). | 
 **file** | **optional.String**| File to remove star from. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

