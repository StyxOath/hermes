package structs

type Webhook struct {
	// The bot/OAuth2 application that created this webhook.
	ApplicationId *Snowflake `json:"application_id,omitempty"`

	// The default user avatar hash of the webhook.
	Avatar *string `json:"avatar,omitempty"`

	// The channel id this webhook is for, if any.
	ChannelId *Snowflake `json:"channel_id,omitempty"`

	// The guild id this webhook is for, if any.
	GuildId *Snowflake `json:"guild_id,omitempty"`

	// The id of the webhook.
	Id Snowflake `json:"id"`

	// The default name of the webhook.
	Name *string `json:"name"`

	// The channel that this webhook is following (returned for Channel Follower Webhooks).
	SourceChannel *Channel `json:"source_channel,omitempty"`

	// The guild of the channel that this webhook is following (returned for Channel Follower Webhooks).
	SourceGuild *Guild `json:"source_guild,omitempty"`

	// The secure token of the webhook (returned for Incoming Webhooks).
	Token *string `json:"token,omitempty"`

	// The type of the webhook.
	Type int `json:"type"`

	// The url used for executing the webhook (returned by the webhooks OAuth2 flow).
	Url *string `json:"url,omitempty"`

	// The user this webhook was created by (not returned when getting a webhook with its token).
	User *User `json:"user,omitempty"`
}
