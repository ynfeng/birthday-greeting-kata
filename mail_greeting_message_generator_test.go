package birthday_greeting_kata

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

func TestGenerateMessage(t *testing.T) {
	now := time.Now()
	month, _ := strconv.Atoi(now.Month().String())
	birthday := NewBirthday(now.Year(), month, now.Day())
	friend := NewFriend("Doe", "John", birthday, "Doe.John@foobar.com")
	msgGenerator := NewMailGreetingMessageGenerator(friend)

	message := msgGenerator.GenMessage()

	assert.Equal(t, "Happy birthday!", message.Subject)
	assert.Equal(t, "Happy birthday, dear John!", message.Content)
}
