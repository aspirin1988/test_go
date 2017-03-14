package conf

type Update struct {
	UpdateID int `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	MessageID int `json:"message_id"`
	Date int `json:"date"`
	Text string `json:"text"`
	From User `json:"from"`
	Chat User `json:"chat"`
	Entities []Entities `json:"entities"`
}

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"username"`
}

type Entities struct {
	Length int `json:"length"`
	Offset int `json:"offset"`
	Type string `json:"type"`
}
