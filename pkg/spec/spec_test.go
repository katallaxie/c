package spec_test

import (
	"testing"

	"github.com/katallaxie/g/pkg/spec"

	"github.com/stretchr/testify/assert"
)

func TestDefault(t *testing.T) {
	t.Parallel()

	s := spec.Default()
	assert.NotNil(t, s)
	assert.Equal(t, spec.DefaultVersion, s.Version)
}
