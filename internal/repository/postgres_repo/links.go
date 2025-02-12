package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	DataBase "github.com/k1v4/url_shortener/pkg/database"
	"github.com/k1v4/url_shortener/pkg/database/postgres"
)

type LinksRepository struct {
	*postgres.Postgres
}

func NewLinksRepository(postgres *postgres.Postgres) *LinksRepository {
	return &LinksRepository{postgres}
}

func (l *LinksRepository) SaveUrl(ctx context.Context, url, shortUrl string) (string, error) {
	const op = "repository.postgres.SaveUrl"

	_, err := l.Pool.Exec(ctx, "INSERT INTO links (url, short_url) VALUES ($1, $2)", url, shortUrl)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return "", DataBase.ErrUrlExists
		}

		return "", fmt.Errorf("%s: %w", op, err)
	}

	return shortUrl, nil
}

func (l *LinksRepository) GetOrigin(ctx context.Context, shortUrl string) (string, error) {
	const op = "repository.postgres.GetOrigin"

	var originUrl string
	err := l.Pool.QueryRow(ctx, "SELECT url FROM links WHERE short_url = $1", shortUrl).Scan(&originUrl)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return originUrl, nil
}

func (l *LinksRepository) GetShortUrl(ctx context.Context, url string) (string, error) {
	const op = "repository.postgres.GetShortUrl"

	var shortUrl string
	err := l.Pool.QueryRow(ctx, "SELECT short_url FROM links WHERE url = $1", url).Scan(&shortUrl)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return shortUrl, nil
}
