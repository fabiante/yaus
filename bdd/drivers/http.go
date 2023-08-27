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
	client  *http.Client
}

func NewHTTPDriver(baseURL string) *HTTPDriver {
	return &HTTPDriver{
		BaseURL: baseURL,
		client:  http.DefaultClient,
	}
}

func (d *HTTPDriver) ShortenURL(input string) (string, error) {
	res, err := d.client.Post(fmt.Sprintf("%s/shorten", d.BaseURL), "application/json", toJSONReader(input))
	if err != nil {
		return "", err
	}
	switch res.StatusCode {
	case 200:
		break
	case 400:
		return "", fmt.Errorf("invalid url (api responded with status 400)")
	default:
		return "", fmt.Errorf("unexpected status %d", res.StatusCode)
	}

	var url string
	err = json.NewDecoder(res.Body).Decode(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (d *HTTPDriver) Resolve(short string) (string, error) {
	panic("driver not yet implemented")
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
