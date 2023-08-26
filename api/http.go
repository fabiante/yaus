package api

import (
	"errors"
	"github.com/fabiante/yaus/app"
	"github.com/gin-gonic/gin"
)

func SetupHTTPServer(api *gin.Engine, service *app.Service) *gin.Engine {
	api.POST("shorten", func(ctx *gin.Context) {
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

	return api
}
