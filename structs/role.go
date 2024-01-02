package structs

type Role struct {
	// int representation of hexadecimal color code.
	Color int `json:"color"`

	// Role flags combined as a bitfield.
	Flags RoleFlags `json:"flags"`

	// If this role is pinned in the user listing.
	Hoist bool `json:"hoist"`

	// Role icon hash.
	Icon *string `json:"icon,omitempty"`

	// Role id.
	Id Snowflake `json:"id"`

	// Whether this role is managed by an integration.
	Managed bool `json:"managed"`

	// Whether this role is mentionable.
	Mentionable bool `json:"mentionable"`

	// Role name.
	Name string `json:"name"`

	// Permission bit set.
	Permissions string `json:"permissions"`

	// Position of this role.
	Position int `json:"position"`

	// The tags this role has.
	Tags *RoleTags `json:"tags,omitempty"`

	// Role unicode emoji.
	UnicodeEmoji *string `json:"unicode_emoji,omitempty"`
}

type RoleFlags int8

var ROLE_FLAGS = struct {
	// Role can be selected by members in an onboarding prompt.
	InPrompt RoleFlags
} {
	InPrompt: 1 << 0,
}

type RoleTags struct {
	// Whether this role is available for purchase.
	AvailableForPurchase *bool `json:"available_for_purchase,omit_empty"`

	// The id of the bot this role belongs to.
	BotId *Snowflake `json:"bot_id,omit_empty"`

	// Whether this role is a guild's linked role.
	GuildConnections *bool `json:"guild_connections,omit_empty"`

	// The id of the integration this role belongs to.
	IntegrationId *Snowflake `json:"integration_id,omit_empty"`

	// Whether this is the guild's Booster role.
	PremiumSubscriber *bool `json:"premium_subscriber,omit_empty"`

	// The id of this role's subscription sku and listing.
	SubscriptionListingId *Snowflake `json:"subscription_listing_id,omit_empty"`
}

