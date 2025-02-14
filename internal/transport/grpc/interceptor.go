package grpc

import (
	"context"
	"github.com/k1v4/url_shortener/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// ContextWithLogger будем пробрасывать в запросы
func ContextWithLogger(l logger.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		l.Info(ctx, "request started", zap.String("method", info.FullMethod), zap.String("handler", info.FullMethod))
		return handler(ctx, req)
	}
}
