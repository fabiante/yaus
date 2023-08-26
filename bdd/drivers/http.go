package drivers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HTTPDriver struct {
	BaseURL string
}

func NewHTTPDriver(baseURL string) *HTTPDriver {
	return &HTTPDriver{BaseURL: baseURL}
}

func (d *HTTPDriver) ShortenURL(input string) (string, error) {
	res, err := http.Post(fmt.Sprintf("%s/shorten", d.BaseURL), "application/json", toJSONReader(input))
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", fmt.Errorf("unexpected status %d", res.StatusCode)
	}

	var url string
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func toJSON(data any) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return bytes
}

func toJSONReader(data any) io.Reader {
	return bytes.NewReader(toJSON(data))
}
