package specs

import (
	"github.com/carlmjohnson/be"
	"testing"
)

func CreateLink(t testing.TB, service Service) {
	url, err := service.ShortenURL("https://en.wikipedia.org/wiki/Go_(programming_language)#Types")
	be.NilErr(t, err)
	be.Unequal(t, "", url)
}
