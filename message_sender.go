package birthday_greeting_kata

type MessageSender interface {
	send(msg Message) error
}
