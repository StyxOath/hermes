package structs

type UserFlags int32

var USER_FLAGS = struct {
	// Discord Employee.
	Staff UserFlags

	// Partnered Server Owner.
	Partner UserFlags

	// HypeSquad Events Member.
	HypeSquad UserFlags

	// Bug Hunter Level 1.
	BugHunterLevel1 UserFlags

	// House Bravery Member.
	HypeSquadOnlineHouse1 UserFlags

	// House Brilliance Member.
	HypeSquadOnlineHouse2 UserFlags

	// House Balance Member.
	HypeSquadOnlineHouse3 UserFlags

	// Early Nitro Supporter.
	PremiumEarlySupporter UserFlags

	// User is a team.
	TeamPseudoUser UserFlags

	// Bug Hunter Level 2.
	BugHunterLevel2 UserFlags

	// Verified Bot.
	VerifiedBot UserFlags

	// Early Verified Bot Developer.
	VerifiedDeveloper UserFlags

	// Moderator Programs Alumni.
	CertifiedModerator UserFlags

	// Bot uses only HTTP interactions and is shown in the online member list.
	BotHttpInteractions UserFlags

	// User is an Active Developer.
	ActiveDeveloper UserFlags
} {
	Staff: 1 << 0,
	Partner: 1 << 1,
	HypeSquad: 1 << 2,
	BugHunterLevel1: 1 << 3,
	HypeSquadOnlineHouse1: 1 << 6,
	HypeSquadOnlineHouse2: 1 << 7,
	HypeSquadOnlineHouse3: 1 << 8,
	PremiumEarlySupporter: 1 << 9,
	TeamPseudoUser: 1 << 10,
	BugHunterLevel2: 1 << 14,
	VerifiedBot: 1 << 16,
	VerifiedDeveloper: 1 << 17,
	CertifiedModerator: 1 << 18,
	BotHttpInteractions: 1 << 19,
	ActiveDeveloper: 1 << 22,
}

