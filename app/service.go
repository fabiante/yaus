package app

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"sync"
)

type Service struct {
	urls     map[uuid.UUID]string
	urlsLock sync.RWMutex
}

func NewService() *Service {
	return &Service{
		urls: make(map[uuid.UUID]string),
	}
}

var (
	ErrInvalidUrl = errors.New("invalid url")
)

// ShortenURL shortens the given input, which must be a URL.
//
// ErrInvalidUrl is returned if the input is invalid.
func (s *Service) ShortenURL(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("%w: may not be empty", ErrInvalidUrl)
	}
	if len([]byte(input)) > 512 {
		return "", fmt.Errorf("%w: may not be longer than 512 bytes", ErrInvalidUrl)
	}
	if !strings.HasPrefix(input, "http") {
		return "", fmt.Errorf("%w: must begin with http", ErrInvalidUrl)
	}

	// generate a random id for this url
	id := uuid.New()

	s.urlsLock.Lock()
	s.urls[id] = input
	s.urlsLock.Unlock()

	return id.String(), nil
}
