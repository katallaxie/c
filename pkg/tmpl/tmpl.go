package tmpl

import (
	"io"
	"strings"
	"text/template"
)

// Parse ...
func Parse(src io.Reader, dst io.Writer, data interface{}) error {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, src)
	if err != nil {
		return err
	}

	tmpl, err := template.New("tmpl").Funcs(tmplFuncs).Parse(buf.String())
	if err != nil {
		return err
	}

	return tmpl.Execute(dst, data)
}
