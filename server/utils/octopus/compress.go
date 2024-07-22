package octopus

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

func Compress(data []byte) ([]byte, error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	defer func() {
		if err := w.Close(); err != nil {
			fmt.Printf("Failed to close gzip writer: %s\n", err)
		}
	}()

	if _, err := w.Write(data); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func Decompress(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := r.Close(); err != nil {
			fmt.Printf("Failed to close gzip reader: %s\n", err)
		}
	}()

	return io.ReadAll(r)
}
