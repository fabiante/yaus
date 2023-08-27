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
	// Custom client which does not follow redirects - required since
	// the API makes use of 3xx redirections.
	client := &http.Client{
		Transport: http.DefaultTransport,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	return &HTTPDriver{
		BaseURL: baseURL,
		client:  client,
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
	res, err := d.client.Get(fmt.Sprintf("%s/s/%s", d.BaseURL, short))
	if err != nil {
		return "", err
	}
	switch res.StatusCode {
	case http.StatusPermanentRedirect:
		break
	default:
		return "", fmt.Errorf("unexpected status %d", res.StatusCode)
	}

	url := res.Header.Get("Location")

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
