# \FilesApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FilesCommentsDelete**](FilesApi.md#FilesCommentsDelete) | **Post** /files.comments.delete | 
[**FilesDelete**](FilesApi.md#FilesDelete) | **Post** /files.delete | 
[**FilesInfo**](FilesApi.md#FilesInfo) | **Get** /files.info | 
[**FilesList**](FilesApi.md#FilesList) | **Get** /files.list | 
[**FilesRevokePublicURL**](FilesApi.md#FilesRevokePublicURL) | **Post** /files.revokePublicURL | 
[**FilesSharedPublicURL**](FilesApi.md#FilesSharedPublicURL) | **Post** /files.sharedPublicURL | 
[**FilesUpload**](FilesApi.md#FilesUpload) | **Post** /files.upload | 


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

# **FilesDelete**
> map[string]interface{} FilesDelete(ctx, optional)


Deletes a file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FilesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;files:write:user&#x60; | 
 **file** | **optional.String**| ID of file to delete. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesInfo**
> map[string]interface{} FilesInfo(ctx, optional)


Gets information about a team file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FilesInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.String**|  | 
 **cursor** | **optional.String**| Parameter for pagination. File comments are paginated for a single file. Set &#x60;cursor&#x60; equal to the &#x60;next_cursor&#x60; attribute returned by the previous request&#39;s &#x60;response_metadata&#x60;. This parameter is optional, but pagination is mandatory: the default value simply fetches the first \&quot;page\&quot; of the collection of comments. See [pagination](/docs/pagination) for more details. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;files:read&#x60; | 
 **limit** | **optional.Int32**| The maximum number of items to return. Fewer than the requested number of items may be returned, even if the end of the list hasn&#39;t been reached. | 
 **file** | **optional.String**| Specify a file by providing its ID. | 
 **page** | **optional.String**|  | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesList**
> map[string]interface{} FilesList(ctx, optional)


Lists & filters team files.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FilesListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.String**|  | 
 **channel** | **optional.String**| Filter files appearing in a specific channel, indicated by its ID. | 
 **tsTo** | **optional.Float32**| Filter files created before this timestamp (inclusive). | 
 **tsFrom** | **optional.Float32**| Filter files created after this timestamp (inclusive). | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;files:read&#x60; | 
 **user** | **optional.String**| Filter files created by a single user. | 
 **page** | **optional.String**|  | 
 **types** | **optional.String**| Filter files by type ([see below](#file_types)). You can pass multiple values in the types argument, like &#x60;types&#x3D;spaces,snippets&#x60;.The default value is &#x60;all&#x60;, which does not filter the list. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesRevokePublicURL**
> map[string]interface{} FilesRevokePublicURL(ctx, optional)


Revokes public/external sharing access for a file

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FilesRevokePublicURLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesRevokePublicURLOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;files:write:user&#x60; | 
 **file** | **optional.String**| File to revoke | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesSharedPublicURL**
> map[string]interface{} FilesSharedPublicURL(ctx, optional)


Enables a file for public/external sharing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FilesSharedPublicURLOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesSharedPublicURLOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;files:write:user&#x60; | 
 **file** | **optional.String**| File to share | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded, application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **FilesUpload**
> map[string]interface{} FilesUpload(ctx, optional)


Uploads or creates a file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***FilesUploadOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FilesUploadOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channels** | **optional.String**| Comma-separated list of channel names or IDs where the file will be shared. | 
 **title** | **optional.String**| Title of file. | 
 **initialComment** | **optional.String**| The message text introducing the file in specified &#x60;channels&#x60;. | 
 **filetype** | **optional.String**| A [file type](/types/file#file_types) identifier. | 
 **filename** | **optional.String**| Filename of file. | 
 **content** | **optional.String**| File contents via a POST variable. If omitting this parameter, you must provide a &#x60;file&#x60;. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;files:write:user&#x60; | 
 **file** | **optional.String**| File contents via &#x60;multipart/form-data&#x60;. If omitting this parameter, you must submit &#x60;content&#x60;. | 
 **threadTs** | **optional.Float32**| Provide another message&#39;s &#x60;ts&#x60; value to upload this file as a reply. Never use a reply&#39;s &#x60;ts&#x60; value; use its parent instead. | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

