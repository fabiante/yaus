package specs

type Service interface {
	ShortenURL(input string) (string, error)
}
