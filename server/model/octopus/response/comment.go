package response

type Text struct {
	Text string `json:"text"`
}

type CommentRes struct {
	Name string `json:"name"`
	Text Text   `json:"text"`
	Date string `json:"date"`
	Mine bool   `json:"mine"`
}
