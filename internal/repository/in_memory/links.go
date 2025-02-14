package in_memory

import (
	"context"
	DataBase "github.com/k1v4/url_shortener/pkg/database"
	"sync"
)

type LinksRepository struct {
	urlStore      map[string]string // Хранилище основное
	shortUrlStore map[string]string // Хранилище для контроля коротких ссылок
	mu            sync.Mutex        // мутекс для синхронизации
}

// NewLinksRepository Инициализация репозитория in-memory
func NewLinksRepository() *LinksRepository {
	return &LinksRepository{
		urlStore:      make(map[string]string),
		shortUrlStore: make(map[string]string),
		mu:            sync.Mutex{},
	}
}

// SaveUrl сохраняем в хранилище
// в качестве параметров передаём ctx context.Context, далее полный url
// и последним аргументом короткая ссылка
func (l *LinksRepository) SaveUrl(ctx context.Context, url, shortUrl string) (string, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// проверка на наличие в хранилище
	if _, ok := l.urlStore[url]; ok {
		return "", DataBase.ErrUrlExists
	}

	// добавление значений
	l.urlStore[url] = shortUrl
	l.shortUrlStore[shortUrl] = url

	return shortUrl, nil
}

// GetOrigin возвращает полный url,
// принимает на вход контекст и сокращённый url
func (l *LinksRepository) GetOrigin(ctx context.Context, shortUrl string) (string, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// проверка на наличие в хранилище
	var value string
	var ok bool
	if value, ok = l.shortUrlStore[shortUrl]; !ok {
		return "", DataBase.ErrUrlNotFound
	}

	return value, nil
}

// GetShortUrl возвращает уже имеющийся сокращенный url
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
