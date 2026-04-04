package service

type ShortenerService interface {
	Shorten(url string) (string, error)
}

type shortenerService struct{}

func NewShortenerService() ShortenerService {
	return &shortenerService{}
}

func (s *shortenerService) Shorten(url string) (string, error) {
	
	return "shortened_url", nil
}