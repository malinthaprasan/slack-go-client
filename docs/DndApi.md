# \DndApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DndEndDnd**](DndApi.md#DndEndDnd) | **Post** /dnd.endDnd | 
[**DndEndSnooze**](DndApi.md#DndEndSnooze) | **Post** /dnd.endSnooze | 
[**DndInfo**](DndApi.md#DndInfo) | **Get** /dnd.info | 
[**DndSetSnooze**](DndApi.md#DndSetSnooze) | **Post** /dnd.setSnooze | 
[**DndTeamInfo**](DndApi.md#DndTeamInfo) | **Get** /dnd.teamInfo | 


# **DndEndDnd**
> map[string]interface{} DndEndDnd(ctx, )


Ends the current user's Do Not Disturb session immediately.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DndEndSnooze**
> map[string]interface{} DndEndSnooze(ctx, )


Ends the current user's snooze mode immediately.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DndInfo**
> map[string]interface{} DndInfo(ctx, optional)


Retrieves a user's current Do Not Disturb status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DndInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DndInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;dnd:read&#x60; | 
 **user** | **optional.String**| User to fetch status for (defaults to current user) | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DndSetSnooze**
> map[string]interface{} DndSetSnooze(ctx, )


Turns on Do Not Disturb mode for the current user, or changes its duration.

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

# **DndTeamInfo**
> map[string]interface{} DndTeamInfo(ctx, optional)


Retrieves the Do Not Disturb status for up to 50 users on a team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DndTeamInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DndTeamInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;dnd:read&#x60; | 
 **users** | **optional.String**| Comma-separated list of users to fetch Do Not Disturb status for | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

