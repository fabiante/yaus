package specs

import "testing"

func TestAll(t *testing.T, service Service) {
	t.Run("create link", func(t *testing.T) {
		CreateLink(t, service)
	})
}
