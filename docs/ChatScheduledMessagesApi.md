# \ChatScheduledMessagesApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatScheduledMessagesList**](ChatScheduledMessagesApi.md#ChatScheduledMessagesList) | **Get** /chat.scheduledMessages.list | 


# **ChatScheduledMessagesList**
> map[string]interface{} ChatScheduledMessagesList(ctx, optional)


Returns a list of scheduled messages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChatScheduledMessagesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChatScheduledMessagesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| For pagination purposes, this is the &#x60;cursor&#x60; value returned from a previous call to &#x60;chat.scheduledmessages.list&#x60; indicating where you want to start this call from. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;none&#x60; | 
 **limit** | **optional.Int32**| Maximum number of original entries to return. | 
 **oldest** | **optional.Float32**| A UNIX timestamp of the oldest value in the time range | 
 **channel** | **optional.String**| The channel of the scheduled messages | 
 **latest** | **optional.Float32**| A UNIX timestamp of the latest value in the time range | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

