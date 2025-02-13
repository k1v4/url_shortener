package grpc

import (
	"context"
	"github.com/k1v4/url_shortener/internal/transport/grpc/mocks"
	linkv1 "github.com/k1v4/url_shortener/pkg/api/link"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestTransport_SaveUrl(t *testing.T) {
	cases := []struct {
		name     string
		req      *linkv1.SaveUrlRequest
		ctx      context.Context
		resp     *linkv1.SaveUrlResponse
		shortUrl string
		wantErr  bool
		respErr  error
		mockErr  error
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
			wantErr: false,
			respErr: nil,
			mockErr: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			iLinkService := mocks.NewILinksService(t)

			iLinkService.
				On("SaveUrl", mock.Anything, tc.req.Url).
				Return(tc.resp.ShortUrl, tc.respErr)

			s := NewLinksService(iLinkService)

			getUrl, err := s.SaveUrl(tc.ctx, tc.req)
			if (err != nil) != tc.wantErr {
				t.Errorf("SaveUrl() error = %v, wantErr %v", err, tc.wantErr)
				return
			}

			if len(getUrl.ShortUrl) != 10 {
				t.Errorf("SaveUrl() getUrl(%s) length = %v, want 10", getUrl, len(getUrl.ShortUrl))
				return
			}
		})
	}

}

//func TestSaveUrl(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockService := mock_grpc.NewMockILinksService(ctrl)
//	server := NewLinksService(mockService)
//
//	tests := []struct {
//		name               string
//		request            *linkv1.SaveUrlRequest
//		expectedRequest    *linkv1.SaveUrlRequest
//		expectedResponse   *linkv1.SaveUrlResponse
//		mockBehavior       func(mock *mock_grpc.MockILinksService)
//		expectedStatusCode codes.Code
//		expectedError      error
//	}{
//		{
//			name:             "Success",
//			request:          &linkv1.SaveUrlRequest{Url: "https://example.com"},
//			expectedRequest:  &linkv1.SaveUrlRequest{Url: "https://example.com"},
//			expectedResponse: &linkv1.SaveUrlResponse{ShortUrl: "abc123AC_F"},
//			mockBehavior: func(mock *mock_grpc.MockILinksService) {
//				mock.EXPECT().
//					SaveUrl(gomock.Any(), "https://example.com").
//					Return("abc123AC_F", nil).
//					Times(1)
//			},
//			expectedStatusCode: codes.OK,
//			expectedError:      nil,
//		},
//		{
//			name:             "Empty URL",
//			request:          &linkv1.SaveUrlRequest{Url: ""},
//			expectedRequest:  &linkv1.SaveUrlRequest{Url: ""},
//			expectedResponse: nil,
//			mockBehavior: func(mock *mock_grpc.MockILinksService) {
//				// Мок не вызывается, так как запрос невалидный
//			},
//			expectedStatusCode: codes.InvalidArgument,
//			expectedError:      status.Error(codes.InvalidArgument, "empty url"),
//		},
//		{
//			name:             "Service Error",
//			request:          &linkv1.SaveUrlRequest{Url: "https://example.com"},
//			expectedRequest:  &linkv1.SaveUrlRequest{Url: "https://example.com"},
//			expectedResponse: nil,
//			mockBehavior: func(mock *mock_grpc.MockILinksService) {
//				mock.EXPECT().
//					SaveUrl(gomock.Any(), "https://example.com").
//					Return("", errors.New("internal error")).
//					Times(1)
//			},
//			expectedStatusCode: codes.Internal,
//			expectedError:      status.Error(codes.Internal, "internal error"),
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			// Настраиваем мок
//			tt.mockBehavior(mockService)
//
//			// Вызываем метод
//			resp, err := server.SaveUrl(context.Background(), tt.request)
//
//			// Проверяем ошибку
//			if tt.expectedError != nil {
//				assert.Error(t, err)
//				assert.Equal(t, tt.expectedStatusCode, status.Code(err))
//			} else {
//				assert.NoError(t, err)
//				assert.Equal(t, tt.expectedResponse, resp)
//			}
//		})
//	}
//}
//
//func TestGetOrigin(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockService := mock_grpc.NewMockILinksService(ctrl)
//	server := NewLinksService(mockService)
//
//	tests := []struct {
//		name               string
//		request            *linkv1.GetOriginRequest
//		expectedRequest    *linkv1.GetOriginRequest
//		expectedResponse   *linkv1.GetOriginResponse
//		mockBehavior       func(mock *mock_grpc.MockILinksService)
//		expectedStatusCode codes.Code
//		expectedError      error
//	}{
//		{
//			name:             "Success",
//			request:          &linkv1.GetOriginRequest{ShortUrl: "abc123_asd"},
//			expectedRequest:  &linkv1.GetOriginRequest{ShortUrl: "abc123_asd"},
//			expectedResponse: &linkv1.GetOriginResponse{Url: "https://example.com"},
//			mockBehavior: func(mock *mock_grpc.MockILinksService) {
//				mock.EXPECT().
//					GetOrigin(gomock.Any(), "abc123_asd").
//					Return("https://example.com", nil).
//					Times(1)
//			},
//			expectedStatusCode: codes.OK,
//			expectedError:      nil,
//		},
//		{
//			name:             "Empty Short URL",
//			request:          &linkv1.GetOriginRequest{ShortUrl: ""},
//			expectedRequest:  &linkv1.GetOriginRequest{ShortUrl: ""},
//			expectedResponse: nil,
//			mockBehavior: func(mock *mock_grpc.MockILinksService) {
//				// Мок не вызывается, так как запрос невалидный
//			},
//			expectedStatusCode: codes.InvalidArgument,
//			expectedError:      status.Error(codes.InvalidArgument, "empty short url"),
//		},
//		{
//			name:             "Service Error",
//			request:          &linkv1.GetOriginRequest{ShortUrl: "abc123"},
//			expectedRequest:  &linkv1.GetOriginRequest{ShortUrl: "abc123"},
//			expectedResponse: nil,
//			mockBehavior: func(mock *mock_grpc.MockILinksService) {
//				mock.EXPECT().
//					GetOrigin(gomock.Any(), "abc123").
//					Return("", errors.New("internal error")).
//					Times(1)
//			},
//			expectedStatusCode: codes.Internal,
//			expectedError:      status.Error(codes.Internal, "internal error"),
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			// Настраиваем мок
//			tt.mockBehavior(mockService)
//
//			// Вызываем метод
//			resp, err := server.GetOrigin(context.Background(), tt.request)
//
//			// Проверяем ошибку
//			if tt.expectedError != nil {
//				assert.Error(t, err)
//				assert.Equal(t, tt.expectedStatusCode, status.Code(err))
//			} else {
//				assert.NoError(t, err)
//				assert.Equal(t, tt.expectedResponse, resp)
//			}
//		})
//	}
//}
