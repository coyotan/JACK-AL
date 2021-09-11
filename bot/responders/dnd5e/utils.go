package dnd5e

import (
	"errors"
	"strconv"
)

func toIntColor(num string) (numColor int, err error) {
	output, err := strconv.ParseInt(num, 0, 64)
	if output < 1 {
		return int(output), err
	} else {
		return 0, errors.New("output was less than 1")
	}
}
