package drivers

import "github.com/fabiante/yaus/app"

type DirectDriver struct {
	service *app.Service
}

func NewDirectDriver(service *app.Service) *DirectDriver {
	return &DirectDriver{service: service}
}

func (d *DirectDriver) ShortenURL(input string) (string, error) {
	return d.service.ShortenURL(input)
}
