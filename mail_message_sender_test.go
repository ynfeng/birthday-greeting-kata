package birthday_greeting_kata

import (
	"testing"
)

func TestSendMessage(t *testing.T) {
	msg := NewMessage("subject", "content")

	sender := new(MailMessageSender)

	sender.send(msg)
}
