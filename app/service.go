package app

import "github.com/google/uuid"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ShortenURL(input string) (string, error) {
	// todo: validate input - users should not store random data - write test first

	// todo: validate input - length should not exceed X - write test first

	// generate a random id for this url
	id := uuid.New()

	// todo: actually store ID + input URL so that they can be resolved later - write test first

	return id.String(), nil
}
