package removeunused

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hello3(t *testing.T) {
	greet := hello3()
	assert.EqualValues(t, greet, "Hello, world.")
}
