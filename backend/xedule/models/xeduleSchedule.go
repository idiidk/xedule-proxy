package models

type XeduleSchedule struct {
	IPublicationDate string `json:"iPublicationDate"`
	Concept          bool   `json:"concept"`
	Apps             []struct {
		Name        string      `json:"name"`
		Summary     string      `json:"summary"`
		Attention   string      `json:"attention"`
		MeetingInfo interface{} `json:"meetingInfo"`
		IStart      string      `json:"iStart"`
		IEnd        string      `json:"iEnd"`
		Atts        []int       `json:"atts"`
		ID          string      `json:"id"`
	} `json:"apps"`
	ID string `json:"id"`
}
