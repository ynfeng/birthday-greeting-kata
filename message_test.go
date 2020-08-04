package birthday_greeting_kata

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateMessage(t *testing.T) {
	msg := NewMessage("subject","content")
	
	assert.Equal(t,msg.Subject,"subject")
	assert.Equal(t,msg.Content,"content")
}
