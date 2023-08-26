package specs

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAll(t *testing.T, service Service) {
	t.Run("feature: shortening a link link", func(t *testing.T) {
		t.Run("given a valid URL will produce a shorter URL", func(t *testing.T) {
			input := "https://en.wikipedia.org/wiki/Go_(programming_language)#Types"
			url, err := service.ShortenURL(input)
			require.NoError(t, err)
			require.NotEmpty(t, url)
			require.Less(t, len(url), len(input)) // produce a shorter url
		})
	})
}
