package api

import (
	"errors"
	"fmt"
	"github.com/fabiante/yaus/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupHTTPServer(api *gin.Engine, service *app.Service) *gin.Engine {
	api.POST("/shorten", func(ctx *gin.Context) {
		var url string
		err := ctx.BindJSON(&url)
		if err != nil {
			_ = ctx.AbortWithError(400, err)
			return
		}

		shortened, err := service.ShortenURL(url)
		if err != nil {
			if errors.Is(err, app.ErrInvalidUrl) {
				_ = ctx.AbortWithError(400, err)
				return
			} else {
				_ = ctx.AbortWithError(500, err)
				return
			}
		}

		ctx.JSON(200, shortened)
	})

	api.GET("/s/:short", func(ctx *gin.Context) {
		short := ctx.Param("short")
		if short == "" {
			_ = ctx.AbortWithError(400, fmt.Errorf("no short to resolve"))
			return
		}

		shortened, err := service.Resolve(short)
		if err != nil {
			if errors.Is(err, app.ErrInvalidUrl) {
				_ = ctx.AbortWithError(400, err)
				return
			} else {
				_ = ctx.AbortWithError(500, err)
				return
			}
		}

		ctx.Redirect(http.StatusFound, shortened)
	})

	return api
}
