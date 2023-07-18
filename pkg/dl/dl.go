package dl

import (
	"archive/zip"
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/katallaxie/g/pkg/spec"
)

// Extract ...
func Extract(ctx context.Context, url string) error {
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
		if zipFile.Name != spec.DefaultFilename {
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
