# \UsergroupsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsergroupsCreate**](UsergroupsApi.md#UsergroupsCreate) | **Post** /usergroups.create | 
[**UsergroupsDisable**](UsergroupsApi.md#UsergroupsDisable) | **Post** /usergroups.disable | 
[**UsergroupsEnable**](UsergroupsApi.md#UsergroupsEnable) | **Post** /usergroups.enable | 
[**UsergroupsList**](UsergroupsApi.md#UsergroupsList) | **Get** /usergroups.list | 
[**UsergroupsUpdate**](UsergroupsApi.md#UsergroupsUpdate) | **Post** /usergroups.update | 
[**UsergroupsUsersList**](UsergroupsApi.md#UsergroupsUsersList) | **Get** /usergroups.users.list | 
[**UsergroupsUsersUpdate**](UsergroupsApi.md#UsergroupsUsersUpdate) | **Post** /usergroups.users.update | 


# **UsergroupsCreate**
> map[string]interface{} UsergroupsCreate(ctx, optional)


Create a User Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsergroupsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **handle** | **optional.String**| A mention handle. Must be unique among channels, users and User Groups. | 
 **description** | **optional.String**| A short description of the User Group. | 
 **channels** | **optional.String**| A comma separated string of encoded channel IDs for which the User Group uses as a default. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;usergroups:write&#x60; | 
 **includeCount** | **optional.Bool**| Include the number of users in each User Group. | 
 **name** | **optional.String**| A name for the User Group. Must be unique among User Groups. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsDisable**
> map[string]interface{} UsergroupsDisable(ctx, optional)


Disable an existing User Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsergroupsDisableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupsDisableOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;usergroups:write&#x60; | 
 **includeCount** | **optional.Bool**| Include the number of users in the User Group. | 
 **usergroup** | **optional.String**| The encoded ID of the User Group to disable. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsEnable**
> map[string]interface{} UsergroupsEnable(ctx, optional)


Enable a User Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsergroupsEnableOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupsEnableOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;usergroups:write&#x60; | 
 **includeCount** | **optional.Bool**| Include the number of users in the User Group. | 
 **usergroup** | **optional.String**| The encoded ID of the User Group to enable. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsList**
> map[string]interface{} UsergroupsList(ctx, optional)


List all User Groups for a team

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsergroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeUsers** | **optional.Bool**| Include the list of users for each User Group. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;usergroups:read&#x60; | 
 **includeCount** | **optional.Bool**| Include the number of users in each User Group. | 
 **includeDisabled** | **optional.Bool**| Include disabled User Groups. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsUpdate**
> map[string]interface{} UsergroupsUpdate(ctx, optional)


Update an existing User Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsergroupsUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupsUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **handle** | **optional.String**| A mention handle. Must be unique among channels, users and User Groups. | 
 **description** | **optional.String**| A short description of the User Group. | 
 **channels** | **optional.String**| A comma separated string of encoded channel IDs for which the User Group uses as a default. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;usergroups:write&#x60; | 
 **includeCount** | **optional.Bool**| Include the number of users in the User Group. | 
 **usergroup** | **optional.String**| The encoded ID of the User Group to update. | 
 **name** | **optional.String**| A name for the User Group. Must be unique among User Groups. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsUsersList**
> map[string]interface{} UsergroupsUsersList(ctx, )


List all users in a User Group

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsUsersUpdate**
> map[string]interface{} UsergroupsUsersUpdate(ctx, optional)


Update the list of users for a User Group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsergroupsUsersUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupsUsersUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **users** | **optional.String**| A comma separated string of encoded user IDs that represent the entire list of users for the User Group. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;usergroups:write&#x60; | 
 **includeCount** | **optional.Bool**| Include the number of users in the User Group. | 
 **usergroup** | **optional.String**| The encoded ID of the User Group to update. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

