package specs

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAll(t *testing.T, service Service) {
	t.Run("feature: shortening a link link", func(t *testing.T) {
		t.Run("given a valid URL will produce a shorter URL which resolves to the original", func(t *testing.T) {
			// create a shortened link
			input := "https://en.wikipedia.org/wiki/Go_(programming_language)#Types"
			url, err := service.ShortenURL(input)
			requireShortenedURL(t, err, url, input)

			// use the shortened link to ensure it resolves to the original input
			resolved, err := service.Resolve(url)
			require.NoError(t, err)
			require.NotEmpty(t, resolved)
			require.Equal(t, input, resolved)
		})

		t.Run("given invalid input will error out", func(t *testing.T) {
			data := []string{
				"",
				"not a url",
			}

			for i, input := range data {
				t.Run(fmt.Sprintf("data[%d]", i), func(t *testing.T) {
					_, err := service.ShortenURL(input)
					require.ErrorContains(t, err, "invalid url")
				})
			}

			// These tests ensure that there is some kind of input size limit.
			// This must not exactly be 512 bytes or characters. It should just generally
			// prevent the service from being used as a database.
			t.Run("validates length", func(t *testing.T) {
				t.Run("allows 512 bytes", func(t *testing.T) {
					input := "http://8123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678"
					url, err := service.ShortenURL(input)
					requireShortenedURL(t, err, url, input)
				})

				t.Run("rejects 513 bytes", func(t *testing.T) {
					input := "http://78123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678123456781234567812345678"
					_, err := service.ShortenURL(input)
					require.ErrorContains(t, err, "invalid url")
				})
			})
		})
	})
}

func requireShortenedURL(t *testing.T, err error, url string, input string) {
	require.NoError(t, err)
	require.NotEmpty(t, url)
	require.Less(t, len(url), len(input)) // produce a shorter url
}
