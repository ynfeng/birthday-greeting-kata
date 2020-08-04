package birthday_greeting_kata

type Message struct {
	Subject string
	Content string
}

func NewMessage(subject string, content string) Message {
	return Message{Subject: subject, Content: content}
}
