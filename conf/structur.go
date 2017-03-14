package conf

type Update struct {
	UpdateID int `json:"update_id"`
	Message Message `json:"message"`
	EditedMessage      Message            `json:"edited_message"`
	ChannelPost        Message            `json:"channel_post"`
	EditedChannelPost  Message            `json:"edited_channel_post"`
}

type Message struct {
	MessageID int `json:"message_id"`
	Date int `json:"date"`
	Text string `json:"text"`
	From User `json:"from"`
	Chat User `json:"chat"`
	Photo []PhotoSize     `json:"photo"`
	Entities []Entities `json:"entities"`
}

type User struct {
	ID int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UserName string `json:"username"`
}

type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"` // optional
}

type Entities struct {
	Length int `json:"length"`
	Offset int `json:"offset"`
	Type string `json:"type"`
}

