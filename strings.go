package helper

import "strings"

func Oneline(multilineJson string) string {
	return strings.Replace(multilineJson, "\n", "", -1)
}
