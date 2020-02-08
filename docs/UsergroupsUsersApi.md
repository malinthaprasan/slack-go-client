# \UsergroupsUsersApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsergroupsUsersList**](UsergroupsUsersApi.md#UsergroupsUsersList) | **Get** /usergroups.users.list | 
[**UsergroupsUsersUpdate**](UsergroupsUsersApi.md#UsergroupsUsersUpdate) | **Post** /usergroups.users.update | 


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

