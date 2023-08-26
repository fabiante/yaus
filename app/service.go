package app

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ShortenURL(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("invalid url: may not be empty")
	}
	if len([]byte(input)) > 512 {
		return "", fmt.Errorf("invalid url: may not be longer than 512 bytes")
	}
	if !strings.HasPrefix(input, "http") {
		return "", fmt.Errorf("invalid url: must begin with http")
	}

	// generate a random id for this url
	id := uuid.New()

	// todo: actually store ID + input URL so that they can be resolved later - write test first

	return id.String(), nil
}
