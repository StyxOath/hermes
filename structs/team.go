package structs

type Team struct {
	// Hash of the image of the team's icon.
	Icon *string `json:"icon"`

	// Unique ID of the team.
	Id Snowflake `json:"id"`

	// Members of the team.
	Members []TeamMember `json:"members"`

	// Name of the team.
	Name string `json:"name"`

	// User ID of the current team owner.
	Owner_user_id Snowflake `json:"owner_user_id"`
}

type TeamMember struct {
	// User's membership state on the team.
	MembershipState int8 `json:"membership_state"`

	// ID of the parent team of which they are a member.
	TeamId Snowflake `json:"team_id"`

	// Avatar, discriminator, ID, and username of the user.
	User struct {
		Avatar *string `json:"avatar,omitempty"`
		Discriminator string `json:"discriminator"`
		Id Snowflake `json:"id"`
		Username string `json:"username"`
	} `json:"user"`

	// Role of the team member.
	Role string `json:"role"`
}

