package structs

type Channel struct {
	// application id of the group DM creator if it is bot-created.
	ApplicationId *Snowflake `json:"application_id,omitempty"`

	// the IDs of the set of tags that have been applied to a thread in a GUILD_FORUM or a GUILD_MEDIA channel.
	AppliedTags *[]Snowflake `json:"applied_tags,omitempty"`

	// the set of tags that can be used in a GUILD_FORUM or a GUILD_MEDIA channel.
	AvailableTags *[]ForumTag `json:"available_tags,omitempty"`

	// the bitrate (in bits) of the voice channel.
	Bitrate *int `json:"bitrate,omitempty"`

	// default duration, copied onto newly created threads, in minutes, threads will stop showing in the channel list after the specified period of inactivity, can be set to: 60, 1440, 4320, 10080.
	DefaultAutoArchiveDuration *int `json:"default_auto_archive_duration,omitempty"`

	// the default forum layout view used to display posts in GUILD_FORUM channels. Defaults to 0, which indicates a layout view has not been set by a channel admin.
	DefaultForumLayout *int `json:"default_forum_layout,omitempty"`

	// the emoji to show in the add reaction button on a thread in a GUILD_FORUM or a GUILD_MEDIA channel.
	DefaultReactionEmoji *DefaultReaction `json:"default_reaction_emoji,omitempty"`

	// the default sort order type used to order posts in GUILD_FORUM and GUILD_MEDIA channels. Defaults to null, which indicates a preferred sort order hasn't been set by a channel admin.
	DefaultSortOrder *int `json:"default_sort_order,omitempty"`

	// the initial rate_limit_per_user to set on newly created threads in a channel. this field is copied to the thread at creation time and does not live update..
	DefaultThreadRateLimitPerUser *int `json:"default_thread_rate_limit_per_user,omitempty"`

	// channel flags combined as a bitfield.
	Flags *int `json:"flags,omitempty"`

	// the id of the guild (may be missing for some channel objects received over gateway guild dispatches).
	GuildId *Snowflake `json:"guild_id,omitempty"`

	// icon hash of the group DM.
	Icon *string `json:"icon,omitempty"`

	// the id of this channel.
	Id Snowflake `json:"id"`

	// the id of the last message sent in this channel (or thread for GUILD_FORUM or GUILD_MEDIA channels) (may not point to an existing or valid message or thread).
	LastMessageId *Snowflake `json:"last_message_id,omitempty"`

	// when the last pinned message was pinned. This may be null in events such as GUILD_CREATE when a message is not pinned..
	LastPinTimestamp *string `json:"last_pin_timestamp,omitempty"`

	// for group DM channels: whether the channel is managed by an application via the gdm.join OAuth2 scope.
	Managed *bool `json:"managed,omitempty"`

	// an approximate count of users in a thread, stops counting at 50.
	MemberCount *int `json:"member_count,omitempty"`

	// thread member object for the current user, if they have joined the thread, only included on certain API endpoints.
	Member *ThreadMember `json:"member,omitempty"`

	// number of messages (not including the initial message or deleted messages) in a thread..
	MessageCount *int `json:"message_count,omitempty"`

	// the name of the channel (1-100 characters).
	Name *string `json:"name,omitempty"`

	// whether the channel is nsfw.
	Nsfw *bool `json:"nsfw,omitempty"`

	// id of the creator of the group DM or thread.
	OwnerId *Snowflake `json:"owner_id,omitempty"`

	// for guild channels: id of the parent category for a channel (each parent category can contain up to 50 channels), for threads: id of the text channel this thread was created.
	ParentId *Snowflake `json:"parent_id,omitempty"`

	// explicit permission overwrites for members and roles.
	PermissionOverwrites *[]ChannelOverwrite `json:"permission_overwrites,omitempty"`

	// computed permissions for the invoking user in the channel, including overwrites, only included when part of the resolved data received on a slash command interaction. This does not include implicit permissions, which may need to be checked separately.
	Permissions *string `json:"permissions,omitempty"`

	// sorting position of the channel.
	Position *int `json:"position,omitempty"`

	// amount of seconds a user has to wait before sending another message (0-21600); bots, as well as users with the permission manage_messages or manage_channel, are unaffected.
	RateLimitPerUser* *int `json:"rate_limit_per_user,omitempty*"`

	// the recipients of the DM.
	Recipients *[]User `json:"recipients,omitempty"`

	// voice region id for the voice channel, automatic when set to null.
	RtcRegion *string `json:"rtc_region,omitempty"`

	// thread-specific fields not needed by other channels.
	ThreadMetadata *ThreadMetadata `json:"thread_metadata,omitempty"`

	// the channel topic (0-4096 characters for GUILD_FORUM and GUILD_MEDIA channels, 0-1024 characters for all others).
	Topic *string `json:"topic,omitempty"`

	// number of messages ever sent in a thread, it's similar to message_count on message creation, but will not decrement the number when a message is deleted.
	TotalMessageSent *int `json:"total_message_sent,omitempty"`

	// the type of channel.
	Type ChannelType `json:"type"`

	// the user limit of the voice channel.
	UserLimit *int `json:"user_limit,omitempty"`

	// the camera video quality mode of the voice channel, 1 when not present.
	VideoQualityMode *int `json:"video_quality_mode,omitempty"`
}

type ChannelOverwrite struct {
	// Permission bit set.
	Allow string `json:"allow"`

	// Permission bit set.
	Deny string `json:"deny"`

	// Role or user id.
	Id Snowflake `json:"id"`

	// Either 0 (role) or 1 (member).
	Type int `json:"type"`
}

type DefaultReaction struct {
	// The id of a guild's custom emoji.
	EmojiId *Snowflake `json:"emoji_id,omitempty"`

	// The unicode character of the emoji.
	EmojiName *string `json:"emoji_name,omitempty"`
}

type ForumTag struct {
	// The id of a guild's custom emoji.
	EmojiId *Snowflake `json:"emoji_id,omitempty"`

	// The unicode character of the emoji.
	EmojiName *string `json:"emoji_name,omitempty"`

	// The id of the tag.
	Id Snowflake `json:"id"`

	// Whether this tag can only be added to or removed from threads by a member with the MANAGE_THREADS permission.
	Moderated bool `json:"moderated"`

	// The name of the tag (0-20 characters).
	Name string `json:"name"`
}

type ThreadMember struct {
	// Any user-thread settings, currently only used for notifications.
	Flags int `json:"flags"`

	// ID of the thread.
	Id *Snowflake `json:"id,omitempty"`

	// Time the user last joined the thread.
	JoinTimestamp string `json:"join_timestamp"`

	// Additional information about the user.
	Member *GuildMember `json:"member,omitempty"`

	// ID of the user.
	UserId *Snowflake `json:"user_id,omitempty"`
}

type ThreadMetadata struct {
	// timestamp when the thread's archive status was last changed, used for calculating recent activity.
	ArchiveTimestamp string `json:"archive_timestamp"`

	// whether the thread is archived.
	Archived bool `json:"archived"`

	// the thread will stop showing in the channel list after auto_archive_duration minutes of inactivity, can be set to: 60, 1440, 4320, 10080.
	AutoArchiveDuration int `json:"auto_archive_duration"`

	// timestamp when the thread was created; only populated for threads created after 2022-01-09.
	CreateTimestamp *string `json:"create_timestamp,omitempty"`

	// whether non-moderators can add other non-moderators to a thread; only available on private threads.
	Invitable *bool `json:"invitable,omitempty"`

	// whether the thread is locked; when a thread is locked, only users with MANAGE_THREADS can unarchive it.
	Locked bool `json:"locked"`
}

