package service

import (
	"context"
	"errors"
	"fmt"
	DataBase "github.com/k1v4/url_shortener/pkg/database"
	"github.com/k1v4/url_shortener/pkg/randomGen"
)

const (
	shortUrlLength = 10
	retries        = 5
)

var (
	errShortUrl = errors.New("short urlSave already exists. Please try again")
)

//go:generate go run github.com/vektra/mockery/v2@v2.43.2 --name=ILinksRepository
type ILinksRepository interface {
	SaveUrl(ctx context.Context, url, shortUrl string) (string, error)
	GetOrigin(ctx context.Context, shortUrl string) (string, error)
	GetShortUrl(ctx context.Context, url string) (string, error)
}

type LinksService struct {
	repo ILinksRepository
}

// NewLinksService конструктор для слоя сервиса
func NewLinksService(repo ILinksRepository) *LinksService {
	return &LinksService{repo: repo}
}

// SaveUrl реализует логику работы сервиса. Принимает полный url.
// Получает сокращение, используя паттерн отказоустойчивости
func (svc *LinksService) SaveUrl(ctx context.Context, url string) (string, error) {
	const op = "LinksService.SaveUrl"
	isUnique := false
	var shortUrl string

	// Пробуем ещё раз если сокращение получилось повторяющимся
	for i := 0; i < retries; i++ {
		shortUrl = randomGen.NewRandomString(shortUrlLength)

		// проверка на уникальность shortUrl
		if val, err := svc.repo.GetOrigin(ctx, shortUrl); err == nil && len(val) > 0 {
			continue
		}

		isUnique = true
		break
	}

	// не смогли создать уникальное сокращение
	if !isUnique {
		return "", fmt.Errorf("%w: %s", errShortUrl, url)
	}

	// вызываем репозиторий
	saveUrl, err := svc.repo.SaveUrl(ctx, url, shortUrl)
	if err != nil {
		// проверяем не был ли добавлен ранее этот адрес
		if errors.Is(err, DataBase.ErrUrlExists) {
			// возвращаем раннее сгенерированный shortUrl
			shortOrigin, err := svc.repo.GetShortUrl(ctx, url)
			if err != nil {
				return "", fmt.Errorf("%s: %w", op, err)
			}

			return shortOrigin, nil
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return saveUrl, nil
}

// GetOrigin реализует логику работы сервиса для получения оригинального url.
func (svc *LinksService) GetOrigin(ctx context.Context, shortUrl string) (string, error) {
	const op = "LinksService.GetOrigin"

	// вызываем репозиторий
	originUrl, err := svc.repo.GetOrigin(ctx, shortUrl)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return originUrl, nil
}
