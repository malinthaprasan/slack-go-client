# \AppsPermissionsUsersApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsPermissionsUsersList**](AppsPermissionsUsersApi.md#AppsPermissionsUsersList) | **Get** /apps.permissions.users.list | 
[**AppsPermissionsUsersRequest**](AppsPermissionsUsersApi.md#AppsPermissionsUsersRequest) | **Get** /apps.permissions.users.request | 


# **AppsPermissionsUsersList**
> DefaultSuccessTemplate AppsPermissionsUsersList(ctx, optional)


Returns list of user grants and corresponding scopes this app has on a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppsPermissionsUsersListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppsPermissionsUsersListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;none&#x60; | 
 **limit** | **optional.Int32**| The maximum number of items to return. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppsPermissionsUsersRequest**
> DefaultSuccessTemplate AppsPermissionsUsersRequest(ctx, optional)


Enables an app to trigger a permissions modal to grant an app access to a user access scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppsPermissionsUsersRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppsPermissionsUsersRequestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scopes** | **optional.String**| A comma separated list of user scopes to request for | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;none&#x60; | 
 **user** | **optional.String**| The user this scope is being requested for | 
 **triggerId** | **optional.String**| Token used to trigger the request | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

