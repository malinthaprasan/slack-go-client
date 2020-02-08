# \AppsPermissionsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsPermissionsInfo**](AppsPermissionsApi.md#AppsPermissionsInfo) | **Get** /apps.permissions.info | 
[**AppsPermissionsRequest**](AppsPermissionsApi.md#AppsPermissionsRequest) | **Get** /apps.permissions.request | 


# **AppsPermissionsInfo**
> map[string]interface{} AppsPermissionsInfo(ctx, optional)


Returns list of permissions this app has on a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppsPermissionsInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppsPermissionsInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;none&#x60; | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppsPermissionsRequest**
> map[string]interface{} AppsPermissionsRequest(ctx, optional)


Allows an app to request additional scopes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppsPermissionsRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppsPermissionsRequestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scopes** | **optional.String**| A comma separated list of scopes to request for | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;none&#x60; | 
 **triggerId** | **optional.String**| Token used to trigger the permissions API | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

