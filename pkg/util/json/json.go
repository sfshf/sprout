package json

import (
	jsoniter "github.com/json-iterator/go"
)

// Temporary:
var (
	Marshal       = jsoniter.Marshal
	Unmarshal     = jsoniter.Unmarshal
	MarshalIndent = jsoniter.MarshalIndent
	NewDecoder    = jsoniter.NewDecoder
	NewEncoder    = jsoniter.NewEncoder
)

func Marshal2String(v interface{}) string {
	s, err := jsoniter.MarshalToString(v)
	if err != nil {
		return "ERROR: " + err.Error()
	}
	return s
}

func MarshalIndent2String(v interface{}) string {
	bs, err := jsoniter.MarshalIndent(v, "", "    ")
	if err != nil {
		return "ERROR: " + err.Error()
	}
	return string(bs)
}
