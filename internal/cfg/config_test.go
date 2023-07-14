package cfg_test

import (
	"testing"

	"github.com/katallaxie/g/internal/cfg"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	c := cfg.New()
	assert.NotNil(t, c)
	assert.False(t, c.Verbose)
	assert.Equal(t, "", c.Template)
	assert.False(t, c.Force)
}

func TestDefault(t *testing.T) {
	t.Parallel()

	c := cfg.Default()
	assert.NotNil(t, c)
	assert.False(t, c.Verbose)
	assert.Equal(t, "", c.Template)
	assert.False(t, c.Force)
}
