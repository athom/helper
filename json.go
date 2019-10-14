package helper

import (
	"encoding/json"
	"fmt"
)

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

func UglyJson(str string) string {
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(str), &data); err != nil {
		panic(err)
	}

	json, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return "format error"
	}
	s := fmt.Sprintf("%v", string(json))
	return Oneline(s)
}
