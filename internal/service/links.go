package service

import "context"

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

func (svc *LinksService) SaveUrl(ctx context.Context, url, shortUrl string) (string, error) {
	saveUrl, err := svc.repo.SaveUrl(ctx, url, shortUrl)
	if err != nil {
		return "", err
	}

	return saveUrl, nil
}

func (svc *LinksService) GetOrigin(ctx context.Context, shortUrl string) (string, error) {
	originUrl, err := svc.repo.GetOrigin(ctx, shortUrl)
	if err != nil {
		return "", err
	}

	return originUrl, nil
}
