package grpc

import (
	"context"
	"errors"
	linkv1 "github.com/k1v4/url_shortener/pkg/api/link"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidUrl         = errors.New("invalid origin url")
	ErrUserExist          = errors.New("user exist")
)

type ILinksService interface {
	SaveUrl(ctx context.Context, url, shortUrl string) (string, error)
	GetOrigin(ctx context.Context, shortUrl string) (string, error)
}

type LinksService struct {
	linkv1.UnimplementedUrlShortenerServer
	service ILinksService
}

func NewLinksService(service ILinksService) *LinksService {
	return &LinksService{service: service}
}

func (s *LinksService) SaveUrl(ctx context.Context, req *linkv1.SaveUrlRequest) (*linkv1.SaveUrlResponse, error) {
	url := req.Url
	if len(strings.TrimSpace(url)) == 0 {
		return nil, status.Error(codes.InvalidArgument, "empty url")
	}

}

func (s *LinksService) GetOrigin(ctx context.Context, req *linkv1.GetOriginRequest) (*linkv1.GetOriginResponse, error) {

}
