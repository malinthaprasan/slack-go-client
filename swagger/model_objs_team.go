/*
 * Slack Web API
 *
 * One way to interact with the Slack platform is its HTTP RPC-based Web API, a collection of methods requiring OAuth 2.0-based user, bot, or workspace tokens blessed with related OAuth scopes.
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ObjsTeam struct {
	Archived bool `json:"archived,omitempty"`
	AvatarBaseUrl string `json:"avatar_base_url,omitempty"`
	Created int32 `json:"created,omitempty"`
	DateCreate int32 `json:"date_create,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	Description string `json:"description,omitempty"`
	Discoverable string `json:"discoverable,omitempty"`
	Domain string `json:"domain"`
	EmailDomain string `json:"email_domain"`
	EnterpriseId *DefsEnterpriseId `json:"enterprise_id,omitempty"`
	EnterpriseName *DefsEnterpriseName `json:"enterprise_name,omitempty"`
	HasComplianceExport bool `json:"has_compliance_export,omitempty"`
	Icon *ObjsIcon `json:"icon"`
	Id *DefsTeam `json:"id"`
	IsAssigned bool `json:"is_assigned,omitempty"`
	IsEnterprise int32 `json:"is_enterprise,omitempty"`
	LimitTs int32 `json:"limit_ts,omitempty"`
	MessagesCount int32 `json:"messages_count,omitempty"`
	MsgEditWindowMins int32 `json:"msg_edit_window_mins,omitempty"`
	Name string `json:"name"`
	OverIntegrationsLimit bool `json:"over_integrations_limit,omitempty"`
	OverStorageLimit bool `json:"over_storage_limit,omitempty"`
	Plan string `json:"plan,omitempty"`
}