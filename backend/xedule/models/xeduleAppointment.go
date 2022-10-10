package models

type XeduleAppointment struct {
	PublicationDate string `json:"iPublicationDate"`
	Concept         bool   `json:"concept"`
	Apps            []struct {
		Name        string      `json:"name"`
		Summary     string      `json:"summary"`
		Attention   string      `json:"attention"`
		MeetingInfo interface{} `json:"meetingInfo"`
		Start       string      `json:"iStart"`
		End         string      `json:"iEnd"`
		Atts        []int       `json:"atts"`
		ID          string      `json:"id"`
	} `json:"apps"`
	ID string `json:"id"`
}
