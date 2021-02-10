package dnd5e

import (
	"strconv"
	"strings"
)

func hexaNumberToInteger(hexaString string) string {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)
	return numberStr
}

func toIntColor(num string) (numColor int, err error) {
	output, err := strconv.ParseInt(hexaNumberToInteger(num), 16, 64)
	return int(output), err
}
