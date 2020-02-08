# \ChannelsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsArchive**](ChannelsApi.md#ChannelsArchive) | **Post** /channels.archive | 
[**ChannelsCreate**](ChannelsApi.md#ChannelsCreate) | **Post** /channels.create | 
[**ChannelsHistory**](ChannelsApi.md#ChannelsHistory) | **Get** /channels.history | 
[**ChannelsInfo**](ChannelsApi.md#ChannelsInfo) | **Get** /channels.info | 
[**ChannelsInvite**](ChannelsApi.md#ChannelsInvite) | **Post** /channels.invite | 
[**ChannelsJoin**](ChannelsApi.md#ChannelsJoin) | **Post** /channels.join | 
[**ChannelsKick**](ChannelsApi.md#ChannelsKick) | **Post** /channels.kick | 
[**ChannelsLeave**](ChannelsApi.md#ChannelsLeave) | **Post** /channels.leave | 
[**ChannelsList**](ChannelsApi.md#ChannelsList) | **Get** /channels.list | 
[**ChannelsMark**](ChannelsApi.md#ChannelsMark) | **Post** /channels.mark | 
[**ChannelsRename**](ChannelsApi.md#ChannelsRename) | **Post** /channels.rename | 
[**ChannelsReplies**](ChannelsApi.md#ChannelsReplies) | **Get** /channels.replies | 
[**ChannelsSetPurpose**](ChannelsApi.md#ChannelsSetPurpose) | **Post** /channels.setPurpose | 
[**ChannelsSetTopic**](ChannelsApi.md#ChannelsSetTopic) | **Post** /channels.setTopic | 
[**ChannelsUnarchive**](ChannelsApi.md#ChannelsUnarchive) | **Post** /channels.unarchive | 


# **ChannelsArchive**
> map[string]interface{} ChannelsArchive(ctx, optional)


Archives a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsArchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsArchiveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **channel** | **optional.String**| Channel to archive | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsCreate**
> map[string]interface{} ChannelsCreate(ctx, optional)


Creates a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validate** | **optional.Bool**| Whether to return errors on invalid channel name instead of modifying it to meet the specified criteria. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **name** | **optional.String**| Name of channel to create | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsHistory**
> map[string]interface{} ChannelsHistory(ctx, optional)


Fetches history of messages and events from a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of messages to return, between 1 and 1000. | 
 **unreads** | **optional.Bool**| Include &#x60;unread_count_display&#x60; in the output? | 
 **inclusive** | **optional.Bool**| Include messages with latest or oldest timestamp in results. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:history&#x60; | 
 **oldest** | **optional.Float32**| Start of time range of messages to include in results. | 
 **channel** | **optional.String**| Channel to fetch history for. | 
 **latest** | **optional.Float32**| End of time range of messages to include in results. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsInfo**
> map[string]interface{} ChannelsInfo(ctx, optional)


Gets information about a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:read&#x60; | 
 **includeLocale** | **optional.Bool**| Set this to &#x60;true&#x60; to receive the locale for this channel. Defaults to &#x60;false&#x60; | 
 **channel** | **optional.String**| Channel to get info on | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsInvite**
> map[string]interface{} ChannelsInvite(ctx, optional)


Invites a user to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsInviteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **user** | **optional.String**| User to invite to channel. | 
 **channel** | **optional.String**| Channel to invite user to. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsJoin**
> map[string]interface{} ChannelsJoin(ctx, optional)


Joins a channel, creating it if needed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsJoinOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsJoinOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validate** | **optional.Bool**| Whether to return errors on invalid channel name instead of modifying it to meet the specified criteria. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **name** | **optional.String**| Name of channel to join | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsKick**
> map[string]interface{} ChannelsKick(ctx, optional)


Removes a user from a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsKickOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsKickOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **user** | **optional.String**| User to remove from channel. | 
 **channel** | **optional.String**| Channel to remove user from. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsLeave**
> map[string]interface{} ChannelsLeave(ctx, optional)


Leaves a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsLeaveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsLeaveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **channel** | **optional.String**| Channel to leave | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsList**
> map[string]interface{} ChannelsList(ctx, optional)


Lists all channels in a Slack team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **excludeMembers** | **optional.Bool**| Exclude the &#x60;members&#x60; collection from each &#x60;channel&#x60; | 
 **cursor** | **optional.String**| Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:read&#x60; | 
 **limit** | **optional.Int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn&#39;t been reached. | 
 **excludeArchived** | **optional.Bool**| Exclude archived channels from the list | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsMark**
> map[string]interface{} ChannelsMark(ctx, optional)


Sets the read cursor in a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsMarkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsMarkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **ts** | **optional.Float32**| Timestamp of the most recently seen message. | 
 **channel** | **optional.String**| Channel to set reading cursor in. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsRename**
> map[string]interface{} ChannelsRename(ctx, optional)


Renames a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsRenameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsRenameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validate** | **optional.Bool**| Whether to return errors on invalid channel name instead of modifying it to meet the specified criteria. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **name** | **optional.String**| New name for channel. | 
 **channel** | **optional.String**| Channel to rename | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsReplies**
> map[string]interface{} ChannelsReplies(ctx, optional)


Retrieve a thread of messages posted to a channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsRepliesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsRepliesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadTs** | **optional.Float32**| Unique identifier of a thread&#39;s parent message | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:history&#x60; | 
 **channel** | **optional.String**| Channel to fetch thread from | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsSetPurpose**
> map[string]interface{} ChannelsSetPurpose(ctx, optional)


Sets the purpose for a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsSetPurposeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsSetPurposeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **purpose** | **optional.String**| The new purpose | 
 **channel** | **optional.String**| Channel to set the purpose of | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsSetTopic**
> map[string]interface{} ChannelsSetTopic(ctx, optional)


Sets the topic for a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsSetTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsSetTopicOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topic** | **optional.String**| The new topic | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **channel** | **optional.String**| Channel to set the topic of | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsUnarchive**
> map[string]interface{} ChannelsUnarchive(ctx, optional)


Unarchives a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChannelsUnarchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsUnarchiveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **channel** | **optional.String**| Channel to unarchive | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

