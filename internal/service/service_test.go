package service

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	mock_service "github.com/k1v4/url_shortener/internal/service/mocks"
	linkv1 "github.com/k1v4/url_shortener/pkg/api/link"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestLinksService_SaveUrl(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockRepo := mock_service.NewMockILinksRepository(ctrl)
//	svc := NewLinksService(mockRepo)
//
//	ctx := context.Background()
//
//	testTable := []struct {
//		name             string
//		requestBody      map[string]string
//		expectedRequest  *linkv1.SaveUrlRequest
//		expectedResponse *linkv1.SaveUrlResponse
//		mockBehavior     func(mockRepo *mock_service.MockILinksRepository)
//	}{
//		{
//			name: "Success",
//			requestBody: map[string]string{
//				"url": "https://example.com",
//			},
//			expectedRequest: &linkv1.SaveUrlRequest{
//				Url: "https://example.com",
//			},
//			expectedResponse: &linkv1.SaveUrlResponse{
//				ShortUrl: "abc123",
//			},
//			mockBehavior: func(mockRepo *mock_service.MockILinksRepository) {
//				// Ожидаем, что GetOrigin вернет пустую строку (shortUrl уникален)
//				mockRepo.EXPECT().
//					GetOrigin(ctx, gomock.Any()). // Используем gomock.Any() для shortUrl
//					Return("", nil).
//					Times(1)
//
//				// Ожидаем, что SaveUrl вернет shortUrl
//				mockRepo.EXPECT().
//					SaveUrl(ctx, "https://example.com", gomock.Any()). // Используем gomock.Any() для shortUrl
//					DoAndReturn(func(ctx context.Context, url, shortUrl string) (string, error) {
//						return "abc123", nil // Возвращаем ожидаемый shortUrl
//					}).
//					Times(1)
//			},
//		},
//		{
//			name: "Short URL conflict",
//			requestBody: map[string]string{
//				"url": "https://example.com",
//			},
//			expectedRequest: &linkv1.SaveUrlRequest{
//				Url: "https://example.com",
//			},
//			expectedResponse: &linkv1.SaveUrlResponse{
//				ShortUrl: "",
//			},
//			mockBehavior: func(mockRepo *mock_service.MockILinksRepository) {
//				// Ожидаем, что GetOrigin всегда возвращает существующий shortUrl
//				mockRepo.EXPECT().
//					GetOrigin(ctx, gomock.Any()). // Используем gomock.Any() для shortUrl
//					Return("existingUrl", nil).
//					Times(retries)
//			},
//		},
//	}
//
//	for _, tt := range testTable {
//		t.Run(tt.name, func(t *testing.T) {
//			// Настраиваем моки
//			tt.mockBehavior(mockRepo)
//
//			// Вызываем метод
//			result, err := svc.SaveUrl(ctx, tt.expectedRequest.Url)
//
//			// Проверяем результат
//			if tt.expectedResponse.ShortUrl == "" {
//				assert.Error(t, err, "Expected an error")
//				assert.Contains(t, err.Error(), "short url already exists", "Error message should match")
//			} else {
//				assert.NoError(t, err, "Expected no error")
//				assert.Equal(t, tt.expectedResponse.ShortUrl, result, "Expected result to match")
//			}
//		})
//	}
//}
//
//func TestLinksService_GetOrigin(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockRepo := mock_service.NewMockILinksRepository(ctrl)
//	svc := NewLinksService(mockRepo)
//
//	ctx := context.Background()
//
//	testTable := []struct {
//		name             string
//		requestBody      map[string]string
//		expectedRequest  *linkv1.GetOriginRequest
//		expectedResponse *linkv1.GetOriginResponse
//		mockBehavior     mockBehaviorGet
//	}{
//		{
//			name: "Success",
//			requestBody: map[string]string{
//				"short_url": "abc123_890",
//			},
//			expectedRequest: &linkv1.GetOriginRequest{
//				ShortUrl: "abc123_890",
//			},
//			expectedResponse: &linkv1.GetOriginResponse{
//				Url: "https://example.com",
//			},
//			mockBehavior: func(mockRepo *mock_service.MockILinksRepository, req *linkv1.GetOriginRequest, expectedResponse *linkv1.GetOriginResponse) {
//				// Ожидаем, что GetOrigin вернет originUrl
//				mockRepo.EXPECT().
//					GetOrigin(ctx, req.ShortUrl).
//					Return(expectedResponse.Url, nil).
//					Times(1)
//			},
//		},
//		{
//			name: "Error",
//			requestBody: map[string]string{
//				"short_url": "abc123_890",
//			},
//			expectedRequest: &linkv1.GetOriginRequest{
//				ShortUrl: "abc123_890",
//			},
//			expectedResponse: &linkv1.GetOriginResponse{
//				Url: "",
//			},
//			mockBehavior: func(mockRepo *mock_service.MockILinksRepository, req *linkv1.GetOriginRequest, expectedResponse *linkv1.GetOriginResponse) {
//				// Ожидаем, что GetOrigin вернет ошибку
//				mockRepo.EXPECT().
//					GetOrigin(ctx, req.ShortUrl).
//					Return("", errors.New("database error")).
//					Times(1)
//			},
//		},
//	}
//
//	for _, tt := range testTable {
//		t.Run(tt.name, func(t *testing.T) {
//			// Настраиваем моки
//			tt.mockBehavior(mockRepo, tt.expectedRequest, tt.expectedResponse)
//
//			// Вызываем метод
//			result, err := svc.GetOrigin(ctx, tt.expectedRequest.ShortUrl)
//
//			// Проверяем результат
//			if tt.expectedResponse.Url == "" {
//				assert.Error(t, err, "Expected an error")
//				assert.Contains(t, err.Error(), "database error", "Error message should match")
//			} else {
//				assert.NoError(t, err, "Expected no error")
//				assert.Equal(t, tt.expectedResponse.Url, result, "Expected result to match")
//			}
//		})
//	}
//}
