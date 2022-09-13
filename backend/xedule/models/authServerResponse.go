package models

type AuthServerResponse struct {
	Cookies []struct {
		Name   string `json:"name"`
		Value  string `json:"value"`
		Domain string `json:"domain"`
	} `json:"cookies"`
	Config struct {
		XeduleURL string `json:"xeduleUrl"`
	} `json:"config"`
}
