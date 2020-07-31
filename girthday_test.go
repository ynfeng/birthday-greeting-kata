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

func TestSameBirthday(t *testing.T) {
	birthday1 := NewBirthday(2020, 10, 1)
	birthday2 := NewBirthday(2020, 10, 1)
	birthday3 := NewBirthday(2020, 1, 1)

	assert.Equal(t, true, birthday1.isSame(birthday2))
	assert.Equal(t, false, birthday2.isSame(birthday3))
}
