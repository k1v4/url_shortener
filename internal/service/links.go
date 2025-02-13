package service

import (
	"context"
	"errors"
	"fmt"
	DataBase "github.com/k1v4/url_shortener/pkg/database"
	"github.com/k1v4/url_shortener/pkg/random"
)

const (
	shortUrlLength = 10
	retries        = 5
)

var (
	errShortUrl = errors.New("short url already exists. Please try again")
)

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
	isUnique := false
	var shortUrl string

	for i := 0; i < retries; i++ {
		shortUrl = random.NewRandomString(shortUrlLength)

		// проверка на уникальность shortUrl
		if val, err := svc.repo.GetOrigin(ctx, shortUrl); err == nil && len(val) > 0 {
			continue
		}

		isUnique = true
		break
	}

	if !isUnique {
		return "", fmt.Errorf("%w: %s", errShortUrl, url)
	}

	saveUrl, err := svc.repo.SaveUrl(ctx, url, shortUrl)
	if err != nil {
		if errors.Is(err, DataBase.ErrUrlExists) {
			origin, err := svc.repo.GetShortUrl(ctx, url)
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
