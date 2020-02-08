# \SearchApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchMessages**](SearchApi.md#SearchMessages) | **Get** /search.messages | 


# **SearchMessages**
> DefaultSuccessTemplate SearchMessages(ctx, optional)


Searches for messages matching a query.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SearchMessagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SearchMessagesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortDir** | **optional.String**| Change sort direction to ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;). | 
 **query** | **optional.String**| Search query. | 
 **sort** | **optional.String**| Return matches sorted by either &#x60;score&#x60; or &#x60;timestamp&#x60;. | 
 **count** | **optional.String**| Pass the number of results you want per \&quot;page\&quot;. Maximum of &#x60;100&#x60;. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;search:read&#x60; | 
 **highlight** | **optional.Bool**| Pass a value of &#x60;true&#x60; to enable query highlight markers (see below). | 
 **page** | **optional.String**|  | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

