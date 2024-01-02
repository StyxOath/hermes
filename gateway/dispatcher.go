package gateway

type DispatchEvent[T any] struct {
	Payload T `json:"d"`
}

type Dispatch struct {
	// Application command permission was updated.
	ApplicationCommandPermissionsUpdate *chan ApplicationCommandPermissionsUpdateEvent

	// Auto Moderation rule was triggered and an action was executed (e.g. a message was blocked).
	AutoModerationActionExecution *chan AutoModerationActionExecutionEvent

	// Auto Moderation rule was created.
	AutoModerationRuleCreate *chan AutoModerationRuleCreateEvent

	// Auto Moderation rule was deleted.
	AutoModerationRuleDelete *chan AutoModerationRuleDeleteEvent

	// Auto Moderation rule was updated.
	AutoModerationRuleUpdate *chan AutoModerationRuleUpdateEvent

	// New guild channel created.
	ChannelCreate *chan ChannelCreateEvent

	// Channel was deleted.
	ChannelDelete *chan ChannelDeleteEvent

	// Message was pinned or unpinned.
	ChannelPinsUpdate *chan ChannelPinsUpdateEvent

	// Channel was updated.
	ChannelUpdate *chan ChannelUpdateEvent

	// Entitlement was created.
	EntitlementCreate *chan EntitlementCreateEvent

	// Entitlement was deleted.
	EntitlementDelete *chan EntitlementDeleteEvent

	// Entitlement was updated or renewed.
	EntitlementUpdate *chan EntitlementUpdateEvent

	// A guild audit log entry was created.
	GuildAuditLogEntryCreate *chan GuildAuditLogEntryCreateEvent

	// User was banned from a guild.
	GuildBanAdd *chan GuildBanAddEvent

	// User was unbanned from a guild.
	GuildBanRemove *chan GuildBanRemoveEvent

	// Lazy-load for unavailable guild, guild became available, or user joined a new guild.
	GuildCreate *chan GuildCreateEvent

	// Guild became unavailable, or user left/was removed from a guild.
	GuildDelete *chan GuildDeleteEvent

	// Guild emojis were updated.
	GuildEmojisUpdate *chan GuildEmojisUpdateEvent

	// Guild integration was updated.
	GuildIntegrationsUpdate *chan GuildIntegrationsUpdateEvent

	// New user joined a guild.
	GuildMemberAdd *chan GuildMemberAddEvent

	// User was removed from a guild.
	GuildMemberRemove *chan GuildMemberRemoveEvent

	// Guild member was updated.
	GuildMemberUpdate *chan GuildMemberUpdateEvent

	// Response to Request Guild Members.
	GuildMembersChunk *chan GuildMembersChunkEvent

	// Guild role was created.
	GuildRoleCreate *chan GuildRoleCreateEvent

	// Guild role was deleted.
	GuildRoleDelete *chan GuildRoleDeleteEvent

	// Guild role was updated.
	GuildRoleUpdate *chan GuildRoleUpdateEvent

	// Guild scheduled event was created.
	GuildScheduledEventCreate *chan GuildScheduledEventCreateEvent

	// Guild scheduled event was deleted.
	GuildScheduledEventDelete *chan GuildScheduledEventDeleteEvent

	// Guild scheduled event was updated.
	GuildScheduledEventUpdate *chan GuildScheduledEventUpdateEvent

	// User subscribed to a guild scheduled event.
	GuildScheduledEventUserAdd *chan GuildScheduledEventUserAddEvent

	// User unsubscribed from a guild scheduled event.
	GuildScheduledEventUserRemove *chan GuildScheduledEventUserRemoveEvent

	// Guild stickers were updated.
	GuildStickersUpdate *chan GuildStickersUpdateEvent

	// Guild was updated.
	GuildUpdate *chan GuildUpdateEvent

	// Defines the heartbeat interval.
	Hello *chan HelloEvent

	// Guild integration was created.
	IntegrationCreate *chan IntegrationCreateEvent

	// Guild integration was deleted.
	IntegrationDelete *chan IntegrationDeleteEvent

	// Guild integration was updated.
	IntegrationUpdate *chan IntegrationUpdateEvent

	// User used an interaction, such as an Application Command.
	InteractionCreate *chan InteractionCreateEvent

	// Failure response to Identify or Resume or invalid active session.
	InvalidSession *chan InvalidSessionEvent

	// Invite to a channel was created.
	InviteCreate *chan InviteCreateEvent

	// Invite to a channel was deleted.
	InviteDelete *chan InviteDeleteEvent

	// Message was created.
	MessageCreate *chan MessageCreateEvent

	// Message was deleted.
	MessageDelete *chan MessageDeleteEvent

	// Multiple messages were deleted at once.
	MessageDeleteBulk *chan MessageDeleteBulkEvent

	// User reacted to a message.
	MessageReactionAdd *chan MessageReactionAddEvent

	// User removed a reaction from a message.
	MessageReactionRemove *chan MessageReactionRemoveEvent

	// All reactions were explicitly removed from a message.
	MessageReactionRemoveAll *chan MessageReactionRemoveAllEvent

	// All reactions for a given emoji were explicitly removed from a message.
	MessageReactionRemoveEmoji *chan MessageReactionRemoveEmojiEvent

	// Message was edited.
	MessageUpdate *chan MessageUpdateEvent

	// Description.
	Name *chan NameEvent

	// User was updated.
	PresenceUpdate *chan PresenceUpdateEvent

	// Contains the initial state information.
	Ready *chan ReadyEvent

	// Server is going away, client should reconnect to gateway and resume.
	Reconnect *chan ReconnectEvent

	// Response to Resume.
	Resumed *chan ResumedEvent

	// Stage instance was created.
	StageInstanceCreate *chan StageInstanceCreateEvent

	// Stage instance was deleted or closed.
	StageInstanceDelete *chan StageInstanceDeleteEvent

	// Stage instance was updated.
	StageInstanceUpdate *chan StageInstanceUpdateEvent

	// Thread created, also sent when being added to a private thread.
	ThreadCreate *chan ThreadCreateEvent

	// Thread was deleted.
	ThreadDelete *chan ThreadDeleteEvent

	// Sent when gaining access to a channel, contains all active threads in that channel.
	ThreadListSync *chan ThreadListSyncEvent

	// Thread member for the current user was updated.
	ThreadMemberUpdate *chan ThreadMemberUpdateEvent

	// Some user(s) were added to or removed from a thread.
	ThreadMembersUpdate *chan ThreadMembersUpdateEvent

	// Thread was updated.
	ThreadUpdate *chan ThreadUpdateEvent

	// User started typing in a channel.
	TypingStart *chan TypingStartEvent

	// Properties about the user changed.
	UserUpdate *chan UserUpdateEvent

	// Guild's voice server was updated.
	VoiceServerUpdate *chan VoiceServerUpdateEvent

	// Someone joined, left, or moved a voice channel.
	VoiceStateUpdate *chan VoiceStateUpdateEvent

	// Guild channel webhook was created, update, or deleted.
	WebhooksUpdate *chan WebhooksUpdateEvent
}
