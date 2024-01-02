package structs

type User struct {
	// The user's banner color encoded as an integer representation of hexadecimal color code.
	AccentColor *int `json:"accent_color,omitempty"`

	// The user's avatar hash.
	Avatar *string `json:"avatar,omitempty"`

	// The user's avatar decoration hash.
	AvatarDecoration *string `json:"avatar_decoration,omitempty"`

	// The user's banner hash.
	Banner *string `json:"banner,omitempty"`

	// Whether the user belongs to an OAuth2 application.
	Bot *bool `json:"bot,omitempty"`

	// The user's Discord-tag.
	Discriminator string `json:"discriminator"`

	// The user's email.
	Email *string `json:"email,omitempty"`

	// The flags on a user's account.
	Flags *UserFlags `json:"flags,omitempty"`

	// The user's display name, if it is set. For bots, this is the application name.
	GlobalName *string `json:"global_name"`

	// The user's id.
	Id Snowflake `json:"id"`

	// The user's chosen language option.
	Locale *string `json:"locale,omitempty"`

	// Whether the user has two factor enabled on their account.
	MFAEnabled *bool `json:"mfa_enabled,omitempty"`

	// The type of Nitro subscription on a user's account.
	PremiumType *PremiumType `json:"premium_type,omitempty"`

	// The public flags on a user's account.
	PublicFlags *UserFlags `json:"public_flags,omitempty"`

	// Whether the user is an Official Discord System user (part of the urgent message system).
	System *bool `json:"system,omitempty"`

	// The user's username, not unique across the platform.
	Username string `json:"username"`

	// Whether the email on this account has been verified.
	Verified *bool `json:"verified,omitempty"`
}

