/*
 * Slack Web API
 *
 * One way to interact with the Slack platform is its HTTP RPC-based Web API, a collection of methods requiring OAuth 2.0-based user, bot, or workspace tokens blessed with related OAuth scopes.
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ObjsUserProfile struct {
	AlwaysActive bool `json:"always_active,omitempty"`
	ApiAppId *DefsAppId `json:"api_app_id,omitempty"`
	AvatarHash string `json:"avatar_hash"`
	BotId *DefsBotId `json:"bot_id,omitempty"`
	DisplayName string `json:"display_name"`
	DisplayNameNormalized string `json:"display_name_normalized"`
	Email string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	GuestChannels string `json:"guest_channels,omitempty"`
	GuestExpirationTs int32 `json:"guest_expiration_ts,omitempty"`
	GuestInvitedBy string `json:"guest_invited_by,omitempty"`
	Image1024 string `json:"image_1024,omitempty"`
	Image192 string `json:"image_192,omitempty"`
	Image24 string `json:"image_24,omitempty"`
	Image32 string `json:"image_32,omitempty"`
	Image48 string `json:"image_48,omitempty"`
	Image512 string `json:"image_512,omitempty"`
	Image72 string `json:"image_72,omitempty"`
	ImageOriginal string `json:"image_original,omitempty"`
	IsCustomImage bool `json:"is_custom_image,omitempty"`
	LastName string `json:"last_name,omitempty"`
	Phone string `json:"phone,omitempty"`
	RealName string `json:"real_name"`
	RealNameNormalized string `json:"real_name_normalized"`
	Skype string `json:"skype,omitempty"`
	StatusEmoji string `json:"status_emoji,omitempty"`
	StatusExpiration int32 `json:"status_expiration,omitempty"`
	StatusText string `json:"status_text,omitempty"`
	StatusTextCanonical string `json:"status_text_canonical,omitempty"`
	Team *DefsWorkspaceId `json:"team,omitempty"`
	Teams *DefsWorkspaceId `json:"teams,omitempty"`
	Title string `json:"title,omitempty"`
}
