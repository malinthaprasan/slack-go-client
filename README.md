# Go API client for swagger

One way to interact with the Slack platform is its HTTP RPC-based Web API, a collection of methods requiring OAuth 2.0-based user, bot, or workspace tokens blessed with related OAuth scopes.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.2.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://slack.com/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminApi* | [**AdminUsersSessionReset**](docs/AdminApi.md#adminuserssessionreset) | **Post** /admin.users.session.reset | 
*AdminUsersSessionApi* | [**AdminUsersSessionReset**](docs/AdminUsersSessionApi.md#adminuserssessionreset) | **Post** /admin.users.session.reset | 
*ApiApi* | [**ApiTest**](docs/ApiApi.md#apitest) | **Get** /api.test | 
*AppsApi* | [**AppsPermissionsInfo**](docs/AppsApi.md#appspermissionsinfo) | **Get** /apps.permissions.info | 
*AppsApi* | [**AppsPermissionsRequest**](docs/AppsApi.md#appspermissionsrequest) | **Get** /apps.permissions.request | 
*AppsApi* | [**AppsPermissionsResourcesList**](docs/AppsApi.md#appspermissionsresourceslist) | **Get** /apps.permissions.resources.list | 
*AppsApi* | [**AppsPermissionsScopesList**](docs/AppsApi.md#appspermissionsscopeslist) | **Get** /apps.permissions.scopes.list | 
*AppsApi* | [**AppsPermissionsUsersList**](docs/AppsApi.md#appspermissionsuserslist) | **Get** /apps.permissions.users.list | 
*AppsApi* | [**AppsPermissionsUsersRequest**](docs/AppsApi.md#appspermissionsusersrequest) | **Get** /apps.permissions.users.request | 
*AppsApi* | [**AppsUninstall**](docs/AppsApi.md#appsuninstall) | **Get** /apps.uninstall | 
*AppsPermissionsApi* | [**AppsPermissionsInfo**](docs/AppsPermissionsApi.md#appspermissionsinfo) | **Get** /apps.permissions.info | 
*AppsPermissionsApi* | [**AppsPermissionsRequest**](docs/AppsPermissionsApi.md#appspermissionsrequest) | **Get** /apps.permissions.request | 
*AppsPermissionsResourcesApi* | [**AppsPermissionsResourcesList**](docs/AppsPermissionsResourcesApi.md#appspermissionsresourceslist) | **Get** /apps.permissions.resources.list | 
*AppsPermissionsScopesApi* | [**AppsPermissionsScopesList**](docs/AppsPermissionsScopesApi.md#appspermissionsscopeslist) | **Get** /apps.permissions.scopes.list | 
*AppsPermissionsUsersApi* | [**AppsPermissionsUsersList**](docs/AppsPermissionsUsersApi.md#appspermissionsuserslist) | **Get** /apps.permissions.users.list | 
*AppsPermissionsUsersApi* | [**AppsPermissionsUsersRequest**](docs/AppsPermissionsUsersApi.md#appspermissionsusersrequest) | **Get** /apps.permissions.users.request | 
*AuthApi* | [**AuthRevoke**](docs/AuthApi.md#authrevoke) | **Get** /auth.revoke | 
*AuthApi* | [**AuthTest**](docs/AuthApi.md#authtest) | **Get** /auth.test | 
*BotsApi* | [**BotsInfo**](docs/BotsApi.md#botsinfo) | **Get** /bots.info | 
*ChannelsApi* | [**ChannelsArchive**](docs/ChannelsApi.md#channelsarchive) | **Post** /channels.archive | 
*ChannelsApi* | [**ChannelsCreate**](docs/ChannelsApi.md#channelscreate) | **Post** /channels.create | 
*ChannelsApi* | [**ChannelsHistory**](docs/ChannelsApi.md#channelshistory) | **Get** /channels.history | 
*ChannelsApi* | [**ChannelsInfo**](docs/ChannelsApi.md#channelsinfo) | **Get** /channels.info | 
*ChannelsApi* | [**ChannelsInvite**](docs/ChannelsApi.md#channelsinvite) | **Post** /channels.invite | 
*ChannelsApi* | [**ChannelsJoin**](docs/ChannelsApi.md#channelsjoin) | **Post** /channels.join | 
*ChannelsApi* | [**ChannelsKick**](docs/ChannelsApi.md#channelskick) | **Post** /channels.kick | 
*ChannelsApi* | [**ChannelsLeave**](docs/ChannelsApi.md#channelsleave) | **Post** /channels.leave | 
*ChannelsApi* | [**ChannelsList**](docs/ChannelsApi.md#channelslist) | **Get** /channels.list | 
*ChannelsApi* | [**ChannelsMark**](docs/ChannelsApi.md#channelsmark) | **Post** /channels.mark | 
*ChannelsApi* | [**ChannelsRename**](docs/ChannelsApi.md#channelsrename) | **Post** /channels.rename | 
*ChannelsApi* | [**ChannelsReplies**](docs/ChannelsApi.md#channelsreplies) | **Get** /channels.replies | 
*ChannelsApi* | [**ChannelsSetPurpose**](docs/ChannelsApi.md#channelssetpurpose) | **Post** /channels.setPurpose | 
*ChannelsApi* | [**ChannelsSetTopic**](docs/ChannelsApi.md#channelssettopic) | **Post** /channels.setTopic | 
*ChannelsApi* | [**ChannelsUnarchive**](docs/ChannelsApi.md#channelsunarchive) | **Post** /channels.unarchive | 
*ChatApi* | [**ChatDelete**](docs/ChatApi.md#chatdelete) | **Post** /chat.delete | 
*ChatApi* | [**ChatDeleteScheduledMessage**](docs/ChatApi.md#chatdeletescheduledmessage) | **Post** /chat.deleteScheduledMessage | 
*ChatApi* | [**ChatGetPermalink**](docs/ChatApi.md#chatgetpermalink) | **Get** /chat.getPermalink | 
*ChatApi* | [**ChatMeMessage**](docs/ChatApi.md#chatmemessage) | **Post** /chat.meMessage | 
*ChatApi* | [**ChatPostEphemeral**](docs/ChatApi.md#chatpostephemeral) | **Post** /chat.postEphemeral | 
*ChatApi* | [**ChatPostMessage**](docs/ChatApi.md#chatpostmessage) | **Post** /chat.postMessage | 
*ChatApi* | [**ChatScheduleMessage**](docs/ChatApi.md#chatschedulemessage) | **Post** /chat.scheduleMessage | 
*ChatApi* | [**ChatScheduledMessagesList**](docs/ChatApi.md#chatscheduledmessageslist) | **Get** /chat.scheduledMessages.list | 
*ChatApi* | [**ChatUnfurl**](docs/ChatApi.md#chatunfurl) | **Post** /chat.unfurl | 
*ChatApi* | [**ChatUpdate**](docs/ChatApi.md#chatupdate) | **Post** /chat.update | 
*ChatScheduledMessagesApi* | [**ChatScheduledMessagesList**](docs/ChatScheduledMessagesApi.md#chatscheduledmessageslist) | **Get** /chat.scheduledMessages.list | 
*ConversationsApi* | [**ConversationsArchive**](docs/ConversationsApi.md#conversationsarchive) | **Post** /conversations.archive | 
*ConversationsApi* | [**ConversationsClose**](docs/ConversationsApi.md#conversationsclose) | **Post** /conversations.close | 
*ConversationsApi* | [**ConversationsCreate**](docs/ConversationsApi.md#conversationscreate) | **Post** /conversations.create | 
*ConversationsApi* | [**ConversationsHistory**](docs/ConversationsApi.md#conversationshistory) | **Get** /conversations.history | 
*ConversationsApi* | [**ConversationsInfo**](docs/ConversationsApi.md#conversationsinfo) | **Get** /conversations.info | 
*ConversationsApi* | [**ConversationsInvite**](docs/ConversationsApi.md#conversationsinvite) | **Post** /conversations.invite | 
*ConversationsApi* | [**ConversationsJoin**](docs/ConversationsApi.md#conversationsjoin) | **Post** /conversations.join | 
*ConversationsApi* | [**ConversationsKick**](docs/ConversationsApi.md#conversationskick) | **Post** /conversations.kick | 
*ConversationsApi* | [**ConversationsLeave**](docs/ConversationsApi.md#conversationsleave) | **Post** /conversations.leave | 
*ConversationsApi* | [**ConversationsList**](docs/ConversationsApi.md#conversationslist) | **Get** /conversations.list | 
*ConversationsApi* | [**ConversationsMembers**](docs/ConversationsApi.md#conversationsmembers) | **Get** /conversations.members | 
*ConversationsApi* | [**ConversationsOpen**](docs/ConversationsApi.md#conversationsopen) | **Post** /conversations.open | 
*ConversationsApi* | [**ConversationsRename**](docs/ConversationsApi.md#conversationsrename) | **Post** /conversations.rename | 
*ConversationsApi* | [**ConversationsReplies**](docs/ConversationsApi.md#conversationsreplies) | **Get** /conversations.replies | 
*ConversationsApi* | [**ConversationsSetPurpose**](docs/ConversationsApi.md#conversationssetpurpose) | **Post** /conversations.setPurpose | 
*ConversationsApi* | [**ConversationsSetTopic**](docs/ConversationsApi.md#conversationssettopic) | **Post** /conversations.setTopic | 
*ConversationsApi* | [**ConversationsUnarchive**](docs/ConversationsApi.md#conversationsunarchive) | **Post** /conversations.unarchive | 
*DialogApi* | [**DialogOpen**](docs/DialogApi.md#dialogopen) | **Get** /dialog.open | 
*DndApi* | [**DndEndDnd**](docs/DndApi.md#dndenddnd) | **Post** /dnd.endDnd | 
*DndApi* | [**DndEndSnooze**](docs/DndApi.md#dndendsnooze) | **Post** /dnd.endSnooze | 
*DndApi* | [**DndInfo**](docs/DndApi.md#dndinfo) | **Get** /dnd.info | 
*DndApi* | [**DndSetSnooze**](docs/DndApi.md#dndsetsnooze) | **Post** /dnd.setSnooze | 
*DndApi* | [**DndTeamInfo**](docs/DndApi.md#dndteaminfo) | **Get** /dnd.teamInfo | 
*EmojiApi* | [**EmojiList**](docs/EmojiApi.md#emojilist) | **Get** /emoji.list | 
*FilesApi* | [**FilesCommentsDelete**](docs/FilesApi.md#filescommentsdelete) | **Post** /files.comments.delete | 
*FilesApi* | [**FilesDelete**](docs/FilesApi.md#filesdelete) | **Post** /files.delete | 
*FilesApi* | [**FilesInfo**](docs/FilesApi.md#filesinfo) | **Get** /files.info | 
*FilesApi* | [**FilesList**](docs/FilesApi.md#fileslist) | **Get** /files.list | 
*FilesApi* | [**FilesRevokePublicURL**](docs/FilesApi.md#filesrevokepublicurl) | **Post** /files.revokePublicURL | 
*FilesApi* | [**FilesSharedPublicURL**](docs/FilesApi.md#filessharedpublicurl) | **Post** /files.sharedPublicURL | 
*FilesApi* | [**FilesUpload**](docs/FilesApi.md#filesupload) | **Post** /files.upload | 
*FilesCommentsApi* | [**FilesCommentsDelete**](docs/FilesCommentsApi.md#filescommentsdelete) | **Post** /files.comments.delete | 
*GroupsApi* | [**GroupsArchive**](docs/GroupsApi.md#groupsarchive) | **Post** /groups.archive | 
*GroupsApi* | [**GroupsCreate**](docs/GroupsApi.md#groupscreate) | **Post** /groups.create | 
*GroupsApi* | [**GroupsCreateChild**](docs/GroupsApi.md#groupscreatechild) | **Post** /groups.createChild | 
*GroupsApi* | [**GroupsHistory**](docs/GroupsApi.md#groupshistory) | **Get** /groups.history | 
*GroupsApi* | [**GroupsInfo**](docs/GroupsApi.md#groupsinfo) | **Get** /groups.info | 
*GroupsApi* | [**GroupsInvite**](docs/GroupsApi.md#groupsinvite) | **Post** /groups.invite | 
*GroupsApi* | [**GroupsKick**](docs/GroupsApi.md#groupskick) | **Post** /groups.kick | 
*GroupsApi* | [**GroupsLeave**](docs/GroupsApi.md#groupsleave) | **Post** /groups.leave | 
*GroupsApi* | [**GroupsList**](docs/GroupsApi.md#groupslist) | **Get** /groups.list | 
*GroupsApi* | [**GroupsMark**](docs/GroupsApi.md#groupsmark) | **Post** /groups.mark | 
*GroupsApi* | [**GroupsOpen**](docs/GroupsApi.md#groupsopen) | **Post** /groups.open | 
*GroupsApi* | [**GroupsRename**](docs/GroupsApi.md#groupsrename) | **Post** /groups.rename | 
*GroupsApi* | [**GroupsReplies**](docs/GroupsApi.md#groupsreplies) | **Get** /groups.replies | 
*GroupsApi* | [**GroupsSetPurpose**](docs/GroupsApi.md#groupssetpurpose) | **Post** /groups.setPurpose | 
*GroupsApi* | [**GroupsSetTopic**](docs/GroupsApi.md#groupssettopic) | **Post** /groups.setTopic | 
*GroupsApi* | [**GroupsUnarchive**](docs/GroupsApi.md#groupsunarchive) | **Post** /groups.unarchive | 
*ImApi* | [**ImClose**](docs/ImApi.md#imclose) | **Post** /im.close | 
*ImApi* | [**ImHistory**](docs/ImApi.md#imhistory) | **Get** /im.history | 
*ImApi* | [**ImList**](docs/ImApi.md#imlist) | **Get** /im.list | 
*ImApi* | [**ImMark**](docs/ImApi.md#immark) | **Post** /im.mark | 
*ImApi* | [**ImOpen**](docs/ImApi.md#imopen) | **Post** /im.open | 
*ImApi* | [**ImReplies**](docs/ImApi.md#imreplies) | **Get** /im.replies | 
*MigrationApi* | [**MigrationExchange**](docs/MigrationApi.md#migrationexchange) | **Get** /migration.exchange | 
*MpimApi* | [**MpimClose**](docs/MpimApi.md#mpimclose) | **Post** /mpim.close | 
*MpimApi* | [**MpimHistory**](docs/MpimApi.md#mpimhistory) | **Get** /mpim.history | 
*MpimApi* | [**MpimList**](docs/MpimApi.md#mpimlist) | **Get** /mpim.list | 
*MpimApi* | [**MpimMark**](docs/MpimApi.md#mpimmark) | **Post** /mpim.mark | 
*MpimApi* | [**MpimOpen**](docs/MpimApi.md#mpimopen) | **Post** /mpim.open | 
*MpimApi* | [**MpimReplies**](docs/MpimApi.md#mpimreplies) | **Get** /mpim.replies | 
*OauthApi* | [**OauthAccess**](docs/OauthApi.md#oauthaccess) | **Get** /oauth.access | 
*OauthApi* | [**OauthToken**](docs/OauthApi.md#oauthtoken) | **Get** /oauth.token | 
*PinsApi* | [**PinsAdd**](docs/PinsApi.md#pinsadd) | **Post** /pins.add | 
*PinsApi* | [**PinsList**](docs/PinsApi.md#pinslist) | **Get** /pins.list | 
*PinsApi* | [**PinsRemove**](docs/PinsApi.md#pinsremove) | **Post** /pins.remove | 
*ReactionsApi* | [**ReactionsAdd**](docs/ReactionsApi.md#reactionsadd) | **Post** /reactions.add | 
*ReactionsApi* | [**ReactionsGet**](docs/ReactionsApi.md#reactionsget) | **Get** /reactions.get | 
*ReactionsApi* | [**ReactionsList**](docs/ReactionsApi.md#reactionslist) | **Get** /reactions.list | 
*ReactionsApi* | [**ReactionsRemove**](docs/ReactionsApi.md#reactionsremove) | **Post** /reactions.remove | 
*RemindersApi* | [**RemindersAdd**](docs/RemindersApi.md#remindersadd) | **Post** /reminders.add | 
*RemindersApi* | [**RemindersComplete**](docs/RemindersApi.md#reminderscomplete) | **Post** /reminders.complete | 
*RemindersApi* | [**RemindersDelete**](docs/RemindersApi.md#remindersdelete) | **Post** /reminders.delete | 
*RemindersApi* | [**RemindersInfo**](docs/RemindersApi.md#remindersinfo) | **Get** /reminders.info | 
*RemindersApi* | [**RemindersList**](docs/RemindersApi.md#reminderslist) | **Get** /reminders.list | 
*RtmApi* | [**RtmConnect**](docs/RtmApi.md#rtmconnect) | **Get** /rtm.connect | 
*SearchApi* | [**SearchMessages**](docs/SearchApi.md#searchmessages) | **Get** /search.messages | 
*StarsApi* | [**StarsAdd**](docs/StarsApi.md#starsadd) | **Post** /stars.add | 
*StarsApi* | [**StarsList**](docs/StarsApi.md#starslist) | **Get** /stars.list | 
*StarsApi* | [**StarsRemove**](docs/StarsApi.md#starsremove) | **Post** /stars.remove | 
*TeamApi* | [**TeamAccessLogs**](docs/TeamApi.md#teamaccesslogs) | **Get** /team.accessLogs | 
*TeamApi* | [**TeamBillableInfo**](docs/TeamApi.md#teambillableinfo) | **Get** /team.billableInfo | 
*TeamApi* | [**TeamInfo**](docs/TeamApi.md#teaminfo) | **Get** /team.info | 
*TeamApi* | [**TeamIntegrationLogs**](docs/TeamApi.md#teamintegrationlogs) | **Get** /team.integrationLogs | 
*TeamApi* | [**TeamProfileGet**](docs/TeamApi.md#teamprofileget) | **Get** /team.profile.get | 
*TeamProfileApi* | [**TeamProfileGet**](docs/TeamProfileApi.md#teamprofileget) | **Get** /team.profile.get | 
*UsergroupsApi* | [**UsergroupsCreate**](docs/UsergroupsApi.md#usergroupscreate) | **Post** /usergroups.create | 
*UsergroupsApi* | [**UsergroupsDisable**](docs/UsergroupsApi.md#usergroupsdisable) | **Post** /usergroups.disable | 
*UsergroupsApi* | [**UsergroupsEnable**](docs/UsergroupsApi.md#usergroupsenable) | **Post** /usergroups.enable | 
*UsergroupsApi* | [**UsergroupsList**](docs/UsergroupsApi.md#usergroupslist) | **Get** /usergroups.list | 
*UsergroupsApi* | [**UsergroupsUpdate**](docs/UsergroupsApi.md#usergroupsupdate) | **Post** /usergroups.update | 
*UsergroupsApi* | [**UsergroupsUsersList**](docs/UsergroupsApi.md#usergroupsuserslist) | **Get** /usergroups.users.list | 
*UsergroupsApi* | [**UsergroupsUsersUpdate**](docs/UsergroupsApi.md#usergroupsusersupdate) | **Post** /usergroups.users.update | 
*UsergroupsUsersApi* | [**UsergroupsUsersList**](docs/UsergroupsUsersApi.md#usergroupsuserslist) | **Get** /usergroups.users.list | 
*UsergroupsUsersApi* | [**UsergroupsUsersUpdate**](docs/UsergroupsUsersApi.md#usergroupsusersupdate) | **Post** /usergroups.users.update | 
*UsersApi* | [**UsersConversations**](docs/UsersApi.md#usersconversations) | **Get** /users.conversations | 
*UsersApi* | [**UsersDeletePhoto**](docs/UsersApi.md#usersdeletephoto) | **Post** /users.deletePhoto | 
*UsersApi* | [**UsersGetPresence**](docs/UsersApi.md#usersgetpresence) | **Get** /users.getPresence | 
*UsersApi* | [**UsersIdentity**](docs/UsersApi.md#usersidentity) | **Get** /users.identity | 
*UsersApi* | [**UsersInfo**](docs/UsersApi.md#usersinfo) | **Get** /users.info | 
*UsersApi* | [**UsersList**](docs/UsersApi.md#userslist) | **Get** /users.list | 
*UsersApi* | [**UsersLookupByEmail**](docs/UsersApi.md#userslookupbyemail) | **Get** /users.lookupByEmail | 
*UsersApi* | [**UsersProfileGet**](docs/UsersApi.md#usersprofileget) | **Get** /users.profile.get | 
*UsersApi* | [**UsersProfileSet**](docs/UsersApi.md#usersprofileset) | **Post** /users.profile.set | 
*UsersApi* | [**UsersSetActive**](docs/UsersApi.md#userssetactive) | **Post** /users.setActive | 
*UsersApi* | [**UsersSetPhoto**](docs/UsersApi.md#userssetphoto) | **Post** /users.setPhoto | 
*UsersApi* | [**UsersSetPresence**](docs/UsersApi.md#userssetpresence) | **Post** /users.setPresence | 
*UsersProfileApi* | [**UsersProfileGet**](docs/UsersProfileApi.md#usersprofileget) | **Get** /users.profile.get | 
*UsersProfileApi* | [**UsersProfileSet**](docs/UsersProfileApi.md#usersprofileset) | **Post** /users.profile.set | 


## Documentation For Models

 - [Blocks](docs/Blocks.md)
 - [BlocksInner](docs/BlocksInner.md)
 - [DefaultSuccessTemplate](docs/DefaultSuccessTemplate.md)
 - [DefsAppId](docs/DefsAppId.md)
 - [DefsBotId](docs/DefsBotId.md)
 - [DefsChannel](docs/DefsChannel.md)
 - [DefsChannelId](docs/DefsChannelId.md)
 - [DefsChannelName](docs/DefsChannelName.md)
 - [DefsCommentId](docs/DefsCommentId.md)
 - [DefsDmId](docs/DefsDmId.md)
 - [DefsEnterpriseId](docs/DefsEnterpriseId.md)
 - [DefsEnterpriseName](docs/DefsEnterpriseName.md)
 - [DefsEnterpriseUserId](docs/DefsEnterpriseUserId.md)
 - [DefsFileId](docs/DefsFileId.md)
 - [DefsGroupId](docs/DefsGroupId.md)
 - [DefsOkFalse](docs/DefsOkFalse.md)
 - [DefsOkTrue](docs/DefsOkTrue.md)
 - [DefsPinnedInfo](docs/DefsPinnedInfo.md)
 - [DefsReminderId](docs/DefsReminderId.md)
 - [DefsSubteamId](docs/DefsSubteamId.md)
 - [DefsTeam](docs/DefsTeam.md)
 - [DefsTopicPurposeCreator](docs/DefsTopicPurposeCreator.md)
 - [DefsTs](docs/DefsTs.md)
 - [DefsUserId](docs/DefsUserId.md)
 - [DefsWorkspaceId](docs/DefsWorkspaceId.md)
 - [ObjsChannel](docs/ObjsChannel.md)
 - [ObjsChannelPurpose](docs/ObjsChannelPurpose.md)
 - [ObjsComment](docs/ObjsComment.md)
 - [ObjsComments](docs/ObjsComments.md)
 - [ObjsConversation](docs/ObjsConversation.md)
 - [ObjsEnterpriseUser](docs/ObjsEnterpriseUser.md)
 - [ObjsFile](docs/ObjsFile.md)
 - [ObjsFileShares](docs/ObjsFileShares.md)
 - [ObjsGroup](docs/ObjsGroup.md)
 - [ObjsIcon](docs/ObjsIcon.md)
 - [ObjsIm](docs/ObjsIm.md)
 - [ObjsMessage](docs/ObjsMessage.md)
 - [ObjsMessageAttachments](docs/ObjsMessageAttachments.md)
 - [ObjsMessageIcons](docs/ObjsMessageIcons.md)
 - [ObjsMessageReplies](docs/ObjsMessageReplies.md)
 - [ObjsPaging](docs/ObjsPaging.md)
 - [ObjsReaction](docs/ObjsReaction.md)
 - [ObjsReminder](docs/ObjsReminder.md)
 - [ObjsResources](docs/ObjsResources.md)
 - [ObjsResponseMetadata](docs/ObjsResponseMetadata.md)
 - [ObjsScopes](docs/ObjsScopes.md)
 - [ObjsSubteam](docs/ObjsSubteam.md)
 - [ObjsSubteamPrefs](docs/ObjsSubteamPrefs.md)
 - [ObjsTeam](docs/ObjsTeam.md)
 - [ObjsTeamProfileField](docs/ObjsTeamProfileField.md)
 - [ObjsUser](docs/ObjsUser.md)
 - [ObjsUserProfile](docs/ObjsUserProfile.md)
 - [ObjsUserProfileShort](docs/ObjsUserProfileShort.md)


## Documentation For Authorization

## slackAuth
- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://slack.com/oauth/authorize
- **Scopes**: 
 - **admin**: admin
 - **admin.users:write**: admin.users:write
 - **bot**: Bot user scope
 - **channels:history**: channels:history
 - **channels:read**: channels:read
 - **channels:write**: channels:write
 - **chat:write**: chat:write
 - **chat:write:bot**: Author messages as a bot
 - **chat:write:user**: chat:write:user
 - **conversations:history**: conversations:history
 - **conversations:read**: conversations:read
 - **conversations:write**: conversations:write
 - **dnd:read**: dnd:read
 - **dnd:write**: dnd:write
 - **emoji:read**: emoji:read
 - **files:read**: files:read
 - **files:write:user**: files:write:user
 - **groups:history**: groups:history
 - **groups:read**: groups:read
 - **groups:write**: groups:write
 - **identity.basic**: identity.basic
 - **im:history**: im:history
 - **im:read**: im:read
 - **im:write**: im:write
 - **links:write**: links:write
 - **mpim:history**: mpim:history
 - **mpim:read**: mpim:read
 - **mpim:write**: mpim:write
 - **none**: No scope required
 - **pins:read**: pins:read
 - **pins:write**: pins:write
 - **reactions:read**: reactions:read
 - **reactions:write**: reactions:write
 - **reminders:read**: reminders:read
 - **reminders:write**: reminders:write
 - **rtm:stream**: rtm:stream
 - **search:read**: search:read
 - **stars:read**: stars:read
 - **stars:write**: stars:write
 - **team:read**: team:read
 - **tokens.basic**: tokens.basic
 - **usergroups:read**: usergroups:read
 - **usergroups:write**: usergroups:write
 - **users.profile:read**: users.profile:read
 - **users.profile:write**: users.profile:write
 - **users:read**: users:read
 - **users:read.email**: users:read.email
 - **users:write**: users:write

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.
```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

## Author



