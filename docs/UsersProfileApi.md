# \UsersProfileApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersProfileGet**](UsersProfileApi.md#UsersProfileGet) | **Get** /users.profile.get | 
[**UsersProfileSet**](UsersProfileApi.md#UsersProfileSet) | **Post** /users.profile.set | 


# **UsersProfileGet**
> map[string]interface{} UsersProfileGet(ctx, optional)


Retrieves a user's profile information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersProfileGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersProfileGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;users.profile:read&#x60; | 
 **includeLabels** | **optional.Bool**| Include labels for each ID in custom profile fields | 
 **user** | **optional.String**| User to retrieve profile info for | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersProfileSet**
> map[string]interface{} UsersProfileSet(ctx, optional)


Set the profile information for a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersProfileSetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersProfileSetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **profile** | **optional.String**| Collection of key:value pairs presented as a URL-encoded JSON hash. At most 50 fields may be set. Each field name is limited to 255 characters. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;users.profile:write&#x60; | 
 **user** | **optional.String**| ID of user to change. This argument may only be specified by team admins on paid teams. | 
 **value** | **optional.String**| Value to set a single key to. Usable only if &#x60;profile&#x60; is not passed. | 
 **name** | **optional.String**| Name of a single key to set. Usable only if &#x60;profile&#x60; is not passed. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

