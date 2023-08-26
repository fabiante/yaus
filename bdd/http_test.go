package bdd

import (
	"github.com/fabiante/yaus/api"
	"github.com/fabiante/yaus/app"
	"github.com/fabiante/yaus/bdd/drivers"
	"github.com/fabiante/yaus/bdd/specs"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"testing"
)

func TestHTTP(t *testing.T) {
	server := startTestserver()
	defer server.Close()

	driver := drivers.NewHTTPDriver(server.URL)

	t.Run("create link", func(t *testing.T) {
		specs.CreateLink(t, driver)
	})
}

func startTestserver() *httptest.Server {
	service := app.NewService()
	gin.SetMode(gin.TestMode)
	server := api.SetupHTTPServer(gin.New(), service)
	return httptest.NewServer(server)
}
