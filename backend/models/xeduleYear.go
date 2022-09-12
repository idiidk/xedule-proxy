package models

type XeduleYear []struct {
	Oru            int         `json:"oru"`
	Year           int         `json:"year"`
	Schs           []int       `json:"schs"`
	Deps           interface{} `json:"deps"`
	Avis           interface{} `json:"avis"`
	PeriodCount    int         `json:"periodCount"`
	HasCalendar    bool        `json:"hasCalendar"`
	Cal            string      `json:"cal"`
	IStart         string      `json:"iStart"`
	IEnd           string      `json:"iEnd"`
	IStartOfDay    string      `json:"iStartOfDay"`
	IEndOfDay      string      `json:"iEndOfDay"`
	FirstDayOfWeek int         `json:"firstDayOfWeek"`
	LastDayOfWeek  int         `json:"lastDayOfWeek"`
	ID             string      `json:"id"`
}
