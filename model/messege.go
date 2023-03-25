package model

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type ChatRoom struct {
	Messages []Message `json:"messages"`
}
