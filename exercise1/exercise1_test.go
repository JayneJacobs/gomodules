package exercise1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_glassdoor(t *testing.T) {
	greetTest := glassdoor()
	assert.EqualValues(t, greetTest, "I can eat glass and it doesn't hurt me.")
}
