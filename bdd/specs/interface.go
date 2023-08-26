package specs

type Service interface {
	ShortenURL(input string) (string, error)
	Resolve(short string) (string, error)
}
