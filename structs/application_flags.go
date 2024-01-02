package structs

type ApplicationFlags int32

var APPLICATION_FLAGS = &struct {
	// Indicates if an app uses the Auto Moderation API.
	ApplicationAutoModerationRuleCreateBadge ApplicationFlags
	
	// Intent required for bots in 100 or more servers to receive presence_update events.
	GatewayPresence ApplicationFlags
	
	// Intent required for bots in under 100 servers to receive presence_update events, found on the Bot page in your app's settings.
	GatewayPresenceLimited ApplicationFlags
	
	// Intent required for bots in 100 or more servers to receive member-related events like guild_member_add. See the list of member-related events under GUILD_MEMBERS.
	GatewayGuildMembers ApplicationFlags
	
	// Intent required for bots in under 100 servers to receive member-related events like guild_member_add, found on the Bot page in your app's settings. See the list of member-related events under GUILD_MEMBERS.
	GatewayGuildMembersLimited ApplicationFlags
	
	// Indicates unusual growth of an app that prevents verification.
	VerificationPendingGuildLimit ApplicationFlags
	
	// Indicates if an app is embedded within the Discord client (currently unavailable publicly).
	Embedded ApplicationFlags
	
	// Intent required for bots in 100 or more servers to receive message content.
	GatewayMessageContent ApplicationFlags
	
	// Intent required for bots in under 100 servers to receive message content, found on the Bot page in your app's settings.
	GatewayMessageContentLimited ApplicationFlags
	
	// Indicates if an app has registered global application commands.
	ApplicationCommandBadge ApplicationFlags
} {
	ApplicationAutoModerationRuleCreateBadge: 1 << 6,
	GatewayPresence: 1 << 12,
	GatewayPresenceLimited: 1 << 13,
	GatewayGuildMembers: 1 << 14,
	GatewayGuildMembersLimited: 1 << 15,
	VerificationPendingGuildLimit: 1 << 16,
	Embedded: 1 << 17,
	GatewayMessageContent: 1 << 18,
	GatewayMessageContentLimited: 1 << 19,
	ApplicationCommandBadge: 1 << 23,
}

