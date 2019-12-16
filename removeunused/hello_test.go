package removeunused

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Testhello3(t *testing.T) {
	greet := hello3()
	assert.EqualValues(t, greet, "Hello")
}
