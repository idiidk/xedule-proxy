package models

type XeduleGroup []struct {
	AttDLs []interface{} `json:"attDLs"`
	Orus   []int         `json:"orus"`
	Code   string        `json:"code"`
	ID     string        `json:"id"`
}
