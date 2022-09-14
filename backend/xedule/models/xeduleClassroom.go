package models

type XeduleClassroom struct {
	Orus []int    `json:"orus"`
	Code string   `json:"code"`
	ID   string   `json:"id"`
	Tsss []string `json:"tsss,omitempty"`
}
