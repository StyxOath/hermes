package structs

type Application struct {
	// Approximate count of guilds the app has been added to.
	ApproximateGuildCount int `json:"approximate_guild_count,omitempty"`
		
	// Partial user object for the bot user associated with the app.
	Bot *User `json:"bot,omitempty"`
		
	// When false, only the app owner can add the app to guilds.
	BotPublic bool `json:"bot_public"`
	
	// When true, the app's bot will only join upon completion of the full OAuth2 code grant flow.
	BotRequireCodeGrant bool `json:"bot_require_code_grant"`

	// App's default rich presence invite cover image hash.
	CoverImage *string `json:"cover_image,omitempty"`
		
	// Default custom authorization URL for the app, if enabled.
	CustomInstallURL *string `json:"custom_install_url,omitempty"`
	
	// Deprecated and will be removed in v11. An empty string.
	Description string `json:"description"`
	
	// App's public flags.
	Flags ApplicationFlags `json:"flags,omitempty"`

	// Guild associated with the app. For example, a developer support server.
	Guild *Guild `json:"guild,omitempty"`
	
	// Guild associated with the app. For example, a developer support server.
	GuildId Snowflake `json:"guild_id,omitempty"`
	
	// Hex encoded key for verification in interactions and the GameSDK's GetTicket.
	Icon *string `json:"icon,omitempty"`
	
	// ID of the app.
	Id Snowflake `json:"id"`

	// Partial object of the associated guild.
	InstallParams *InstallParams `json:"install_params,omitempty"`
	
	// Interactions endpoint URL for the app.
	InteractionsEndpointURL *string `json:"interactions_endpoint_url,omitempty"`
			
	// Partial object of the associated guild.
	Owner *User `json:"owner,omitempty"`
		
	// ID of the "Game SKU" that is created, if exists.
	PrimarySkuID Snowflake `json:"primary_sku_id,omitempty"`
	
	// URL of the app's Privacy Policy.
	PrivacyPolicyURL *string `json:"privacy_policy_url,omitempty"`

	// Role connection verification URL for the app.
	RedirectURIs []string `json:"redirect_uris,omitempty"`
	
	// List of tags describing the content and functionality of the app. Max of 5 tags.
	RoleConnectionsVerificationURL *string `json:"role_connections_verification_url,omitempty"`
	
	// List of RPC origin URLs, if RPC is enabled.
	RPCOrigins []string `json:"rpc_origins,omitempty"`

	// If this app is a game sold on Discord, this field will be the URL slug that links to the store page.
	Slug *string `json:"slug,omitempty"`
	
	// Deprecated and will be removed in v11. An empty string.
	Summary string `json:"summary"`

	// List of tags describing the content and functionality of the app. Max of 5 tags.
	Tags []string `json:"tags,omitempty"`
	
	// If the app belongs to a team, this will be a list of the members of that team.
	Team *Team `json:"team,omitempty"`

	// URL of the app's Terms of Service.
	TermsOfServiceURL *string `json:"terms_of_service_url,omitempty"`

	// When true, the app's bot will only join upon completion of the full OAuth2 code grant flow.
	VerifyKey string `json:"verify_key"`
}

