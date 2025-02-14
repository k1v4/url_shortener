package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/k1v4/url_shortener/internal/service/mocks"
	DataBase "github.com/k1v4/url_shortener/pkg/database"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestService_GetOrigin(t *testing.T) {
	cases := []struct {
		name     string
		ctx      context.Context
		shortUrl string
		url      string
		isMock   bool
		wantErr  bool
		respErr  error
	}{
		{
			name:     "success",
			ctx:      context.Background(),
			shortUrl: "",
			url:      "http://google.com",
			isMock:   true,
			wantErr:  false,
			respErr:  nil,
		},
		{
			name:     "error",
			ctx:      context.Background(),
			shortUrl: "",
			url:      "http://google.com",
			isMock:   true,
			wantErr:  true,
			respErr:  fmt.Errorf("%s:error", "LinksService.GetOrigin"),
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			iRepo := mocks.NewILinksRepository(t)

			if tt.isMock {
				iRepo.
					On("GetOrigin", mock.Anything, tt.shortUrl).
					Return(tt.url, tt.respErr)
			}

			s := NewLinksService(iRepo)

			origin, err := s.GetOrigin(tt.ctx, tt.shortUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && len(origin) == 0 {
				t.Errorf("GetOrigin(%s) length = %v, want not nil", origin, len(origin))
				return
			}
		})
	}
}

func TestService_SaveUrl(t *testing.T) {
	cases := []struct {
		name             string
		ctx              context.Context
		urlSave          string
		shortUrl         string
		isMockGetOrigin  bool
		isMockSave       bool
		isMockGetShort   bool
		wantErrSave      bool
		wantErrGetOrigin bool
		wantErrGetShort  bool
		wantTotalErr     bool
		respErrGetOrigin error
		respErrSave      error
		respErrGetSHort  error
		getOriginResp    string
	}{
		{
			name:             "success",
			ctx:              context.Background(),
			urlSave:          "http://google.com",
			shortUrl:         "1234567890",
			isMockGetOrigin:  true,
			isMockSave:       true,
			isMockGetShort:   false,
			wantErrSave:      false,
			wantErrGetOrigin: false,
			wantErrGetShort:  false,
			wantTotalErr:     false,
			respErrSave:      nil,
			respErrGetOrigin: nil,
			respErrGetSHort:  nil,
			getOriginResp:    "",
		},
		{
			name:             "Duplicate url",
			ctx:              context.Background(),
			urlSave:          "http://google.com",
			shortUrl:         "1234567890",
			isMockGetOrigin:  true,
			isMockSave:       true,
			isMockGetShort:   true,
			wantErrGetOrigin: false,
			wantErrSave:      true,
			wantErrGetShort:  false,
			wantTotalErr:     false,
			respErrGetOrigin: nil,
			respErrSave:      DataBase.ErrUrlExists,
			respErrGetSHort:  nil,
			getOriginResp:    "",
		},
		{
			name:             "Duplicate url",
			ctx:              context.Background(),
			urlSave:          "http://google.com",
			shortUrl:         "1234567890",
			isMockGetOrigin:  true,
			isMockSave:       true,
			isMockGetShort:   true,
			wantErrGetOrigin: false,
			wantErrSave:      true,
			wantErrGetShort:  true,
			wantTotalErr:     true,
			respErrGetOrigin: nil,
			respErrSave:      DataBase.ErrUrlExists,
			respErrGetSHort:  errors.New("some error"),
			getOriginResp:    "",
		},
		{
			name:             "Duplicate url",
			ctx:              context.Background(),
			urlSave:          "http://google.com",
			shortUrl:         "1234567890",
			isMockGetOrigin:  true,
			isMockSave:       true,
			isMockGetShort:   false,
			wantErrGetOrigin: false,
			wantErrSave:      true,
			wantErrGetShort:  false,
			wantTotalErr:     true,
			respErrGetOrigin: nil,
			respErrSave:      errors.New("some error"),
			respErrGetSHort:  nil,
			getOriginResp:    "",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			iRepo := mocks.NewILinksRepository(t)

			if tt.isMockGetOrigin {
				iRepo.
					On("GetOrigin", mock.Anything, mock.Anything).
					Return(tt.getOriginResp, tt.respErrGetOrigin)
			}

			if tt.isMockSave {
				iRepo.
					On("SaveUrl", mock.Anything, tt.urlSave, mock.Anything).
					Return(tt.shortUrl, tt.respErrSave)
			}

			if tt.isMockGetShort {
				iRepo.
					On("GetShortUrl", tt.ctx, tt.urlSave).
					Return(tt.shortUrl, tt.respErrGetSHort)
			}

			s := NewLinksService(iRepo)

			url, err := s.SaveUrl(tt.ctx, tt.urlSave)
			if (err != nil) != tt.wantTotalErr {
				t.Errorf("SaveUrl() error = %v, wantErr %v", err, tt.wantErrSave)
				return
			}

			if !tt.wantTotalErr && len(url) != 10 {
				fmt.Println(url)
				t.Errorf("GetOrigin(%s) length = %v, want 10", url, len(url))
				return
			}
		})
	}
}
