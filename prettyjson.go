package tools

import (
	"bytes"
	"encoding/json"
)

const (
	empty = ""
	tab = "\t"
)

func prettyString(data string) (string, error) {
	buffer := new(bytes.Buffer)
	err := json.Indent(buffer, []byte(data), empty, tab)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func prettyStruct(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)
	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}

func PrettyJSON(data interface{}) (string, error) {
	switch data.(type) {
	case string:
		return prettyString(data.(string))
	default:
		return prettyStruct(data)
	}
}
