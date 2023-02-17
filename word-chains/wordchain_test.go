package wordchain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsEmptyListForInvalidStartEnd(t *testing.T) {
	result := Solve("cat", "sheep")
	assert.Empty(t, result)
}

