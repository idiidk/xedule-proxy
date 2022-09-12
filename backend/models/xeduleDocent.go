package models

type XeduleDocent []struct {
	AttTLs []interface{} `json:"attTLs"`
	AttGLs []interface{} `json:"attGLs"`
	Name   string        `json:"name"`
	Orus   []int         `json:"orus"`
	Tsss   []string      `json:"tsss,omitempty"`
	Code   string        `json:"code"`
	ID     string        `json:"id"`
}
