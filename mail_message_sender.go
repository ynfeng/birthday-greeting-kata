package birthday_greeting_kata

import "fmt"

type MailMessageSender struct {
}

func (sender MailMessageSender) send(msg Message) error {
	fmt.Println("subject:", msg.Subject, "content:", msg.Content)
	return nil
}

func NewMailMessageSender() MessageSender {
	return MailMessageSender{}
}
