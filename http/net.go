package http

import (
	"context"
	"io"
	"net/http"
)

// DownloadImage download image to bytes
func DownloadImage(_ context.Context, url string) ([]byte, error) {
	var (
		resp *http.Response
		err  error
	)

	if resp, err = http.Get(url); err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
