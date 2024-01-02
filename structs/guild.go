package structs

type Guild struct {
	// Id of afk channel.
	AfkChannelId *Snowflake `json:"afk_channel_id,omitempty"`

	// Afk timeout in seconds.
	AfkTimeout int `json:"afk_timeout"`

	// Application id of the guild creator if it is bot-created.
	ApplicationId *Snowflake `json:"application_id,omitempty"`

	// Approximate number of members in this guild, returned from the GET /guilds/<id> and /users/@me/guilds endpoints when with_counts is true.
	ApproximateMemberCount *int `json:"approximate_member_count,omitempty"`

	// Approximate number of non-offline members in this guild, returned from the GET /guilds/<id> and /users/@me/guilds endpoints when with_counts is true.
	ApproximatePresenceCount *int `json:"approximate_presence_count,omitempty"`

	// Banner hash.
	Banner *string `json:"banner,omitempty"`

	// Default message notifications level.
	DefaultMessageNotifications int `json:"default_message_notifications"`

	// The description of a guild.
	Description *string `json:"description,omitempty"`

	// Discovery splash hash; only present for guilds with the "DISCOVERABLE" feature.
	DiscoverySplash *string `json:"discovery_splash,omitempty"`

	// Custom guild emojis.
	Emojis []Emoji `json:"emojis"`

	// Explicit content filter level.
	ExplicitContentFilter int `json:"explicit_content_filter"`

	// Enabled guild features.
	Features []GuildFeature `json:"features"`

	// Icon hash.
	Icon *string `json:"icon,omitempty"`

	// Icon hash, returned when in the template object.
	IconHash *string `json:"icon_hash,omitempty"`

	// Guild id.
	Id Snowflake `json:"id"`

	// The maximum number of members for the guild.
	MaxMembers *int `json:"max_members,omitempty"`

	// The maximum number of presences for the guild (null is always returned, apart from the largest of guilds).
	MaxPresences *int `json:"max_presences,omitempty"`

	// The maximum amount of users in a stage video channel.
	MaxStageVideoChannelUsers *int `json:"max_stage_video_channel_users,omitempty"`

	// The maximum amount of users in a video channel.
	MaxVideoChannelUsers *int `json:"max_video_channel_users,omitempty"`

	// Required MFA level for the guild.
	MFALevel int `json:"mfa_level"`

	// Guild name (2-100 characters, excluding trailing and leading whitespace).
	Name string `json:"name"`

	// Guild NSFW level.
	NSFWLevel int `json:"nsfw_level"`

	// Id of owner.
	OwnerId Snowflake `json:"owner_id"`

	// True if the user is the owner of the guild.
	Owner *bool `json:"owner,omitempty"`

	// Total permissions for the user in the guild (excludes overwrites and implicit permissions).
	Permissions *string `json:"permissions,omitempty"`

	// The preferred locale of a Community guild; used in server discovery and notices from Discord, and sent in interactions; defaults to "en-US".
	PreferredLocale string `json:"preferred_locale"`

	// Whether the guild has the boost progress bar enabled.
	PremiumProgressBarEnabled bool `json:"premium_progress_bar_enabled"`

	// The number of boosts this guild currently has.
	PremiumSubscriptionCount *int `json:"premium_subscription_count,omitempty"`

	// Premium tier (Server Boost level).
	PremiumTier int `json:"premium_tier"`

	// The id of the channel where admins and moderators of Community guilds receive notices from Discord.
	PublicUpdatesChannelId *Snowflake `json:"public_updates_channel_id,omitempty"`

	// Voice region id for the guild (deprecated).
	Region *string `json:"region,omitempty"`

	// Roles in the guild.
	Roles []Role `json:"roles"`

	// The id of the channel where Community guilds can display rules and/or guidelines.
	RulesChannelId *Snowflake `json:"rules_channel_id,omitempty"`

	// The id of the channel where admins and moderators of Community guilds receive safety alerts from Discord.
	SafetyAlertsChannelId *Snowflake `json:"safety_alerts_channel_id,omitempty"`

	// Splash hash.
	Splash *string `json:"splash,omitempty"`

	// Custom guild stickers.
	Stickers *[]Sticker `json:"stickers,omitempty"`

	// System channel flags.
	SystemChannel_Flags int `json:"system_channel_flags"`

	// The id of the channel where guild notices such as welcome messages and boost events are posted.
	SystemChannelId *Snowflake `json:"system_channel_id,omitempty"`

	// The vanity url code for the guild.
	VanityUrlCode *string `json:"vanity_url_code,omitempty"`

	// Verification level required for the guild.
	VerificationLevel int `json:"verification_level"`

	// The welcome screen of a Community guild, shown to new members, returned in an Invite's guild object.
	WelcomeScreen *WelcomeScreen `json:"welcome_screen,omitempty"`

	// The channel id that the widget will generate an invite to, or null if set to no invite.
	WidgetChannelId *Snowflake `json:"widget_channel_id,omitempty"`

	// True if the server widget is enabled.
	WidgetEnabled *bool `json:"widget_enabled,omitempty"`
}

