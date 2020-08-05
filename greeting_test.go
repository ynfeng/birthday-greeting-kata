package birthday_greeting_kata

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	addressBook := NewFileAddressBook("./friends.data")
	greeting := NewMailBirthdayGreeting(addressBook)
	greeting.execute()
}
