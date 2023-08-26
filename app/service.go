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
	ErrInvalidUrl   = errors.New("invalid url")
	ErrInvalidShort = errors.New("invalid short")
	// ErrNotFound communicates that a shortened url can not
	// be resolved to the original url.
	// It could also mean that there is simply no shortened url.
	ErrNotFound = errors.New("short not found")
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

// Resolve resolves the given short to the original URL.
//
// Returns ErrInvalidShort if the input has some invalid format.
//
// Returns ErrNotFound if resolving failed - meaning no original URL is found.
func (s *Service) Resolve(short string) (string, error) {
	id, err := uuid.Parse(short)
	if err != nil {
		return "", fmt.Errorf("%w: short must be an uuid", ErrInvalidShort)
	}

	s.urlsLock.RLock()
	url, found := s.urls[id]
	s.urlsLock.RUnlock()

	if !found {
		return "", fmt.Errorf("%w: id %s", ErrNotFound, id)
	} else {
		return url, nil
	}
}
