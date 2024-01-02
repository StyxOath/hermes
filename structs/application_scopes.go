package structs

type ApplicationScope string

var SCOPES = struct {
	// allows your app to fetch data from a user's "Now Playing/Recently Played" list — not currently available for apps.
	ActivitiesRead ApplicationScope

	// allows your app to update a user's activity - not currently available for apps (NOT REQUIRED FOR GAMESDK ACTIVITY MANAGER).
	ActivitiesWrite ApplicationScope

	// allows your app to read build data for a user's applications.
	ApplicationsBuildsRead ApplicationScope

	// allows your app to upload/update builds for a user's applications - requires Discord approval.
	ApplicationsBuildsUpload ApplicationScope

	// allows your app to add commands to a guild - included by default with the bot scope.
	ApplicationsCommands ApplicationScope

	// allows your app to update its commands using a Bearer token - client credentials grant only.
	ApplicationsCommandsUpdate ApplicationScope

	// allows your app to update permissions for its commands in a guild a user has permissions to.
	ApplicationsCommandsPermissionsUpdate ApplicationScope

	// allows your app to read entitlements for a user's applications.
	ApplicationsEntitlements ApplicationScope

	// allows your app to read and update store data (SKUs, store listings, achievements, etc.) for a user's applications.
	ApplicationsStoreUpdate ApplicationScope

	// for oauth2 bots, this puts the bot in the user's selected guild by default.
	Bot ApplicationScope

	// allows /users/@me/connections to return linked third-party accounts.
	Connections ApplicationScope

	// allows your app to see information about the user's DMs and group DMs - requires Discord approval.
	DmChannelsRead ApplicationScope

	// enables /users/@me to return an email.
	Email ApplicationScope

	// allows your app to join users to a group dm.
	GdmJoin ApplicationScope

	// allows /users/@me/guilds to return basic information about all of a user's guilds.
	Guilds ApplicationScope

	// allows /guilds/{guild.id}/members/{user.id} to be used for joining users to a guild.
	GuildsJoin ApplicationScope

	// allows /users/@me/guilds/{guild.id}/member to return a user's member information in a guild.
	GuildsMembersRead ApplicationScope

	// allows /users/@me without email.
	Identify ApplicationScope

	// for local rpc server api access, this allows you to read messages from all client channels (otherwise restricted to channels/guilds your app creates).
	MessagesRead ApplicationScope

	// allows your app to know a user's friends and implicit relationships - requires Discord approval.
	RelationshipsRead ApplicationScope

	// allows your app to update a user's connection and metadata for the app.
	RoleConnectionsWrite ApplicationScope

	// for local rpc server access, this allows you to control a user's local Discord client - requires Discord approval.
	Rpc ApplicationScope

	// for local rpc server access, this allows you to update a user's activity - requires Discord approval.
	RpcActivitiesWrite ApplicationScope

	// for local rpc server access, this allows you to receive notifications pushed out to the user - requires Discord approval.
	RpcNotificationsRead ApplicationScope

	// for local rpc server access, this allows you to read a user's voice settings and listen for voice events - requires Discord approval.
	RpcVoiceRead ApplicationScope

	// for local rpc server access, this allows you to update a user's voice settings - requires Discord approval.
	RpcVoiceWrite ApplicationScope

	// allows your app to connect to voice on user's behalf and see all the voice members - requires Discord approval.
	Voice ApplicationScope

	// this generates a webhook that is returned in the oauth token response for authorization code grants.
	WebhookIncoming ApplicationScope
} {
	ActivitiesRead: "activities.read",
	ActivitiesWrite: "activities.write",
	ApplicationsBuildsRead: "applications.builds.read",
	ApplicationsBuildsUpload: "applications.builds.upload",
	ApplicationsCommands: "applications.commands",
	ApplicationsCommandsUpdate: "applications.commands.update",
	ApplicationsCommandsPermissionsUpdate: "applications.commands.permissions.update",
	ApplicationsEntitlements: "applications.entitlements",
	ApplicationsStoreUpdate: "applications.store.update",
	Bot: "bot",
	Connections: "connections",
	DmChannelsRead: "dm_channels.read",
	Email: "email",
	GdmJoin: "gdm.join",
	Guilds: "guilds",
	GuildsJoin: "guilds.join",
	GuildsMembersRead: "guilds.members.read",
	Identify: "identify",
	MessagesRead: "messages.read",
	RelationshipsRead: "relationships.read",
	RoleConnectionsWrite: "role_connections.write",
	Rpc: "rpc",
	RpcActivitiesWrite: "rpc.activities.write",
	RpcNotificationsRead: "rpc.notifications.read",
	RpcVoiceRead: "rpc.voice.read",
	RpcVoiceWrite: "rpc.voice.write",
	Voice: "voice",
	WebhookIncoming: "webhook.incoming",
}

