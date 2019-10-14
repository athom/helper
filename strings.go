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
	json, err := json.MarshalIndent([]byte(str), "", "  ")
	return fmt.Sprintf(json)
}
