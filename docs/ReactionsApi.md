# \ReactionsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReactionsAdd**](ReactionsApi.md#ReactionsAdd) | **Post** /reactions.add | 
[**ReactionsGet**](ReactionsApi.md#ReactionsGet) | **Get** /reactions.get | 
[**ReactionsList**](ReactionsApi.md#ReactionsList) | **Get** /reactions.list | 
[**ReactionsRemove**](ReactionsApi.md#ReactionsRemove) | **Post** /reactions.remove | 


# **ReactionsAdd**
> map[string]interface{} ReactionsAdd(ctx, optional)


Adds a reaction to an item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReactionsAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReactionsAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Reaction (emoji) name. | 
 **fileComment** | **optional.String**| File comment to add reaction to. Now that [file threads](/changelog/2018-05-file-threads-soon-tread#whats_changed) work the way you&#39;d expect, this argument is deprecated. Specify the timestamp and channel of the message associated with a file instead. | 
 **timestamp** | **optional.Float32**| Timestamp of the message to add reaction to. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;reactions:write&#x60; | 
 **file** | **optional.String**| File to add reaction to. Now that [file threads](/changelog/2018-05-file-threads-soon-tread#whats_changed) work the way you&#39;d expect, this argument is deprecated. Specify the timestamp and channel of the message associated with a file instead. | 
 **channel** | **optional.String**| Channel where the message to add reaction to was posted. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactionsGet**
> interface{} ReactionsGet(ctx, optional)


Gets reactions for an item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReactionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReactionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **full** | **optional.Bool**| If true always return the complete reaction list. | 
 **fileComment** | **optional.String**| File comment to get reactions for. | 
 **timestamp** | **optional.Float32**| Timestamp of the message to get reactions for. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;reactions:read&#x60; | 
 **file** | **optional.String**| File to get reactions for. | 
 **channel** | **optional.String**| Channel where the message to get reactions for was posted. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactionsList**
> map[string]interface{} ReactionsList(ctx, optional)


Lists reactions made by a user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReactionsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReactionsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.String**|  | 
 **full** | **optional.Bool**| If true always return the complete reaction list. | 
 **cursor** | **optional.String**| Parameter for pagination. Set &#x60;cursor&#x60; equal to the &#x60;next_cursor&#x60; attribute returned by the previous request&#39;s &#x60;response_metadata&#x60;. This parameter is optional, but pagination is mandatory: the default value simply fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more details. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;reactions:read&#x60; | 
 **limit** | **optional.Int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the list hasn&#39;t been reached. | 
 **user** | **optional.String**| Show reactions made by this user. Defaults to the authed user. | 
 **page** | **optional.String**|  | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReactionsRemove**
> map[string]interface{} ReactionsRemove(ctx, optional)


Removes a reaction from an item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReactionsRemoveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReactionsRemoveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Reaction (emoji) name. | 
 **fileComment** | **optional.String**| File comment to remove reaction from. | 
 **timestamp** | **optional.Float32**| Timestamp of the message to remove reaction from. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;reactions:write&#x60; | 
 **file** | **optional.String**| File to remove reaction from. | 
 **channel** | **optional.String**| Channel where the message to remove reaction from was posted. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

