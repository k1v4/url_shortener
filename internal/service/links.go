package service

import (
	"context"
	"errors"
	"fmt"
	DataBase "github.com/k1v4/url_shortener/pkg/database"
	"github.com/k1v4/url_shortener/pkg/random"
)

const shortUrlLength = 10

type ILinksRepository interface {
	SaveUrl(ctx context.Context, url, shortUrl string) (string, error)
	GetOrigin(ctx context.Context, shortUrl string) (string, error)
	GetShortUrl(ctx context.Context, url string) (string, error)
}

type LinksService struct {
	repo ILinksRepository
}

func NewLinksService(repo ILinksRepository) *LinksService {
	return &LinksService{repo: repo}
}

func (svc *LinksService) SaveUrl(ctx context.Context, url string) (string, error) {
	const op = "LinksService.SaveUrl"

	shortUrl := random.NewRandomString(shortUrlLength)

	saveUrl, err := svc.repo.SaveUrl(ctx, url, shortUrl)
	if err != nil {
		if errors.Is(err, DataBase.ErrUrlExists) {
			origin, err := svc.repo.GetOrigin(ctx, shortUrl)
			if err != nil {
				return "", fmt.Errorf("%s: %w", op, err)
			}

			return origin, nil
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return saveUrl, nil
}

func (svc *LinksService) GetOrigin(ctx context.Context, shortUrl string) (string, error) {
	const op = "LinksService.GetOrigin"

	originUrl, err := svc.repo.GetOrigin(ctx, shortUrl)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return originUrl, nil
}
