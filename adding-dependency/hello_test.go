package helloadd

import "testing"

import "github.com/stretchr/testify/assert"

func TestHello(t *testing.T) {
	greeting := greet()

	assert.EqualValues(t, greeting, "Hello, world.")
}
