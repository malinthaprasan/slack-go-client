/*
 * Slack Web API
 *
 * One way to interact with the Slack platform is its HTTP RPC-based Web API, a collection of methods requiring OAuth 2.0-based user, bot, or workspace tokens blessed with related OAuth scopes.
 *
 * API version: 1.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ObjsMessage struct {
	Attachments []ObjsMessageAttachments `json:"attachments,omitempty"`
	Blocks *Blocks `json:"blocks,omitempty"`
	BotId *interface{} `json:"bot_id,omitempty"`
	ClientMsgId string `json:"client_msg_id,omitempty"`
	Comment *ObjsComment `json:"comment,omitempty"`
	DisplayAsBot bool `json:"display_as_bot,omitempty"`
	File *ObjsFile `json:"file,omitempty"`
	Files []ObjsFile `json:"files,omitempty"`
	Icons *ObjsMessageIcons `json:"icons,omitempty"`
	Inviter *DefsUserId `json:"inviter,omitempty"`
	IsDelayedMessage bool `json:"is_delayed_message,omitempty"`
	IsIntro bool `json:"is_intro,omitempty"`
	IsStarred bool `json:"is_starred,omitempty"`
	LastRead *DefsTs `json:"last_read,omitempty"`
	LatestReply *DefsTs `json:"latest_reply,omitempty"`
	Name string `json:"name,omitempty"`
	OldName string `json:"old_name,omitempty"`
	ParentUserId *DefsUserId `json:"parent_user_id,omitempty"`
	Permalink string `json:"permalink,omitempty"`
	PinnedTo []DefsChannel `json:"pinned_to,omitempty"`
	Purpose string `json:"purpose,omitempty"`
	Reactions []ObjsReaction `json:"reactions,omitempty"`
	Replies []ObjsMessageReplies `json:"replies,omitempty"`
	ReplyCount int32 `json:"reply_count,omitempty"`
	ReplyUsers []DefsUserId `json:"reply_users,omitempty"`
	ReplyUsersCount int32 `json:"reply_users_count,omitempty"`
	SourceTeam *DefsWorkspaceId `json:"source_team,omitempty"`
	Subscribed bool `json:"subscribed,omitempty"`
	Subtype string `json:"subtype,omitempty"`
	Team *DefsWorkspaceId `json:"team,omitempty"`
	Text string `json:"text"`
	ThreadTs *DefsTs `json:"thread_ts,omitempty"`
	Topic string `json:"topic,omitempty"`
	Ts *DefsTs `json:"ts"`
	Type_ string `json:"type"`
	UnreadCount int32 `json:"unread_count,omitempty"`
	Upload bool `json:"upload,omitempty"`
	User *DefsUserId `json:"user,omitempty"`
	UserProfile *ObjsUserProfileShort `json:"user_profile,omitempty"`
	UserTeam *DefsWorkspaceId `json:"user_team,omitempty"`
	Username string `json:"username,omitempty"`
}