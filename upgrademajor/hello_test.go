package upgrademajor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_greeting(t *testing.T) {
	greet := Hello()
	assert.EqualValues(t, greet, "Hello, world.")
}
