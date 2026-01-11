package redisclient

type QueueMessage struct {
	ID      string `json:"id"`
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}
