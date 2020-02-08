# \PinsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PinsAdd**](PinsApi.md#PinsAdd) | **Post** /pins.add | 
[**PinsList**](PinsApi.md#PinsList) | **Get** /pins.list | 
[**PinsRemove**](PinsApi.md#PinsRemove) | **Post** /pins.remove | 


# **PinsAdd**
> map[string]interface{} PinsAdd(ctx, optional)


Pins an item to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PinsAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PinsAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileComment** | **optional.String**| File comment to pin. | 
 **timestamp** | **optional.Float32**| Timestamp of the message to pin. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;pins:write&#x60; | 
 **file** | **optional.String**| File to pin. | 
 **channel** | **optional.String**| Channel to pin the item in. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PinsList**
> interface{} PinsList(ctx, optional)


Lists items pinned to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PinsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PinsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;pins:read&#x60; | 
 **channel** | **optional.String**| Channel to get pinned items for. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PinsRemove**
> map[string]interface{} PinsRemove(ctx, optional)


Un-pins an item from a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PinsRemoveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PinsRemoveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fileComment** | **optional.String**| File comment to un-pin. | 
 **timestamp** | **optional.Float32**| Timestamp of the message to un-pin. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;pins:write&#x60; | 
 **file** | **optional.String**| File to un-pin. | 
 **channel** | **optional.String**| Channel where the item is pinned to. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

