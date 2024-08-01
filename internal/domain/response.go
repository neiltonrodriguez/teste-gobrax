package domain

import (
	"encoding/json"
)

type Response struct {
	Meta  Meta        `json:"meta,omitempty"`
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"errors,omitempty"`
}

type Meta struct {
	Count      int                 `json:"count"`
	Pagination *paginationResponse `json:"pagination,omitempty"`
}

func ToJson(v interface{}) string {
	jsonToReturn, e := json.Marshal(v)
	if e != nil {
		return ""
	} else {
		return string(jsonToReturn)
	}
}

func FromJson(v interface{}, jsonString string) error {
	return json.Unmarshal([]byte(jsonString), &v)
}
