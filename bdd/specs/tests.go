package specs

import (
	"github.com/carlmjohnson/be"
	"testing"
)

func TestAll(t *testing.T, service Service) {
	t.Run("feature: shortening a link link", func(t *testing.T) {
		t.Run("given a valid URL", func(t *testing.T) {
			input := "https://en.wikipedia.org/wiki/Go_(programming_language)#Types"
			url, err := service.ShortenURL(input)
			be.NilErr(t, err)
			be.Unequal(t, "", url)
		})
	})
}
