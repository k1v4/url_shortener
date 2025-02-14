package service

import (
	"context"
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
		name           string
		ctx            context.Context
		urlSave        string
		shortUrl       string
		isMockSave     bool
		isMockGet      bool
		isMockGetShort bool
		wantErrSave    bool
		wantErrGet     bool
		respErrSave    error
	}{
		{
			name:           "success",
			ctx:            context.Background(),
			urlSave:        "http://google.com",
			shortUrl:       "1234567890",
			isMockSave:     true,
			isMockGetShort: false,
			isMockGet:      true,
			wantErrSave:    false,
			wantErrGet:     false,
			respErrSave:    nil,
		},
		{
			name:           "url exist",
			ctx:            context.Background(),
			urlSave:        "http://google.com",
			shortUrl:       "1234567890",
			isMockSave:     true,
			isMockGetShort: true,
			isMockGet:      true,
			wantErrSave:    false,
			wantErrGet:     false,
			respErrSave:    DataBase.ErrUrlExists,
		},
		{
			name:           "GetShort error",
			ctx:            context.Background(),
			urlSave:        "http://google.com",
			shortUrl:       "1234567890",
			isMockSave:     true,
			isMockGetShort: true,
			isMockGet:      true,
			wantErrSave:    false,
			wantErrGet:     false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			iRepo := mocks.NewILinksRepository(t)

			if tt.isMockGet {
				iRepo.
					On("GetOrigin", mock.Anything, mock.Anything).
					Return("", nil)
			}

			if tt.isMockSave {
				iRepo.
					On("SaveUrl", mock.Anything, tt.urlSave, mock.Anything).
					Return(tt.shortUrl, tt.respErrSave)
			}

			if tt.isMockGetShort {
				iRepo.
					On("GetShortUrl", mock.Anything, tt.urlSave).
					Return(tt.respErrSave, nil)
			}

			s := NewLinksService(iRepo)

			_, err := s.SaveUrl(tt.ctx, tt.urlSave)
			if (err != nil) != tt.wantErrSave {
				t.Errorf("SaveUrl() error = %v, wantErr %v", err, tt.wantErrSave)
				return
			}

			//if !tt.wantErrSave && len(url) != 10 {
			//	fmt.Println(url)
			//	t.Errorf("GetOrigin(%s) length = %v, want 10", url, len(url))
			//	return
			//}
		})
	}
}
