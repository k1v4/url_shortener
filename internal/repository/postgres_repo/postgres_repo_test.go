package postgres_repo

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/joho/godotenv"
	"github.com/k1v4/url_shortener/internal/config"
	DataBase "github.com/k1v4/url_shortener/pkg/database"
	"github.com/k1v4/url_shortener/pkg/database/postgres"
	"github.com/k1v4/url_shortener/pkg/randomGen"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLinksRepository(t *testing.T) {
	ctx := context.Background()

	err := godotenv.Load("../../../.env") // Укажите полный путь к файлу .env
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	os.Setenv("POSTGRES_HOST", "localhost")

	cfg := config.New()
	if cfg == nil {
		panic("load config fail")
	}

	urlDB := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBConfig.UserName,
		cfg.DBConfig.Password,
		cfg.DBConfig.Host,
		cfg.DBConfig.Port,
		cfg.DBConfig.DbName,
	)

	pg, err := postgres.New(urlDB, postgres.MaxPoolSize(cfg.DBConfig.PoolMax))
	if err != nil {
		panic(err)
	}

	linksRepository := NewLinksRepository(pg)

	url := gofakeit.URL()
	shortUrl := randomGen.NewRandomString(10)

	saveUrl, err := linksRepository.SaveUrl(ctx, url, shortUrl)
	assert.Nil(t, err)
	assert.NotEmpty(t, saveUrl)
	assert.Equal(t, shortUrl, saveUrl)

	saveDuplicateUrl, err := linksRepository.SaveUrl(ctx, url, shortUrl)
	assert.Error(t, err)
	assert.ErrorIs(t, err, DataBase.ErrUrlExists)
	assert.Empty(t, saveDuplicateUrl)

	getShortUrl, err := linksRepository.GetShortUrl(ctx, url)
	assert.Nil(t, err)
	assert.NotEmpty(t, getShortUrl)
	assert.Equal(t, shortUrl, getShortUrl)

	getShortEmpty, err := linksRepository.GetShortUrl(ctx, "https:/hahahah.com")
	assert.Error(t, err)
	assert.Empty(t, getShortEmpty)

	getOrigin, err := linksRepository.GetOrigin(ctx, shortUrl)
	assert.Nil(t, err)
	assert.NotEmpty(t, getOrigin)
	assert.Equal(t, url, getOrigin)

	getOriginEmpty, err := linksRepository.GetOrigin(ctx, "1234567890")
	assert.Error(t, err)
	assert.Empty(t, getOriginEmpty)

	linksRepository.Close()

	url = gofakeit.URL()
	shortUrl = randomGen.NewRandomString(10)

	saveUrl, err = linksRepository.SaveUrl(ctx, url, shortUrl)
	assert.Error(t, err)
	assert.Empty(t, saveUrl)
}
