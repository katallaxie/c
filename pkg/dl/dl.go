package dl

import (
	"archive/zip"
	"bytes"
	"context"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/katallaxie/g/pkg/spec"
)

// Extract ...
func Extract(ctx context.Context, prefix string, url string) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return err
	}

	s := spec.Default()
	for _, zipFile := range zipReader.File {
		parts := strings.Split(zipFile.Name, string(os.PathSeparator))
		name := strings.Join(parts[1:], string(os.PathSeparator))

		if name != spec.DefaultFilename {
			continue
		}

		bb, err := readZipFile(zipFile)
		if err != nil {
			return err
		}

		err = s.UnmarshalYAML(bb)
		if err != nil {
			return err
		}
	}

	for _, zipFile := range zipReader.File {
		parts := strings.Split(zipFile.Name, string(os.PathSeparator))
		name := strings.Join(parts[1:], string(os.PathSeparator))

		if len(s.Templates) > 0 && s.TemplateMap()[name] == "" {
			continue
		}
	}

	return nil
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}
