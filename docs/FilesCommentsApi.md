# \FilesCommentsApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesCommentsDelete**](FilesCommentsApi.md#FilesCommentsDelete) | **Post** /files.comments.delete | 


# **FilesCommentsDelete**
> map[string]interface{} FilesCommentsDelete(ctx, optional)


Deletes an existing comment on a file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FilesCommentsDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesCommentsDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;files:write:user&#x60; | 
 **id** | **optional.String**| The comment to delete. | 
 **file** | **optional.String**| File to delete a comment from. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

