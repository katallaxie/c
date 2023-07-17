package tmpl_test

import (
	"bytes"
	"io"
	"runtime"
	"strings"
	"testing"

	"github.com/katallaxie/g/pkg/tmpl"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		name string
		in   io.Reader
		out  string
		data interface{}
	}{
		{
			name: "replace values",
			in:   strings.NewReader("{{ OS }}"),
			out:  runtime.GOOS,
		},
		{
			name: "replace no values in data",
			in:   strings.NewReader("{{ .NOVALUE }}"),
			out:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := tmpl.Parse(tt.in, &buf, tt.data)
			assert.NoError(t, err)
			assert.Equal(t, tt.out, buf.String())
		})
	}
}
