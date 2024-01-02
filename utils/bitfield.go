package utils

import "golang.org/x/exp/constraints"

func BitfieldHasAll[T constraints.Integer](bitfield T, expected T) bool {
	return bitfield & expected == expected
}

func BitfieldHasAny[T constraints.Integer](bitfield T, expected T) bool {
	return bitfield & expected > 0
}
