# \ImApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImClose**](ImApi.md#ImClose) | **Post** /im.close | 
[**ImHistory**](ImApi.md#ImHistory) | **Get** /im.history | 
[**ImList**](ImApi.md#ImList) | **Get** /im.list | 
[**ImMark**](ImApi.md#ImMark) | **Post** /im.mark | 
[**ImOpen**](ImApi.md#ImOpen) | **Post** /im.open | 
[**ImReplies**](ImApi.md#ImReplies) | **Get** /im.replies | 


# **ImClose**
> map[string]interface{} ImClose(ctx, )


Close a direct message channel.

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

# **ImHistory**
> map[string]interface{} ImHistory(ctx, optional)


Fetches history of messages and events from direct message channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of messages to return, between 1 and 1000. | 
 **unreads** | **optional.Bool**| Include &#x60;unread_count_display&#x60; in the output? | 
 **inclusive** | **optional.Bool**| Include messages with latest or oldest timestamp in results. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;im:history&#x60; | 
 **oldest** | **optional.Float32**| Start of time range of messages to include in results. | 
 **channel** | **optional.String**| Direct message channel to fetch history for. | 
 **latest** | **optional.Float32**| End of time range of messages to include in results. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImList**
> map[string]interface{} ImList(ctx, optional)


Lists direct message channels for the calling user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;im:read&#x60; | 
 **limit** | **optional.Int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn&#39;t been reached. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImMark**
> map[string]interface{} ImMark(ctx, optional)


Sets the read cursor in a direct message channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImMarkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImMarkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;im:write&#x60; | 
 **channel** | **optional.String**| Direct message channel to set reading cursor in. | 
 **ts** | **optional.Float32**| Timestamp of the most recently seen message. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImOpen**
> map[string]interface{} ImOpen(ctx, optional)


Opens a direct message channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImOpenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImOpenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;im:write&#x60; | 
 **returnIm** | **optional.Bool**| Boolean, indicates you want the full IM channel definition in the response. | 
 **user** | **optional.String**| User to open a direct message channel with. | 
 **includeLocale** | **optional.Bool**| Set this to &#x60;true&#x60; to receive the locale for this im. Defaults to &#x60;false&#x60; | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImReplies**
> map[string]interface{} ImReplies(ctx, optional)


Retrieve a thread of messages posted to a direct message conversation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImRepliesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImRepliesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadTs** | **optional.Float32**| Unique identifier of a thread&#39;s parent message | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;im:history&#x60; | 
 **channel** | **optional.String**| Direct message channel to fetch thread from | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

