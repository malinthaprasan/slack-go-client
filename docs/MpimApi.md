# \MpimApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MpimClose**](MpimApi.md#MpimClose) | **Post** /mpim.close | 
[**MpimHistory**](MpimApi.md#MpimHistory) | **Get** /mpim.history | 
[**MpimList**](MpimApi.md#MpimList) | **Get** /mpim.list | 
[**MpimMark**](MpimApi.md#MpimMark) | **Post** /mpim.mark | 
[**MpimOpen**](MpimApi.md#MpimOpen) | **Post** /mpim.open | 
[**MpimReplies**](MpimApi.md#MpimReplies) | **Get** /mpim.replies | 


# **MpimClose**
> map[string]interface{} MpimClose(ctx, optional)


Closes a multiparty direct message channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MpimCloseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MpimCloseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;mpim:write&#x60; | 
 **channel** | **optional.String**| MPIM to close. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpimHistory**
> map[string]interface{} MpimHistory(ctx, optional)


Fetches history of messages and events from a multiparty direct message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MpimHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MpimHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of messages to return, between 1 and 1000. | 
 **unreads** | **optional.Bool**| Include &#x60;unread_count_display&#x60; in the output? | 
 **inclusive** | **optional.Bool**| Include messages with latest or oldest timestamp in results. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;mpim:history&#x60; | 
 **oldest** | **optional.Float32**| Start of time range of messages to include in results. | 
 **channel** | **optional.String**| Multiparty direct message to fetch history for. | 
 **latest** | **optional.Float32**| End of time range of messages to include in results. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpimList**
> map[string]interface{} MpimList(ctx, optional)


Lists multiparty direct message channels for the calling user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MpimListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MpimListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Parameter for pagination. Set &#x60;cursor&#x60; equal to the &#x60;next_cursor&#x60; attribute returned by the previous request&#39;s &#x60;response_metadata&#x60;. This parameter is optional, but pagination is mandatory: the default value simply fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more details. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;mpim:read&#x60; | 
 **limit** | **optional.Int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the list hasn&#39;t been reached. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpimMark**
> map[string]interface{} MpimMark(ctx, optional)


Sets the read cursor in a multiparty direct message channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MpimMarkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MpimMarkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;mpim:write&#x60; | 
 **ts** | **optional.Float32**| Timestamp of the most recently seen message. | 
 **channel** | **optional.String**| multiparty direct message channel to set reading cursor in. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpimOpen**
> map[string]interface{} MpimOpen(ctx, optional)


This method opens a multiparty direct message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MpimOpenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MpimOpenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;mpim:write&#x60; | 
 **users** | **optional.String**| Comma separated lists of users.  The ordering of the users is preserved whenever a MPIM group is returned. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MpimReplies**
> map[string]interface{} MpimReplies(ctx, optional)


Retrieve a thread of messages posted to a direct message conversation from a multiparty direct message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MpimRepliesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MpimRepliesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadTs** | **optional.Float32**| Unique identifier of a thread&#39;s parent message. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;mpim:history&#x60; | 
 **channel** | **optional.String**| Multiparty direct message channel to fetch thread from. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

