package structs

type PremiumType int8

var PREMIUM_TYPES = struct {
	// None
	None PremiumType

	// Nitro Classic
	NitroClassic PremiumType

	// Nitro
	Nitro PremiumType
	
	// Nitro Basic
	NitroBasic PremiumType
} {
	None: 0,
	NitroClassic: 1,
	Nitro: 2,
	NitroBasic: 3,
}

