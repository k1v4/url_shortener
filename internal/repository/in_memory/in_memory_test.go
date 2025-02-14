package in_memory

import (
	"context"
	DataBase "github.com/k1v4/url_shortener/pkg/database"
	"github.com/k1v4/url_shortener/pkg/randomGen"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinksRepository_SaveUrl(t *testing.T) {
	repo := NewLinksRepository()
	ctx := context.Background()
	url := "http://www.google.com"
	shortUrl := randomGen.NewRandomString(10)

	saveUrl, err := repo.SaveUrl(ctx, url, shortUrl)
	assert.Nil(t, err)
	assert.NotEmpty(t, saveUrl)
	assert.Equal(t, shortUrl, saveUrl)

	saveUrlReply, err := repo.SaveUrl(ctx, url, shortUrl)
	assert.Error(t, err)
	assert.ErrorIs(t, err, DataBase.ErrUrlExists)
	assert.Empty(t, saveUrlReply)
}

func TestLinksRepository_GetOrigin(t *testing.T) {
	repo := NewLinksRepository()
	ctx := context.Background()

	url := "http://www.google.com"
	shortUrl := randomGen.NewRandomString(10)

	saveUrl, err := repo.SaveUrl(ctx, url, shortUrl)
	assert.Nil(t, err)
	assert.NotEmpty(t, saveUrl)
	assert.Equal(t, shortUrl, saveUrl)

	getOrigin, err := repo.GetOrigin(ctx, shortUrl)
	assert.Nil(t, err)
	assert.NotEmpty(t, getOrigin)
	assert.Equal(t, url, getOrigin)

	getOriginEmpty, err := repo.GetOrigin(ctx, "1234567890")
	assert.Error(t, err)
	assert.ErrorIs(t, err, DataBase.ErrUrlNotFound)
	assert.Empty(t, getOriginEmpty)
}

func TestLinksRepository_GetShortUrl(t *testing.T) {
	repo := NewLinksRepository()
	ctx := context.Background()

	url := "http://www.google.com"
	shortUrl := randomGen.NewRandomString(10)

	saveUrl, err := repo.SaveUrl(ctx, url, shortUrl)
	assert.Nil(t, err)
	assert.NotEmpty(t, saveUrl)
	assert.Equal(t, shortUrl, saveUrl)

	getShort, err := repo.GetShortUrl(ctx, url)
	assert.Nil(t, err)
	assert.NotEmpty(t, getShort)
	assert.Equal(t, shortUrl, getShort)

	getShortEmpty, err := repo.GetShortUrl(ctx, "https:/hahahah.com")
	assert.Error(t, err)
	assert.ErrorIs(t, err, DataBase.ErrUrlNotFound)
	assert.Empty(t, getShortEmpty)
}
