package structs

type InstallParams struct {
	// Permissions to request for the bot role.
	Permissions Permission   `json:"permissions,string"`

	// Scopes to add the application to the server with.
	Scopes      []ApplicationScope `json:"scopes"`
}

