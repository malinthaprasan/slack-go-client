# \AppsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppsPermissionsInfo**](AppsApi.md#AppsPermissionsInfo) | **Get** /apps.permissions.info | 
[**AppsPermissionsRequest**](AppsApi.md#AppsPermissionsRequest) | **Get** /apps.permissions.request | 
[**AppsPermissionsResourcesList**](AppsApi.md#AppsPermissionsResourcesList) | **Get** /apps.permissions.resources.list | 
[**AppsPermissionsScopesList**](AppsApi.md#AppsPermissionsScopesList) | **Get** /apps.permissions.scopes.list | 
[**AppsPermissionsUsersList**](AppsApi.md#AppsPermissionsUsersList) | **Get** /apps.permissions.users.list | 
[**AppsPermissionsUsersRequest**](AppsApi.md#AppsPermissionsUsersRequest) | **Get** /apps.permissions.users.request | 
[**AppsUninstall**](AppsApi.md#AppsUninstall) | **Get** /apps.uninstall | 


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

# **AppsPermissionsResourcesList**
> map[string]interface{} AppsPermissionsResourcesList(ctx, optional)


Returns list of resource grants this app has on a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppsPermissionsResourcesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppsPermissionsResourcesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;none&#x60; | 
 **limit** | **optional.Int32**| The maximum number of items to return. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AppsPermissionsScopesList**
> map[string]interface{} AppsPermissionsScopesList(ctx, optional)


Returns list of scopes this app has on a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppsPermissionsScopesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppsPermissionsScopesListOpts struct

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

# **AppsUninstall**
> map[string]interface{} AppsUninstall(ctx, optional)


Uninstalls your app from a workspace.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AppsUninstallOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AppsUninstallOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientSecret** | **optional.String**| Issued when you created your application. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;none&#x60; | 
 **clientId** | **optional.String**| Issued when you created your application. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

