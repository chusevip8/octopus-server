package response

type Text struct {
	Text string `json:"text"`
}

type CommentRes struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Text Text   `json:"text"`
	Date string `json:"date"`
	Mine bool   `json:"mine"`
	Img  string `json:"img"`
}
