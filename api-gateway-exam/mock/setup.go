package mock

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

const testHost = "http://localhost:9191"

func NewRequest(method string, uri string, body []byte) *http.Request {
	req, err := http.NewRequest(method, testHost+uri, nil)
	if err != nil {
		return nil
	}

	req.Header.Set("Host", "localhost")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36")
	req.Header.Set("X-Forwarded-For", "79.104.42.249")

	if body != nil {
		req.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	return req
}

func OpenFile(fileName string) ([]byte, error) {
	return os.ReadFile(fileName)
}
