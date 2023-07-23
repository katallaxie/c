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

func TestUnmarshalYAML(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name      string
		spec      []byte
		err       error
		expectErr bool
	}{
		{
			name: "empty",
			spec: []byte(
				`version: 1
name: test`,
			),
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := spec.Default()
			assert.NotNil(t, s)
			err := s.UnmarshalYAML(tt.spec)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
