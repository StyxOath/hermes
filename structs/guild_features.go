package structs

type GuildFeature string

var GUILD_FEATURES = struct {
	// Guild has access to set an animated guild banner image.
	AnimatedBanner GuildFeature

	// Guild has access to set an animated guild icon.
	AnimatedIcon GuildFeature

	// Guild is using the old permissions configuration behavior.
	ApplicationCommandPermissionsV2 GuildFeature

	// Guild has set up auto moderation rules.
	AutoModeration GuildFeature

	// Guild has access to set a guild banner image.
	Banner GuildFeature

	// Guild can enable welcome screen, Membership Screening, stage channels and discovery, and receives community updates.
	Community GuildFeature

	// Guild has enabled monetization.
	CreatorMonetizableProvisional GuildFeature

	// Guild has enabled the role subscription promo page.
	CreatorStorePage GuildFeature

	// Guild has been set as a support server on the App Directory.
	DeveloperSupportServer GuildFeature

	// Guild is able to be discovered in the directory.
	Discoverable GuildFeature

	// Guild is able to be featured in the directory.
	Featurable GuildFeature

	// Description.
	Feature GuildFeature

	// Guild has access to set an invite splash background.
	InviteSplash GuildFeature

	// Guild has paused invites, preventing new users from joining.
	InvitesDisabled GuildFeature

	// Guild has enabled Membership Screening.
	MemberVerificationGateEnabled GuildFeature

	// Guild has increased custom sticker slots.
	MoreStickers GuildFeature

	// Guild has access to create announcement channels.
	News GuildFeature

	// Guild is partnered.
	Partnered GuildFeature

	// Guild can be previewed before joining via Membership Screening or the directory.
	PreviewEnabled GuildFeature

	// Guild has disabled alerts for join raids in the configured safety alerts channel.
	RaidAlertsDisabled GuildFeature

	// Guild is able to set role icons.
	RoleIcons GuildFeature

	// Guild has role subscriptions that can be purchased.
	RoleSubscriptionsAvailableForPurchase GuildFeature

	// Guild has enabled role subscriptions.
	RoleSubscriptionsEnabled GuildFeature

	// Guild has enabled ticketed events.
	TicketedEventsEnabled GuildFeature

	// Guild has access to set a vanity URL.
	VanityUrl GuildFeature

	// Guild is verified.
	Verified GuildFeature

	// Guild has access to set 384kbps bitrate in voice (previously VIP voice servers).
	VipRegions GuildFeature

	// Guild has enabled the welcome screen.
	WelcomeScreenEnabled GuildFeature
} {
	AnimatedBanner: "ANIMATED_BANNER",
	AnimatedIcon: "ANIMATED_ICON",
	ApplicationCommandPermissionsV2: "APPLICATION_COMMAND_PERMISSIONS_V2",
	AutoModeration: "AUTO_MODERATION",
	Banner: "BANNER",
	Community: "COMMUNITY",
	CreatorMonetizableProvisional: "CREATOR_MONETIZABLE_PROVISIONAL",
	CreatorStorePage: "CREATOR_STORE_PAGE",
	DeveloperSupportServer: "DEVELOPER_SUPPORT_SERVER",
	Discoverable: "DISCOVERABLE",
	Featurable: "FEATURABLE",
	InviteSplash: "INVITE_SPLASH",
	InvitesDisabled: "INVITES_DISABLED",
	MemberVerificationGateEnabled: "MEMBER_VERIFICATION_GATE_ENABLED",
	MoreStickers: "MORE_STICKERS",
	News: "NEWS",
	Partnered: "PARTNERED",
	PreviewEnabled: "PREVIEW_ENABLED",
	RaidAlertsDisabled: "RAID_ALERTS_DISABLED",
	RoleIcons: "ROLE_ICONS",
	RoleSubscriptionsAvailableForPurchase: "ROLE_SUBSCRIPTIONS_AVAILABLE_FOR_PURCHASE",
	RoleSubscriptionsEnabled: "ROLE_SUBSCRIPTIONS_ENABLED",
	TicketedEventsEnabled: "TICKETED_EVENTS_ENABLED",
	VanityUrl: "VANITY_URL",
	Verified: "VERIFIED",
	VipRegions: "VIP_REGIONS",
	WelcomeScreenEnabled: "WELCOME_SCREEN_ENABLED",
}

