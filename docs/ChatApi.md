# \ChatApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChatDelete**](ChatApi.md#ChatDelete) | **Post** /chat.delete | 
[**ChatDeleteScheduledMessage**](ChatApi.md#ChatDeleteScheduledMessage) | **Post** /chat.deleteScheduledMessage | 
[**ChatGetPermalink**](ChatApi.md#ChatGetPermalink) | **Get** /chat.getPermalink | 
[**ChatMeMessage**](ChatApi.md#ChatMeMessage) | **Post** /chat.meMessage | 
[**ChatPostEphemeral**](ChatApi.md#ChatPostEphemeral) | **Post** /chat.postEphemeral | 
[**ChatPostMessage**](ChatApi.md#ChatPostMessage) | **Post** /chat.postMessage | 
[**ChatScheduleMessage**](ChatApi.md#ChatScheduleMessage) | **Post** /chat.scheduleMessage | 
[**ChatScheduledMessagesList**](ChatApi.md#ChatScheduledMessagesList) | **Get** /chat.scheduledMessages.list | 
[**ChatUnfurl**](ChatApi.md#ChatUnfurl) | **Post** /chat.unfurl | 
[**ChatUpdate**](ChatApi.md#ChatUpdate) | **Post** /chat.update | 


# **ChatDelete**
> map[string]interface{} ChatDelete(ctx, optional)


Deletes a message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChatDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChatDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **asUser** | **optional.Bool**| Pass true to delete the message as the authed user with &#x60;chat:write:user&#x60; scope. [Bot users](/bot-users) in this context are considered authed users. If unused or false, the message will be deleted with &#x60;chat:write:bot&#x60; scope. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;chat:write&#x60; | 
 **ts** | **optional.Float32**| Timestamp of the message to be deleted. | 
 **channel** | **optional.String**| Channel containing the message to be deleted. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChatDeleteScheduledMessage**
> map[string]interface{} ChatDeleteScheduledMessage(ctx, )


Deletes a pending scheduled message from the queue.

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

# **ChatGetPermalink**
> map[string]interface{} ChatGetPermalink(ctx, optional)


Retrieve a permalink URL for a specific extant message

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChatGetPermalinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChatGetPermalinkOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;none&#x60; | 
 **messageTs** | **optional.Float32**| A message&#39;s &#x60;ts&#x60; value, uniquely identifying it within a channel | 
 **channel** | **optional.String**| The ID of the conversation or channel containing the message | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChatMeMessage**
> map[string]interface{} ChatMeMessage(ctx, optional)


Share a me message into a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChatMeMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChatMeMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **optional.String**| Text of the message to send. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;chat:write:user&#x60; | 
 **channel** | **optional.String**| Channel to send message to. Can be a public channel, private group or IM channel. Can be an encoded ID, or a name. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChatPostEphemeral**
> map[string]interface{} ChatPostEphemeral(ctx, optional)


Sends an ephemeral message to a user in a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChatPostEphemeralOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChatPostEphemeralOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadTs** | **optional.Float32**| Provide another message&#39;s &#x60;ts&#x60; value to post this message in a thread. Avoid using a reply&#39;s &#x60;ts&#x60; value; use its parent&#39;s value instead. Ephemeral messages in threads are only shown if there is already an active thread. | 
 **blocks** | **optional.String**| A JSON-based array of structured blocks, presented as a URL-encoded string. | 
 **attachments** | **optional.String**| A JSON-based array of structured attachments, presented as a URL-encoded string. | 
 **asUser** | **optional.Bool**| Pass true to post the message as the authed user. Defaults to true if the chat:write:bot scope is not included. Otherwise, defaults to false. | 
 **parse** | **optional.String**| Change how messages are treated. Defaults to &#x60;none&#x60;. See [below](#formatting). | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;chat:write&#x60; | 
 **text** | **optional.String**| Text of the message to send. See below for an explanation of [formatting](#formatting). This field is usually required, unless you&#39;re providing only &#x60;attachments&#x60; instead. | 
 **user** | **optional.String**| &#x60;id&#x60; of the user who will receive the ephemeral message. The user should be in the channel specified by the &#x60;channel&#x60; argument. | 
 **linkNames** | **optional.Bool**| Find and link channel names and usernames. | 
 **channel** | **optional.String**| Channel, private group, or IM channel to send message to. Can be an encoded ID, or a name. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChatPostMessage**
> map[string]interface{} ChatPostMessage(ctx, optional)


Sends a message to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChatPostMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChatPostMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attachments** | **optional.String**| A JSON-based array of structured attachments, presented as a URL-encoded string. | 
 **unfurlLinks** | **optional.Bool**| Pass true to enable unfurling of primarily text-based content. | 
 **text** | **optional.String**| Text of the message to send. See below for an explanation of [formatting](#formatting). This field is usually required, unless you&#39;re providing only &#x60;attachments&#x60; instead. Provide no more than 40,000 characters or [risk truncation](/changelog/2018-04-truncating-really-long-messages). | 
 **unfurlMedia** | **optional.Bool**| Pass false to disable unfurling of media content. | 
 **parse** | **optional.String**| Change how messages are treated. Defaults to &#x60;none&#x60;. See [below](#formatting). | 
 **asUser** | **optional.Bool**| Pass true to post the message as the authed user, instead of as a bot. Defaults to false. See [authorship](#authorship) below. | 
 **mrkdwn** | **optional.Bool**| Disable Slack markup parsing by setting to &#x60;false&#x60;. Enabled by default. | 
 **channel** | **optional.String**| Channel, private group, or IM channel to send message to. Can be an encoded ID, or a name. See [below](#channels) for more details. | 
 **username** | **optional.String**| Set your bot&#39;s user name. Must be used in conjunction with &#x60;as_user&#x60; set to false, otherwise ignored. See [authorship](#authorship) below. | 
 **blocks** | **optional.String**| A JSON-based array of structured blocks, presented as a URL-encoded string. | 
 **iconEmoji** | **optional.String**| Emoji to use as the icon for this message. Overrides &#x60;icon_url&#x60;. Must be used in conjunction with &#x60;as_user&#x60; set to &#x60;false&#x60;, otherwise ignored. See [authorship](#authorship) below. | 
 **linkNames** | **optional.Bool**| Find and link channel names and usernames. | 
 **replyBroadcast** | **optional.Bool**| Used in conjunction with &#x60;thread_ts&#x60; and indicates whether reply should be made visible to everyone in the channel or conversation. Defaults to &#x60;false&#x60;. | 
 **threadTs** | **optional.Float32**| Provide another message&#39;s &#x60;ts&#x60; value to make this message a reply. Avoid using a reply&#39;s &#x60;ts&#x60; value; use its parent instead. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;chat:write&#x60; | 
 **iconUrl** | **optional.String**| URL to an image to use as the icon for this message. Must be used in conjunction with &#x60;as_user&#x60; set to false, otherwise ignored. See [authorship](#authorship) below. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChatScheduleMessage**
> map[string]interface{} ChatScheduleMessage(ctx, optional)


Schedules a message to be sent to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChatScheduleMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChatScheduleMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **threadTs** | **optional.Float32**| Provide another message&#39;s &#x60;ts&#x60; value to make this message a reply. Avoid using a reply&#39;s &#x60;ts&#x60; value; use its parent instead. | 
 **blocks** | **optional.String**| A JSON-based array of structured blocks, presented as a URL-encoded string. | 
 **attachments** | **optional.String**| A JSON-based array of structured attachments, presented as a URL-encoded string. | 
 **unfurlLinks** | **optional.Bool**| Pass true to enable unfurling of primarily text-based content. | 
 **text** | **optional.String**| Text of the message to send. See below for an explanation of [formatting](#formatting). This field is usually required, unless you&#39;re providing only &#x60;attachments&#x60; instead. Provide no more than 40,000 characters or [risk truncation](/changelog/2018-04-truncating-really-long-messages). | 
 **linkNames** | **optional.Bool**| Find and link channel names and usernames. | 
 **unfurlMedia** | **optional.Bool**| Pass false to disable unfurling of media content. | 
 **parse** | **optional.String**| Change how messages are treated. Defaults to &#x60;none&#x60;. See [below](#formatting). | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;chat:write&#x60; | 
 **asUser** | **optional.Bool**| Pass true to post the message as the authed user, instead of as a bot. Defaults to false. See [authorship](#authorship) below. | 
 **postAt** | **optional.String**| Unix EPOCH timestamp of time in future to send the message. | 
 **channel** | **optional.String**| Channel, private group, or DM channel to send message to. Can be an encoded ID, or a name. See [below](#channels) for more details. | 
 **replyBroadcast** | **optional.Bool**| Used in conjunction with &#x60;thread_ts&#x60; and indicates whether reply should be made visible to everyone in the channel or conversation. Defaults to &#x60;false&#x60;. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **ChatUnfurl**
> map[string]interface{} ChatUnfurl(ctx, optional)


Provide custom unfurl behavior for user-posted URLs

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChatUnfurlOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChatUnfurlOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAuthMessage** | **optional.String**| Provide a simply-formatted string to send as an ephemeral message to the user as invitation to authenticate further and enable full unfurling behavior | 
 **userAuthRequired** | **optional.Bool**| Set to &#x60;true&#x60; or &#x60;1&#x60; to indicate the user must install your Slack app to trigger unfurls for this domain | 
 **unfurls** | **optional.String**| URL-encoded JSON map with keys set to URLs featured in the the message, pointing to their unfurl blocks or message attachments. | 
 **ts** | **optional.String**| Timestamp of the message to add unfurl behavior to. | 
 **userAuthUrl** | **optional.String**| Send users to this custom URL where they will complete authentication in your app to fully trigger unfurling. Value should be properly URL-encoded. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;links:write&#x60; | 
 **channel** | **optional.String**| Channel ID of the message | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChatUpdate**
> map[string]interface{} ChatUpdate(ctx, optional)


Updates a message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ChatUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChatUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blocks** | **optional.String**| A JSON-based array of structured blocks, presented as a URL-encoded string. | 
 **attachments** | **optional.String**| A JSON-based array of structured attachments, presented as a URL-encoded string. This field is required when not presenting &#x60;text&#x60;. | 
 **text** | **optional.String**| New text for the message, using the [default formatting rules](/docs/formatting). It&#39;s not required when presenting &#x60;attachments&#x60;. | 
 **ts** | **optional.Float32**| Timestamp of the message to be updated. | 
 **parse** | **optional.String**| Change how messages are treated. Defaults to &#x60;client&#x60;, unlike &#x60;chat.postMessage&#x60;. See [below](#formatting). | 
 **asUser** | **optional.Bool**| Pass true to update the message as the authed user. [Bot users](/bot-users) in this context are considered authed users. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;chat:write&#x60; | 
 **linkNames** | **optional.Bool**| Find and link channel names and usernames. Defaults to &#x60;none&#x60;. See [below](#formatting). | 
 **channel** | **optional.String**| Channel containing the message to be updated. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

