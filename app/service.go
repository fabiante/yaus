package app

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ShortenURL(input string) (string, error) {
	return input, nil
}
