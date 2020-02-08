# \ConversationsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConversationsArchive**](ConversationsApi.md#ConversationsArchive) | **Post** /conversations.archive | 
[**ConversationsClose**](ConversationsApi.md#ConversationsClose) | **Post** /conversations.close | 
[**ConversationsCreate**](ConversationsApi.md#ConversationsCreate) | **Post** /conversations.create | 
[**ConversationsHistory**](ConversationsApi.md#ConversationsHistory) | **Get** /conversations.history | 
[**ConversationsInfo**](ConversationsApi.md#ConversationsInfo) | **Get** /conversations.info | 
[**ConversationsInvite**](ConversationsApi.md#ConversationsInvite) | **Post** /conversations.invite | 
[**ConversationsJoin**](ConversationsApi.md#ConversationsJoin) | **Post** /conversations.join | 
[**ConversationsKick**](ConversationsApi.md#ConversationsKick) | **Post** /conversations.kick | 
[**ConversationsLeave**](ConversationsApi.md#ConversationsLeave) | **Post** /conversations.leave | 
[**ConversationsList**](ConversationsApi.md#ConversationsList) | **Get** /conversations.list | 
[**ConversationsMembers**](ConversationsApi.md#ConversationsMembers) | **Get** /conversations.members | 
[**ConversationsOpen**](ConversationsApi.md#ConversationsOpen) | **Post** /conversations.open | 
[**ConversationsRename**](ConversationsApi.md#ConversationsRename) | **Post** /conversations.rename | 
[**ConversationsReplies**](ConversationsApi.md#ConversationsReplies) | **Get** /conversations.replies | 
[**ConversationsSetPurpose**](ConversationsApi.md#ConversationsSetPurpose) | **Post** /conversations.setPurpose | 
[**ConversationsSetTopic**](ConversationsApi.md#ConversationsSetTopic) | **Post** /conversations.setTopic | 
[**ConversationsUnarchive**](ConversationsApi.md#ConversationsUnarchive) | **Post** /conversations.unarchive | 


# **ConversationsArchive**
> map[string]interface{} ConversationsArchive(ctx, optional)


Archives a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsArchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsArchiveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **channel** | **optional.String**| ID of conversation to archive | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsClose**
> map[string]interface{} ConversationsClose(ctx, optional)


Closes a direct message or multi-person direct message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsCloseOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsCloseOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **channel** | **optional.String**| Conversation to close. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsCreate**
> map[string]interface{} ConversationsCreate(ctx, optional)


Initiates a public or private channel-based conversation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **userIds** | **optional.String**| **Required** for workspace apps. A list of between 1 and 30 human users that will be added to the newly-created conversation. This argument has no effect when used by classic Slack apps. | 
 **name** | **optional.String**| Name of the public or private channel to create | 
 **isPrivate** | **optional.Bool**| Create a private channel instead of a public one | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsHistory**
> map[string]interface{} ConversationsHistory(ctx, optional)


Fetches a conversation's history of messages and events.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsHistoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsHistoryOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inclusive** | **optional.Bool**| Include messages with latest or oldest timestamp in results only when either timestamp is specified. | 
 **cursor** | **optional.String**| Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:history&#x60; | 
 **limit** | **optional.Int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn&#39;t been reached. | 
 **oldest** | **optional.Float32**| Start of time range of messages to include in results. | 
 **channel** | **optional.String**| Conversation ID to fetch history for. | 
 **latest** | **optional.Float32**| End of time range of messages to include in results. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsInfo**
> map[string]interface{} ConversationsInfo(ctx, optional)


Retrieve information about a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeNumMembers** | **optional.Bool**| Set to &#x60;true&#x60; to include the member count for the specified conversation. Defaults to &#x60;false&#x60; | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:read&#x60; | 
 **channel** | **optional.String**| Conversation ID to learn more about | 
 **includeLocale** | **optional.Bool**| Set this to &#x60;true&#x60; to receive the locale for this conversation. Defaults to &#x60;false&#x60; | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsInvite**
> map[string]interface{} ConversationsInvite(ctx, optional)


Invites users to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsInviteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsInviteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **users** | **optional.String**| A comma separated list of user IDs. Up to 30 users may be listed. | 
 **channel** | **optional.String**| The ID of the public or private channel to invite user(s) to. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsJoin**
> map[string]interface{} ConversationsJoin(ctx, optional)


Joins an existing conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsJoinOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsJoinOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;channels:write&#x60; | 
 **channel** | **optional.String**| ID of conversation to join | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsKick**
> map[string]interface{} ConversationsKick(ctx, optional)


Removes a user from a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsKickOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsKickOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **user** | **optional.String**| User ID to be removed. | 
 **channel** | **optional.String**| ID of conversation to remove user from. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsLeave**
> map[string]interface{} ConversationsLeave(ctx, optional)


Leaves a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsLeaveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsLeaveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **channel** | **optional.String**| Conversation to leave | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsList**
> map[string]interface{} ConversationsList(ctx, optional)


Lists all channels in a Slack team.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:read&#x60; | 
 **limit** | **optional.Int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the list hasn&#39;t been reached. Must be an integer no larger than 1000. | 
 **excludeArchived** | **optional.Bool**| Set to &#x60;true&#x60; to exclude archived channels from the list | 
 **types** | **optional.String**| Mix and match channel types by providing a comma-separated list of any combination of &#x60;public_channel&#x60;, &#x60;private_channel&#x60;, &#x60;mpim&#x60;, &#x60;im&#x60; | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsMembers**
> map[string]interface{} ConversationsMembers(ctx, optional)


Retrieve members of a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsMembersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsMembersOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:read&#x60; | 
 **limit** | **optional.Int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn&#39;t been reached. | 
 **channel** | **optional.String**| ID of the conversation to retrieve members for | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsOpen**
> map[string]interface{} ConversationsOpen(ctx, optional)


Opens or resumes a direct message or multi-person direct message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsOpenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsOpenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **returnIm** | **optional.Bool**| Boolean, indicates you want the full IM channel definition in the response. | 
 **users** | **optional.String**| Comma separated lists of users. If only one user is included, this creates a 1:1 DM.  The ordering of the users is preserved whenever a multi-person direct message is returned. Supply a &#x60;channel&#x60; when not supplying &#x60;users&#x60;. | 
 **channel** | **optional.String**| Resume a conversation by supplying an &#x60;im&#x60; or &#x60;mpim&#x60;&#39;s ID. Or provide the &#x60;users&#x60; field instead. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsRename**
> map[string]interface{} ConversationsRename(ctx, optional)


Renames a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsRenameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsRenameOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **name** | **optional.String**| New name for conversation. | 
 **channel** | **optional.String**| ID of conversation to rename | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsReplies**
> map[string]interface{} ConversationsReplies(ctx, optional)


Retrieve a thread of messages posted to a conversation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsRepliesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsRepliesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inclusive** | **optional.Bool**| Include messages with latest or oldest timestamp in results only when either timestamp is specified. | 
 **ts** | **optional.Float32**| Unique identifier of a thread&#39;s parent message. | 
 **cursor** | **optional.String**| Paginate through collections of data by setting the &#x60;cursor&#x60; parameter to a &#x60;next_cursor&#x60; attribute returned by a previous request&#39;s &#x60;response_metadata&#x60;. Default value fetches the first \&quot;page\&quot; of the collection. See [pagination](/docs/pagination) for more detail. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:history&#x60; | 
 **limit** | **optional.Int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the users list hasn&#39;t been reached. | 
 **oldest** | **optional.Float32**| Start of time range of messages to include in results. | 
 **channel** | **optional.String**| Conversation ID to fetch thread from. | 
 **latest** | **optional.Float32**| End of time range of messages to include in results. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsSetPurpose**
> map[string]interface{} ConversationsSetPurpose(ctx, optional)


Sets the purpose for a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsSetPurposeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsSetPurposeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **purpose** | **optional.String**| A new, specialer purpose | 
 **channel** | **optional.String**| Conversation to set the purpose of | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsSetTopic**
> map[string]interface{} ConversationsSetTopic(ctx, optional)


Sets the topic for a conversation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsSetTopicOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsSetTopicOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **topic** | **optional.String**| The new topic string. Does not support formatting or linkification. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **channel** | **optional.String**| Conversation to set the topic of | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConversationsUnarchive**
> map[string]interface{} ConversationsUnarchive(ctx, optional)


Reverses conversation archival.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConversationsUnarchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConversationsUnarchiveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;conversations:write&#x60; | 
 **channel** | **optional.String**| ID of conversation to unarchive | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

