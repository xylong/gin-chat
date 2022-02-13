package model

import "encoding/json"

// Response 回复
type Response struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func NewResponse(t string, data interface{}) *Response {
	return &Response{Type: t, Data: data}
}

func (r *Response) Json() []byte {
	str, err := json.Marshal(r)
	if err != nil {
		return []byte("")
	}

	return str
}
