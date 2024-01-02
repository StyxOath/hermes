package structs

type WelcomeScreen struct {
	// The server description shown in the welcome screen.
	Description *string `json:"description,omitempty"`

	// The channels shown in the welcome screen, up to 5.
	WelcomeChannels []WelcomeScreenChannel `json:"welcome_channels"`
}

type WelcomeScreenChannel struct {
	// The channel's id.
	ChannelId Snowflake `json:"channel_id"`

	// The description shown for the channel.
	Description string `json:"description"`

	// The emoji id, if the emoji is custom.
	EmojiId *Snowflake `json:"emoji_id,omitempty"`

	// The emoji name if custom, the unicode character if standard, or null if no emoji is set.
	EmojiName *string `json:"emoji_name,omitempty"`
}

