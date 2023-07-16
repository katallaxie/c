package tmpl

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTmplFuncs(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		index string
		want  interface{}
	}{
		{
			name:  "ARCH",
			index: "ARCH",
			want:  runtime.GOARCH,
		},
		{
			name:  "OS",
			index: "OS",
			want:  runtime.GOOS,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tmplFuncs[tt.index].(func() string)()
			assert.Equal(t, tt.want, got)
		})
	}
}
