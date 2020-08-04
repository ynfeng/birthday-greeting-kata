package birthday_greeting_kata

import "fmt"

type MailGreetingMessageGenerator struct {
	friend Friend
}

func (generator MailGreetingMessageGenerator) GenMessage() Message {
	content := fmt.Sprintf("Happy birthday, dear %s!", generator.friend.FirstName)
	return Message{Subject: "Happy birthday!", Content: content}
}

func NewMailGreetingMessageGenerator(friend Friend) MessageGenerator {
	return MailGreetingMessageGenerator{friend: friend}
}
