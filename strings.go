package helper

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Oneline(multilineJson string) string {
	return strings.Replace(multilineJson, "\n", "", -1)
}

func PrettyJson(str string) string {
	json, err := json.MarshalIndent(str, "", "  ")
	if err != nil {
		return "format error"
	}
	return fmt.Sprintf("%v", string(json))
}
