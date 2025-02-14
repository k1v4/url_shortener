package grpc

import (
	"context"
	linkv1 "github.com/k1v4/url_shortener/pkg/api/link"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	urlCheck "net/url"
	"strings"
)

const (
	MsgInvalidShortUrl  = "invalid shortUrl"
	MsgInvalidOriginUrl = "invalid origin url"
)

//go:generate go run github.com/vektra/mockery/v2@v2.43.2 --name=ILinksService
type ILinksService interface {
	SaveUrl(ctx context.Context, url string) (string, error)
	GetOrigin(ctx context.Context, shortUrl string) (string, error)
}

type LinksService struct {
	linkv1.UnimplementedUrlShortenerServer
	service ILinksService
}

// NewLinksService конструктор
func NewLinksService(service ILinksService) *LinksService {
	return &LinksService{service: service}
}

// SaveUrl функция транспортного слоя для сохранения url в репозиторий
func (s *LinksService) SaveUrl(ctx context.Context, req *linkv1.SaveUrlRequest) (*linkv1.SaveUrlResponse, error) {
	url := req.Url
	// проверяем на пустоту
	if len(strings.TrimSpace(url)) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty url")
	}

	// проверяем на наличие
	if _, err := urlCheck.ParseRequestURI(url); err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid url")
	}

	// обращаемся к сервису
	saveUrl, err := s.service.SaveUrl(ctx, url)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &linkv1.SaveUrlResponse{ShortUrl: saveUrl}, nil
}

// GetOrigin функция транспортного слоя для получения сокращенного url в репозиторий
func (s *LinksService) GetOrigin(ctx context.Context, req *linkv1.GetOriginRequest) (*linkv1.GetOriginResponse, error) {
	shortUrl := req.GetShortUrl()

	// проверка на пустоту
	if len(strings.TrimSpace(shortUrl)) == 0 {
		return nil, status.Error(codes.InvalidArgument, MsgInvalidShortUrl)
	}

	// по условию длина должна быть == 10
	if len(shortUrl) != 10 {
		return nil, status.Error(codes.InvalidArgument, MsgInvalidShortUrl)
	}

	// обращаемся к сервису
	origin, err := s.service.GetOrigin(ctx, shortUrl)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &linkv1.GetOriginResponse{Url: origin}, nil
}
