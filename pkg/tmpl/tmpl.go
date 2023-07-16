package tmpl

import (
	"bytes"
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

	var out bytes.Buffer
	tmpl, err := template.New("tmpl").Funcs(tmplFuncs).Parse(buf.String())
	if err != nil {
		return err
	}

	err = tmpl.Execute(&out, data)
	if err != nil {
		return err
	}

	b := bytes.ReplaceAll(out.Bytes(), []byte("<no value>"), []byte(""))
	_, err = io.Copy(dst, bytes.NewReader(b))
	if err != nil {
		return err
	}

	return nil
}
