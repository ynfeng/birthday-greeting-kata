package birthday_greeting_kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateFriend(t *testing.T) {
	friend := NewFriend("Doe", "John", NewBirthday(1982, 10, 8), "john.doe@foobar.com")

	assert.Equal(t, "Doe", friend.LastName)
	assert.Equal(t, "John", friend.FirstName)
	assert.Equal(t, NewBirthday(1982, 10, 8), friend.Birthday)
	assert.Equal(t, "john.doe@foobar.com", friend.Email)
}
