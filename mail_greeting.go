package birthday_greeting_kata

type MailBirthdayGreeting struct {
	addressBook AddressBook
}

func (greeting MailBirthdayGreeting) execute() {
	friends := greeting.addressBook.Friends()
	sender := NewMailMessageSender()
	for _, friend := range friends {
		msgGenerator := NewMailGreetingMessageGenerator(friend)
		msg := msgGenerator.GenMessage()
		sender.send(msg)
	}
}

func NewMailBirthdayGreeting(book FileAddressBook) MailBirthdayGreeting {
	return MailBirthdayGreeting{addressBook: book}
}
