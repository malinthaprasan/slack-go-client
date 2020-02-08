# \OauthApi

All URIs are relative to *https://slack.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OauthAccess**](OauthApi.md#OauthAccess) | **Get** /oauth.access | 
[**OauthToken**](OauthApi.md#OauthToken) | **Get** /oauth.token | 


# **OauthAccess**
> DefaultSuccessTemplate OauthAccess(ctx, optional)


Exchanges a temporary OAuth verifier code for an access token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OauthAccessOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthAccessOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **code** | **optional.String**| The &#x60;code&#x60; param returned via the OAuth callback. | 
 **redirectUri** | **optional.String**| This must match the originally submitted URI (if one was sent). | 
 **clientId** | **optional.String**| Issued when you created your application. | 
 **clientSecret** | **optional.String**| Issued when you created your application. | 
 **singleChannel** | **optional.Bool**| Request the user to add your app only to a single channel. | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OauthToken**
> DefaultSuccessTemplate OauthToken(ctx, optional)


Exchanges a temporary OAuth verifier code for a workspace token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***OauthTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OauthTokenOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientSecret** | **optional.String**| Issued when you created your application. | 
 **code** | **optional.String**| The &#x60;code&#x60; param returned via the OAuth callback. | 
 **singleChannel** | **optional.Bool**| Request the user to add your app only to a single channel. | 
 **clientId** | **optional.String**| Issued when you created your application. | 
 **redirectUri** | **optional.String**| This must match the originally submitted URI (if one was sent). | 

### Return type

[**DefaultSuccessTemplate**](Default success template.md)

### Authorization

[slackAuth](../README.md#slackAuth)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

