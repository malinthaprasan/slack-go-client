/*
 * Slack Web API
 *
 * One way to interact with the Slack platform is its HTTP RPC-based Web API, a collection of methods requiring OAuth 2.0-based user, bot, or workspace tokens blessed with related OAuth scopes.
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ObjsIm struct {
	Created int32 `json:"created"`
	Id *DefsDmId `json:"id"`
	IsAppHome bool `json:"is_app_home,omitempty"`
	IsExtShared bool `json:"is_ext_shared,omitempty"`
	IsIm bool `json:"is_im"`
	IsOrgShared bool `json:"is_org_shared"`
	IsShared bool `json:"is_shared,omitempty"`
	IsUserDeleted bool `json:"is_user_deleted"`
	Priority float32 `json:"priority,omitempty"`
	User *DefsUserId `json:"user"`
}
