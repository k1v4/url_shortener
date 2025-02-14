package grpc

import (
	"context"
	"github.com/k1v4/url_shortener/internal/transport/grpc/mocks"
	linkv1 "github.com/k1v4/url_shortener/pkg/api/link"
	"github.com/k1v4/url_shortener/pkg/logger"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestTransport_SaveUrl(t *testing.T) {
	cases := []struct {
		name     string
		req      *linkv1.SaveUrlRequest
		ctx      context.Context
		resp     *linkv1.SaveUrlResponse
		shortUrl string
		isMock   bool
		wantErr  bool
		respErr  error
	}{
		{
			name: "success",
			req: &linkv1.SaveUrlRequest{
				Url: "http://google.com",
			},
			ctx: context.Background(),
			resp: &linkv1.SaveUrlResponse{
				ShortUrl: "1234567890",
			},
			isMock:  true,
			wantErr: false,
			respErr: nil,
		},
		{
			name: "empty_url",
			req: &linkv1.SaveUrlRequest{
				Url: "",
			},
			ctx: context.Background(),
			resp: &linkv1.SaveUrlResponse{
				ShortUrl: "",
			},
			isMock:  false,
			wantErr: true,
			respErr: status.Error(codes.InvalidArgument, "empty url"),
		},
		{
			name: "fail_service_func",
			req: &linkv1.SaveUrlRequest{
				Url: "http://google.com",
			},
			ctx: context.Background(),
			resp: &linkv1.SaveUrlResponse{
				ShortUrl: "",
			},
			isMock:  true,
			wantErr: true,
			respErr: status.Error(codes.Internal, "service error"),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			iLinkService := mocks.NewILinksService(t)

			if tc.isMock {
				iLinkService.
					On("SaveUrl", mock.Anything, tc.req.Url).
					Return(tc.resp.ShortUrl, tc.respErr)
			}

			s := NewLinksService(iLinkService)

			getUrl, err := s.SaveUrl(tc.ctx, tc.req)
			if (err != nil) != tc.wantErr {
				t.Errorf("SaveUrl() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr && len(getUrl.ShortUrl) != 10 {
				t.Errorf("SaveUrl() getUrl(%s) length = %v, want 10", getUrl, len(getUrl.ShortUrl))
				return
			}
		})
	}

}

func TestTransport_GetUrl(t *testing.T) {
	cases := []struct {
		name    string
		req     *linkv1.GetOriginRequest
		resp    *linkv1.GetOriginResponse
		ctx     context.Context
		isMock  bool
		wantErr bool
		respErr error
	}{
		{
			name: "success",
			req: &linkv1.GetOriginRequest{
				ShortUrl: "1234567890",
			},
			resp: &linkv1.GetOriginResponse{
				Url: "http://google.com",
			},
			ctx:     context.Background(),
			isMock:  true,
			wantErr: false,
			respErr: nil,
		},
		{
			name: "empty_short_url",
			req: &linkv1.GetOriginRequest{
				ShortUrl: "",
			},
			resp: &linkv1.GetOriginResponse{
				Url: "",
			},
			isMock:  false,
			wantErr: true,
			respErr: status.Error(codes.InvalidArgument, "empty short url"),
		},
		{
			name: "fail_service_func",
			req: &linkv1.GetOriginRequest{
				ShortUrl: "1234567890",
			},
			ctx: context.Background(),
			resp: &linkv1.GetOriginResponse{
				Url: "",
			},
			isMock:  true,
			wantErr: true,
			respErr: status.Error(codes.Internal, "service error"),
		},
		{
			name: "fail validating short url",
			req: &linkv1.GetOriginRequest{
				ShortUrl: "123",
			},
			ctx: context.Background(),
			resp: &linkv1.GetOriginResponse{
				Url: "http://google.com",
			},
			isMock:  false,
			wantErr: true,
			respErr: status.Error(codes.InvalidArgument, MsgInvalidShortUrl),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			iLinkService := mocks.NewILinksService(t)

			if tc.isMock {
				iLinkService.
					On("GetOrigin", mock.Anything, tc.req.ShortUrl).
					Return(tc.resp.Url, tc.respErr)
			}

			s := NewLinksService(iLinkService)

			origin, err := s.GetOrigin(tc.ctx, tc.req)
			if (err != nil) != tc.wantErr {
				t.Errorf("SaveUrl() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if !tc.wantErr && len(origin.Url) == 0 {
				t.Log(len(origin.Url) > 0, !tc.wantErr)
				t.Errorf("GetOrigin(%s) length = %v, want not nil", origin, len(origin.Url))
				return
			}
		})
	}
}

func TestServer_StartStop(t *testing.T) {
	ctx := context.Background()

	shortenerLogger := logger.New(logger.ServiceName)
	ctx = context.WithValue(ctx, logger.LoggerKey, shortenerLogger)

	// Создаем временный gRPC сервер
	grpcPort := 50052
	restPort := 8081
	mockService := &mocks.ILinksService{} // Мок сервиса
	server, err := NewServer(ctx, grpcPort, restPort, mockService)
	require.NoError(t, err, "Failed to create server")

	// Канал для отслеживания ошибок
	errChan := make(chan error, 1)

	// WaitGroup для ожидания завершения сервера
	var wg sync.WaitGroup
	wg.Add(1)

	// Запускаем серверы в отдельной горутине
	go func() {
		defer wg.Done()
		errChan <- server.Start(ctx)
	}()

	// Даем серверам время на запуск
	time.Sleep(1 * time.Second)

	// Останавливаем серверы
	err = server.Stop(ctx)
	assert.NoError(t, err, "Server should stop without errors")

	// Ждем завершения сервера
	wg.Wait()

	// Проверяем ошибку
	select {
	case err := <-errChan:
		// Игнорируем ошибку "http: Server closed", так как это ожидаемо при остановке сервера
		if err != nil && err != http.ErrServerClosed {
			assert.NoError(t, err, "Server should stop without errors")
		}
	case <-ctx.Done():
		t.Fatal("Test timed out")
	}
}
