# \GroupsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GroupsArchive**](GroupsApi.md#GroupsArchive) | **Post** /groups.archive | 
[**GroupsCreate**](GroupsApi.md#GroupsCreate) | **Post** /groups.create | 
[**GroupsCreateChild**](GroupsApi.md#GroupsCreateChild) | **Post** /groups.createChild | 
[**GroupsHistory**](GroupsApi.md#GroupsHistory) | **Get** /groups.history | 
[**GroupsInfo**](GroupsApi.md#GroupsInfo) | **Get** /groups.info | 
[**GroupsInvite**](GroupsApi.md#GroupsInvite) | **Post** /groups.invite | 
[**GroupsKick**](GroupsApi.md#GroupsKick) | **Post** /groups.kick | 
[**GroupsLeave**](GroupsApi.md#GroupsLeave) | **Post** /groups.leave | 
[**GroupsList**](GroupsApi.md#GroupsList) | **Get** /groups.list | 
[**GroupsMark**](GroupsApi.md#GroupsMark) | **Post** /groups.mark | 
[**GroupsOpen**](GroupsApi.md#GroupsOpen) | **Post** /groups.open | 
[**GroupsRename**](GroupsApi.md#GroupsRename) | **Post** /groups.rename | 
[**GroupsReplies**](GroupsApi.md#GroupsReplies) | **Get** /groups.replies | 
[**GroupsSetPurpose**](GroupsApi.md#GroupsSetPurpose) | **Post** /groups.setPurpose | 
[**GroupsSetTopic**](GroupsApi.md#GroupsSetTopic) | **Post** /groups.setTopic | 
[**GroupsUnarchive**](GroupsApi.md#GroupsUnarchive) | **Post** /groups.unarchive | 


# **GroupsArchive**
> map[string]interface{} GroupsArchive(ctx, optional)


Archives a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsArchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsArchiveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **channel** | **optional.String**| Private channel to archive | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsCreate**
> map[string]interface{} GroupsCreate(ctx, optional)


Creates a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validate** | **optional.Bool**| Whether to return errors on invalid channel name instead of modifying it to meet the specified criteria. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **name** | **optional.String**| Name of private channel to create | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsCreateChild**
> map[string]interface{} GroupsCreateChild(ctx, optional)


Clones and archives a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsCreateChildOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsCreateChildOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **channel** | **optional.String**| Private channel to clone and archive. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsHistory**
> map[string]interface{} GroupsHistory(ctx, optional)


Fetches history of messages and events from a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| Number of messages to return, between 1 and 1000. | 
 **unreads** | **optional.Bool**| Include &#x60;unread_count_display&#x60; in the output? | 
 **inclusive** | **optional.Bool**| Include messages with latest or oldest timestamp in results. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:history&#x60; | 
 **oldest** | **optional.Float32**| Start of time range of messages to include in results. | 
 **channel** | **optional.String**| Private channel to fetch history for. | 
 **latest** | **optional.Float32**| End of time range of messages to include in results. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsInfo**
> map[string]interface{} GroupsInfo(ctx, optional)


Gets information about a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:read&#x60; | 
 **includeLocale** | **optional.Bool**| Set this to &#x60;true&#x60; to receive the locale for this group. Defaults to &#x60;false&#x60; | 
 **channel** | **optional.String**| Private channel to get info on | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsInvite**
> map[string]interface{} GroupsInvite(ctx, optional)


Invites a user to a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsInviteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **user** | **optional.String**| User to invite. | 
 **channel** | **optional.String**| Private channel to invite user to. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsKick**
> map[string]interface{} GroupsKick(ctx, optional)


Removes a user from a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsKickOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsKickOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **user** | **optional.String**| User to remove from private channel. | 
 **channel** | **optional.String**| Private channel to remove user from. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsLeave**
> map[string]interface{} GroupsLeave(ctx, optional)


Leaves a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsLeaveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsLeaveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **channel** | **optional.String**| Private channel to leave | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsList**
> map[string]interface{} GroupsList(ctx, optional)


Lists private channels that the calling user has access to.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Parameter for pagination. Set &#x60;cursor&#x60; equal to the &#x60;next_cursor&#x60; attribute returned by the previous request&#39;s &#x60;response_metadata&#x60;. This parameter is optional, but pagination is mandatory: the default value simply fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more details. | 
 **excludeMembers** | **optional.Bool**| Exclude the &#x60;members&#x60; from each &#x60;group&#x60; | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:read&#x60; | 
 **excludeArchived** | **optional.Bool**| Don&#39;t return archived private channels. | 
 **limit** | **optional.Int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the list hasn&#39;t been reached. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsMark**
> map[string]interface{} GroupsMark(ctx, optional)


Sets the read cursor in a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsMarkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsMarkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **ts** | **optional.Float32**| Timestamp of the most recently seen message. | 
 **channel** | **optional.String**| Private channel to set reading cursor in. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsOpen**
> map[string]interface{} GroupsOpen(ctx, optional)


Opens a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsOpenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsOpenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **channel** | **optional.String**| Private channel to open. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsRename**
> map[string]interface{} GroupsRename(ctx, optional)


Renames a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsRenameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsRenameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validate** | **optional.Bool**| Whether to return errors on invalid channel name instead of modifying it to meet the specified criteria. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **name** | **optional.String**| New name for private channel. | 
 **channel** | **optional.String**| Private channel to rename | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsReplies**
> map[string]interface{} GroupsReplies(ctx, optional)


Retrieve a thread of messages posted to a private channel

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsRepliesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsRepliesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadTs** | **optional.Float32**| Unique identifier of a thread&#39;s parent message | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:history&#x60; | 
 **channel** | **optional.String**| Private channel to fetch thread from | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsSetPurpose**
> map[string]interface{} GroupsSetPurpose(ctx, optional)


Sets the purpose for a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsSetPurposeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsSetPurposeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **purpose** | **optional.String**| The new purpose | 
 **channel** | **optional.String**| Private channel to set the purpose of | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsSetTopic**
> map[string]interface{} GroupsSetTopic(ctx, optional)


Sets the topic for a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsSetTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsSetTopicOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topic** | **optional.String**| The new topic | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **channel** | **optional.String**| Private channel to set the topic of | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GroupsUnarchive**
> map[string]interface{} GroupsUnarchive(ctx, optional)


Unarchives a private channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GroupsUnarchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GroupsUnarchiveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;groups:write&#x60; | 
 **channel** | **optional.String**| Private channel to unarchive | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

