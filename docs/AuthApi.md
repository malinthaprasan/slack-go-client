# \AuthApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthRevoke**](AuthApi.md#AuthRevoke) | **Get** /auth.revoke | 
[**AuthTest**](AuthApi.md#AuthTest) | **Get** /auth.test | 


# **AuthRevoke**
> map[string]interface{} AuthRevoke(ctx, optional)


Revokes a token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthRevokeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AuthRevokeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **test** | **optional.Bool**| Setting this parameter to &#x60;1&#x60; triggers a _testing mode_ where the specified token will not actually be revoked. | 
 **token** | **optional.String**| Authentication token. Requires scope: &#x60;none&#x60; | 

### Return type

[**map[string]interface{}**](interface{}.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AuthTest**
> map[string]interface{} AuthTest(ctx, )


Checks authentication & identity.

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

