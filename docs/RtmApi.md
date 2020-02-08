# \RtmApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RtmConnect**](RtmApi.md#RtmConnect) | **Get** /rtm.connect | 


# **RtmConnect**
> map[string]interface{} RtmConnect(ctx, optional)


Starts a Real Time Messaging session.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RtmConnectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RtmConnectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **presenceSub** | **optional.Bool**| Only deliver presence events when requested by subscription. See [presence subscriptions](/docs/presence-and-status#subscriptions). | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;rtm:stream&#x60; | 
 **batchPresenceAware** | **optional.Bool**| Batch presence deliveries via subscription. Enabling changes the shape of &#x60;presence_change&#x60; events. See [batch presence](/docs/presence-and-status#batching). | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

