package idgenerator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUlIdGenerator_Generate(t *testing.T) {
	idGen := New()
	id := idGen.Generate()
	assert.NotEmpty(t, id)
	assert.Equal(t, len(id), 26)
}
