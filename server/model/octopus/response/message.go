package response

type MessageRes struct {
	Name string `json:"name"`
	Text Text   `json:"text"`
	Date string `json:"date"`
	Mine bool   `json:"mine"`
}
