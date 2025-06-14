package wsc

type Message struct {
	Sender      *Client
	MessageType int
	Message     []byte
}

type JsonMessageBody struct {
	Name    string `json:"sender_name"`
	Message string `json:"message"`
}
