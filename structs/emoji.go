package structs

type Emoji struct {
	// Whether this emoji is animated.
	Animated *bool `json:"animated,omitempty"`

	// Whether this emoji can be used, may be false due to loss of Server Boosts.
	Available *bool `json:"available,omitempty"`

	// Emoji id.
	Id *Snowflake `json:"id"`

	// Whether this emoji is managed.
	Managed *bool `json:"managed,omitempty"`

	// Emoji name.
	Name *string `json:"name"`

	// Whether this emoji must be wrapped in colons.
	RequireColons *bool `json:"require_colons,omitempty"`

	// Roles allowed to use this emoji.
	Roles *[]string `json:"roles,omitempty"`

	// User that created this emoji.
	User *User `json:"user,omitempty"`
}

