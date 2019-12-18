package createmodule

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_greeting(t *testing.T) {
	greeting, err := greet("Joe")
	assert.EqualValues(t, "", err)
	assert.EqualValues(t, "hello: Joe", greeting)
}

func TestgreetingNoName(t *testing.T) {
	greeting, err := greet("e")
	assert.EqualValues(t, "No Name was given", err)
	assert.EqualValues(t, "hello: ", greeting)
}
func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := Hello("Joe"); got != want {
		t.Errorf("Hello()= got %q want %q", got, want)
	}
}
