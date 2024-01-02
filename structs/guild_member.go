package structs

type GuildMember struct {
	// The member's guild avatar hash.
	Avatar *string `json:"avatar,omitempty"`

	// When the user's timeout will expire and the user will be able to communicate in the guild again, null or a time in the past if the user is not timed out.
	CommunicationDisabledUntil *string `json:"communication_disabled_until,omitempty"`

	// Whether the user is deafened in voice channels.
	Deaf bool `json:"deaf"`

	// Guild member flags represented as a bit set, defaults to 0.
	Flags int `json:"flags"`

	// When the user joined the guild.
	JoinedAt string `json:"joined_at"`

	// Whether the user is muted in voice channels.
	Mute bool `json:"mute"`

	// This user's guild nickname.
	Nick *string `json:"nick,omitempty"`

	// Whether the user has not yet passed the guild's Membership Screening requirements.
	Pending *bool `json:"pending,omitempty"`

	// Total permissions of the member in the channel, including overwrites, returned when in the interaction object.
	Permissions *string `json:"permissions,omitempty"`

	// When the user started boosting the guild.
	PremiumSince *string `json:"premium_since,omitempty"`

	// Array of role object ids.
	Roles []Snowflake `json:"roles"`

	// The user this guild member represents.
	User *User `json:"user,omitempty"`
}

