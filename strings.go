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
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		panic(err)
	}

	json, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "format error"
	}
	return fmt.Sprintf("%v", string(json))
}
