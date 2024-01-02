package structs

type Message struct {
	// Sent with Rich Presence-related chat embeds.
	Activity *MessageActivity `json:"activity,omitempty"`

	// If the message is an Interaction or application-owned webhook, this is the id of the application.
	ApplicationId *Snowflake `json:"application_id,omitempty"`

	// Sent with Rich Presence-related chat embeds.
	Application *Application `json:"application,omitempty"`

	// Any attached files.
	Attachments []MessageAttachment `json:"attachments"`

	// The author of this message (not guaranteed to be a valid user, see below).
	Author User `json:"author"`

	// Id of the channel the message was sent in.
	ChannelId Snowflake `json:"channel_id"`

	// Sent if the message contains components like buttons, action rows, or other interactive components.
	Components []MessageComponent `json:"components"`

	// Contents of the message.
	Content string `json:"content"`

	// When this message was edited (or null if never).
	EditedTimestamp *string `json:"edited_timestamp"`

	// Any embedded content.
	Embeds []MessageEmbed `json:"embeds**"`

	// Message flags combined as a bitfield.
	Flags *int `json:"flags,omitempty"`

	// Id of the message.
	Id Snowflake `json:"id"`

	// Sent if the message is a response to an Interaction.
	Interaction *MessageInteraction `json:"interaction,omitempty"`

	// Channels specifically mentioned in this message.
	MentionChannels *ChannelMention `json:"mention_channels,omitempty"`

	// Whether this message mentions everyone.
	MentionEveryone bool `json:"mention_everyone"`

	// Roles specifically mentioned in this message.
	MentionRoles []string `json:"mention_roles"`

	// Users specifically mentioned in the message.
	Mentions []User `json:"mentions"`

	// Data showing the source of a crosspost, channel follow add, pin, or reply message.
	MessageReference *MessageReference `json:"message_reference,omitempty"`

	// Used for validating a message was sent.
	Nonce string `json:"nonce,omitempty"`

	// Whether this message is pinned.
	Pinned bool `json:"pinned"`

	// A generally increasing int (there may be gaps or duplicates) that represents the approximate position of the message in a thread, it can be used to estimate the relative position of the message in a thread in company with total_message_sent on parent thread
	Position int `json:"position"`

	// Reactions to the message.
	Reactions *[]Reaction `json:"reactions,omitempty"`

	// The message associated with the message_reference.
	ReferencedMessage *Message `json:"referenced_message,omitempty"`

	// Data for users, members, channels, and roles in the message's auto-populated select menus.
	Resolved *ResolvedData `json:"resolved,omitempty"`

	// Data of the role subscription purchase or renewal that prompted this ROLE_SUBSCRIPTION_PURCHASE message.
	RoleSubscriptionData *RoleSubscriptionData `json:"role_subscription_data,omitempty"`

	// Sent if the message contains stickers.
	StickerItems *[]StickerItem `json:"sticker_items,omitempty"`

	// The thread that was started from this message, includes thread member object.
	Thread *Channel `json:"thread,omitempty"`

	// When this message was sent.
	Timestamp string `json:"timestamp"`

	// Whether this was a TTS message.
	Tts bool `json:"tts"`

	// Type of message.
	Type int `json:"type"`

	// If the message is generated by a webhook, this is the webhook's id.
	WebhookId *Snowflake `json:"webhook_id,omitempty"`
}

type MessageActivity struct {
	// Party id from a Rich Presence event.
	PartyId *string `json:"party_id,omitempty"`

	// Type of message activity.
	Type int8 `json:"type"`
}

type MessageAttachment struct {
	// The attachment's media type.
	ContentType *string `json:"content_type,omitempty"`

	// Description for the file (max 1024 characters).
	Description *string `json:"description,omitempty"`

	// The duration of the audio file (currently for voice messages).
	DurationSecs *float64 `json:"duration_secs,omitempty"`

	// Whether this attachment is ephemeral.
	Ephemeral *bool `json:"ephemeral,omitempty"`

	// Name of file attached.
	Filename string `json:"filename"`

	// Attachment flags combined as a bitfield.
	Flags *int `json:"flags,omitempty"`

	// Height of file (if image).
	Height *int `json:"height,omitempty"`

	// Attachment id.
	Id Snowflake `json:"id"`

	// A proxied url of file.
	ProxyUrl string `json:"proxy_url"`

	// Size of file in bytes.
	Size int `json:"size"`

	// Source url of file.
	Url string `json:"url"`

	// Base64 encoded bytearray representing a sampled waveform (currently for voice messages).
	Waveform *string `json:"waveform,omitempty"`

	// Width of file (if image).
	Width *int `json:"width,omitempty"`
}

type MessageComponent struct {

}

type MessageEmbed struct {
	// Author information.
	Author *EmbedAuthor `json:"author,omitempty"`

	// Color code of the embed.
	Color *int `json:"color,omitempty"`

	// Description of embed.
	Description *string `json:"description,omitempty"`

	// Fields information, max of 25.
	Fields *[]EmbedField `json:"fields,omitempty"`

	// Footer information.
	Footer *EmbedFooter `json:"footer,omitempty"`

	// Image information.
	Image *EmbedImage `json:"image,omitempty"`

	// Provider information.
	Provider *EmbedProvider `json:"provider,omitempty"`

	// Thumbnail information.
	Thumbnail *EmbedThumbnail `json:"thumbnail,omitempty"`

	// Timestamp of embed content.
	Timestamp *string `json:"timestamp,omitempty"`

	// Title of embed.
	Title *string `json:"title,omitempty"`

	// Type of embed (always "rich" for webhook embeds).
	Type *string `json:"type,omitempty"`

	// Url of embed.
	Url *string `json:"url,omitempty"`

	// Video information.
	Video *EmbedVideo `json:"video,omitempty"`
}

type EmbedAuthor struct {
	IconUrl *string `json:"icon_url,omitempty"`
	Name string `json:"name"`
	ProxyIconUrl *string `json:"proxy_icon_url,omitempty"`
	Url *string `json:"url,omitempty"`
}

type EmbedField struct {
	Inline bool `json:"inline"`
	Name string `json:"name"`
	Value string `json:"string"`
}

type EmbedFooter struct {
	IconUrl *string `json:"icon_url,omitempty"`
	ProxyIconUrl *string `json:"proxy_icon_url,omitempty"`
	Text string `json:"text"`
}

type EmbedImage struct {
	Height *int `json:"height,omitempty"`
	ProxyUrl *string `json:"proxy_url,omitempty"`
	Url string `json:"url"`
	Width *int `json:"width,omitempty"`
}

type EmbedProvider struct {
	Name *string `json:"name,omitempty"`
	Url *string `json:"url,omitempty"`
}

type EmbedThumbnail EmbedImage

type EmbedVideo EmbedImage

type MessageInteraction struct {

}

type ChannelMention struct {

}

type MessageReference struct {

}

type Reaction struct {

}

type ResolvedData struct {

}

type StickerItem struct {

}

type RoleSubscriptionData struct {

}
