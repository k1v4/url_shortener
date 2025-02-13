package in_memory

import (
	"context"
	DataBase "github.com/k1v4/url_shortener/pkg/database"
	"sync"
)

type LinksRepository struct {
	urlStore      map[string]string
	shortUrlStore map[string]string
	mu            sync.Mutex
}

func NewLinksRepository() *LinksRepository {
	return &LinksRepository{
		urlStore:      make(map[string]string),
		shortUrlStore: make(map[string]string),
		mu:            sync.Mutex{},
	}
}

func (l *LinksRepository) SaveUrl(ctx context.Context, url, shortUrl string) (string, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if _, ok := l.urlStore[url]; ok {
		return "", DataBase.ErrUrlExists
	}

	l.urlStore[url] = shortUrl
	l.shortUrlStore[shortUrl] = url

	return shortUrl, nil
}

func (l *LinksRepository) GetOrigin(ctx context.Context, shortUrl string) (string, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var value string
	var ok bool
	if value, ok = l.shortUrlStore[shortUrl]; !ok {
		return "", DataBase.ErrUrlNotFound
	}

	return value, nil
}

func (l *LinksRepository) GetShortUrl(ctx context.Context, url string) (string, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	var value string
	var ok bool
	if value, ok = l.urlStore[url]; !ok {
		return "", DataBase.ErrUrlNotFound
	}

	return value, nil

}
