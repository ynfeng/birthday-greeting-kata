package birthday_greeting_kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFriendsFromFile(t *testing.T) {
	addressBook := NewFileAddressBook("./Friends.data")

	friends := addressBook.Friends()

	assert.Equal(t, len(friends), 2)

	john := friends[0]
	mary := friends[1]

	assert.Equal(t, "John", john.FirstName)
	assert.Equal(t, "Doe", john.LastName)
	assert.Equal(t, "john.doe@foobar.com", john.Email)
	assert.Equal(t, NewBirthday(1982, 10, 8), john.Birthday)
	assert.Equal(t, "Mary", mary.FirstName)
	assert.Equal(t, "Ann", mary.LastName)
	assert.Equal(t, "mary.ann@foobar.com", mary.Email)
	assert.Equal(t, NewBirthday(1975, 9, 11), mary.Birthday)
}
