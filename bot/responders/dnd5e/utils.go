package dnd5e

import (
	"strconv"
)

func toIntColor(num string) (numColor int, err error) {
	output, err := strconv.ParseUint(num, 0, 64)
	return int(output), err
}
