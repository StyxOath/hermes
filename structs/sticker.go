package structs

type Sticker struct {
	// Deprecated previously the sticker asset hash, now an empty string.
	Asset *string `json:"asset,omitempty"`

	// Whether this guild sticker can be used, may be false due to loss of Server Boosts.
	Available *bool `json:"available,omitempty"`

	// Description of the sticker.
	Description *string `json:"description,omitempty"`

	// Type of sticker format.
	FormatType StickerFormat `json:"format_type"`

	// Id of the guild that owns this sticker.
	GuildId *Snowflake `json:"guild_id,omitempty"`

	// Id of the sticker.
	Id Snowflake `json:"id"`

	// Name of the sticker.
	Name string `json:"name"`

	// For standard stickers, id of the pack the sticker is from.
	PackId *Snowflake `json:"pack_id,omitempty"`

	// The standard sticker's sort order within its pack.
	SortValue *int `json:"sort_value,omitempty"`

	// Autocomplete/suggestion tags for the sticker (max 200 characters).
	Tags string `json:"tags"`

	// Type of sticker.
	Type StickerType `json:"type"`

	// The user that uploaded the guild sticker.
	User *User `json:"user,omitempty"`
}

type StickerFormat int8

var STICKER_FORMATS = struct {
	PNG StickerFormat
	APNG StickerFormat
	LOTTIE StickerFormat
	GIF StickerFormat
} {
	PNG: 1,
	APNG: 2,
	LOTTIE: 3,
	GIF: 4,
}

type StickerType int8

var STICKER_TYPES = struct {
	// An official sticker in a pack.
	Standard StickerType

	// A sticker uploaded to a guild for the guild's members.
	Guild StickerType
} {
	Standard: 1,
	Guild: 2,
}
