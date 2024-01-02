package structs

type AuditLogEntry struct {
	// Type of action that occurred.
	ActionType AuditLogEvent `json:"action_type"`

	// Changes made to the target_id.
	Changes *[]AuditLogChange `json:"changes,omitempty"`

	// ID of the entry.
	Id Snowflake `json:"id"`

	// Additional info for certain event types.
	Options *AuditEntryInfo `json:"options,omitempty"`

	// Reason for the change (1-512 characters).
	Reason *string `json:"reason,omitempty"`

	// ID of the affected entity (webhook, user, role, etc.).
	TargetId *string `json:"target_id,omitempty"`

	// User or app that made the changes.
	UserId *Snowflake `json:"user_id,omitempty"`
}

type AuditLogChange []interface{}

type AuditEntryInfo struct {
	// ID of the app whose permissions were targeted.
	ApplicationId *Snowflake `json:"application_id,omitempty"`

	// Name of the Auto Moderation rule that was triggered.
	AutoModerationRuleName *string `json:"auto_moderation_rule_name,omitempty"`

	// Trigger type of the Auto Moderation rule that was triggered.
	AutoModerationRuleTriggerType *string `json:"auto_moderation_rule_trigger_type,omitempty"`

	// Channel in which the entities were targeted.
	ChannelId *Snowflake `json:"channel_id,omitempty"`

	// Number of entities that were targeted.
	Count *string `json:"count,omitempty"`

	// Number of days after which inactive members were kicked.
	DeleteMemberDays *string `json:"delete_member_days,omitempty"`

	// ID of the overwritten entity.
	Id *Snowflake `json:"id,omitempty"`

	// The type of integration which performed the action.
	IntegrationType *string `json:"integration_type,omitempty"`

	// Number of members removed by the prune.
	MembersRemoved *string `json:"members_removed,omitempty"`

	// ID of the message that was targeted.
	MessageId *Snowflake `json:"message_id,omitempty"`

	// Name of the role if type is "0" (not present if type is "1").
	RoleName *string `json:"role_name,omitempty"`

	// Type of overwritten entity - role ("0") or member ("1").
	Type *string `json:"type,omitempty"`
}

type AuditLogEvent int16

var AUDIT_LOG_EVENTS = struct {
	// Server settings were updated.
	GuildUpdate AuditLogEvent

	// Channel was created.
	ChannelCreate AuditLogEvent

	// Channel settings were updated.
	ChannelUpdate AuditLogEvent

	// Channel was deleted.
	ChannelDelete AuditLogEvent

	// Permission overwrite was added to a channel.
	ChannelOverwriteCreate AuditLogEvent

	// Permission overwrite was updated for a channel.
	ChannelOverwriteUpdate AuditLogEvent

	// Permission overwrite was deleted from a channel.
	ChannelOverwriteDelete AuditLogEvent

	// Member was removed from server.
	MemberKick AuditLogEvent

	// Members were pruned from server.
	MemberPrune AuditLogEvent

	// Member was banned from server.
	MemberBanAdd AuditLogEvent

	// Server ban was lifted for a member.
	MemberBanRemove AuditLogEvent

	// Member was updated in server.
	MemberUpdate AuditLogEvent

	// Member was added or removed from a role.
	MemberRoleUpdate AuditLogEvent

	// Member was moved to a different voice channel.
	MemberMove AuditLogEvent

	// Member was disconnected from a voice channel.
	MemberDisconnect AuditLogEvent

	// Bot user was added to server.
	BotAdd AuditLogEvent

	// Role was created.
	RoleCreate AuditLogEvent

	// Role was edited.
	RoleUpdate AuditLogEvent

	// Role was deleted.
	RoleDelete AuditLogEvent

	// Server invite was created.
	InviteCreate AuditLogEvent

	// Server invite was updated.
	InviteUpdate AuditLogEvent

	// Server invite was deleted.
	InviteDelete AuditLogEvent

	// Webhook was created.
	WebhookCreate AuditLogEvent

	// Webhook properties or channel were updated.
	WebhookUpdate AuditLogEvent

	// Webhook was deleted.
	WebhookDelete AuditLogEvent

	// Emoji was created.
	EmojiCreate AuditLogEvent

	// Emoji name was updated.
	EmojiUpdate AuditLogEvent

	// Emoji was deleted.
	EmojiDelete AuditLogEvent

	// Single message was deleted.
	MessageDelete AuditLogEvent

	// Multiple messages were deleted.
	MessageBulkDelete AuditLogEvent

	// Message was pinned to a channel.
	MessagePin AuditLogEvent

	// Message was unpinned from a channel.
	MessageUnpin AuditLogEvent

	// App was added to server.
	IntegrationCreate AuditLogEvent

	// App was updated (as an example, its scopes were updated).
	IntegrationUpdate AuditLogEvent

	// App was removed from server.
	IntegrationDelete AuditLogEvent

	// Stage instance was created (stage channel becomes live).
	StageInstanceCreate AuditLogEvent

	// Stage instance details were updated.
	StageInstanceUpdate AuditLogEvent

	// Stage instance was deleted (stage channel no longer live).
	StageInstanceDelete AuditLogEvent

	// Sticker was created.
	StickerCreate AuditLogEvent

	// Sticker details were updated.
	StickerUpdate AuditLogEvent

	// Sticker was deleted.
	StickerDelete AuditLogEvent

	// Event was created.
	GuildScheduledEventCreate AuditLogEvent

	// Event was updated.
	GuildScheduledEventUpdate AuditLogEvent

	// Event was cancelled.
	GuildScheduledEventDelete AuditLogEvent

	// Thread was created in a channel.
	ThreadCreate AuditLogEvent

	// Thread was updated.
	ThreadUpdate AuditLogEvent

	// Thread was deleted.
	ThreadDelete AuditLogEvent

	// Permissions were updated for a command.
	ApplicationCommandPermissionUpdate AuditLogEvent

	// Auto Moderation rule was created.
	AutoModerationRuleCreate AuditLogEvent

	// Auto Moderation rule was updated.
	AutoModerationRuleUpdate AuditLogEvent

	// Auto Moderation rule was deleted.
	AutoModerationRuleDelete AuditLogEvent

	// Message was blocked by Auto Moderation.
	AutoModerationBlockMessage AuditLogEvent

	// Message was flagged by Auto Moderation.
	AutoModerationFlagToChannel AuditLogEvent

	// Member was timed out by Auto Moderation.
	AutoModerationUserCommunicationDisabled AuditLogEvent

	// Creator monetization request was created.
	CreatorMonetizationRequestCreated AuditLogEvent

	// Creator monetization terms were accepted.
	CreatorMonetizationTermsAccepted AuditLogEvent
} {
	GuildUpdate: 1,
	ChannelCreate: 10,
	ChannelUpdate: 11,
	ChannelDelete: 12,
	ChannelOverwriteCreate: 13,
	ChannelOverwriteUpdate: 14,
	ChannelOverwriteDelete: 15,
	MemberKick: 20,
	MemberPrune: 21,
	MemberBanAdd: 22,
	MemberBanRemove: 23,
	MemberUpdate: 24,
	MemberRoleUpdate: 25,
	MemberMove: 26,
	MemberDisconnect: 27,
	BotAdd: 28,
	RoleCreate: 30,
	RoleUpdate: 31,
	RoleDelete: 32,
	InviteCreate: 40,
	InviteUpdate: 41,
	InviteDelete: 42,
	WebhookCreate: 50,
	WebhookUpdate: 51,
	WebhookDelete: 52,
	EmojiCreate: 60,
	EmojiUpdate: 61,
	EmojiDelete: 62,
	MessageDelete: 72,
	MessageBulkDelete: 73,
	MessagePin: 74,
	MessageUnpin: 75,
	IntegrationCreate: 80,
	IntegrationUpdate: 81,
	IntegrationDelete: 82,
	StageInstanceCreate: 83,
	StageInstanceUpdate: 84,
	StageInstanceDelete: 85,
	StickerCreate: 90,
	StickerUpdate: 91,
	StickerDelete: 92,
	GuildScheduledEventCreate: 100,
	GuildScheduledEventUpdate: 101,
	GuildScheduledEventDelete: 102,
	ThreadCreate: 110,
	ThreadUpdate: 111,
	ThreadDelete: 112,
	ApplicationCommandPermissionUpdate: 121,
	AutoModerationRuleCreate: 140,
	AutoModerationRuleUpdate: 141,
	AutoModerationRuleDelete: 142,
	AutoModerationBlockMessage: 143,
	AutoModerationFlagToChannel: 144,
	AutoModerationUserCommunicationDisabled: 145,
	CreatorMonetizationRequestCreated: 150,
	CreatorMonetizationTermsAccepted: 151,
}

