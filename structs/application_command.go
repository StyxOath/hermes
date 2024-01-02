package structs

type GuildApplicationCommandPermissions struct {
	// ID of the application the command belongs to.
	ApplicationId Snowflake `json:"application_id"`

	// ID of the guild.
	GuildId Snowflake `json:"guild_id"`

	// ID of the command or the application ID.
	Id Snowflake `json:"id"`

	// Permissions for the command in the guild, max of 100.
	Permissions []ApplicationCommandPermissions `json:"permissions"`
}

type ApplicationCommandPermissions struct {
	// ID of the role, user, or channel. It can also be a permission constant.
	Id Snowflake `json:"id"`

	// True to allow, false, to disallow.
	Permission bool `json:"permission"`

	// Role (1), user (2), or channel (3).
	Type int8 `json:"type"`
}

