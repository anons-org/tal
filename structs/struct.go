package structs


type MessageSender struct {
	UserId int `json:"user_id"`
}

type Event struct {
	MessageType string `json:"message_type"`
	PostType string `json:"post_type"`
	RawMessage string `json:"raw_message"`
	Message string `json:"message"`
	Sender MessageSender
}