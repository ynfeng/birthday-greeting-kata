package birthday_greeting_kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateBirthday(t *testing.T) {
	birthday := NewBirthday(2020, 1, 1)

	assert.Equal(t, 2020, birthday.Year)
	assert.Equal(t, 1, birthday.Month)
	assert.Equal(t, 1, birthday.Day)
}

