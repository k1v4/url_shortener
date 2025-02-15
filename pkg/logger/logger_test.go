package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func TestLogger_Info(t *testing.T) {
	testLogger := zaptest.NewLogger(t)
	defer testLogger.Sync()

	l := &logger{
		serviceName: "test_service",
		logger:      testLogger,
	}

	ctx := context.WithValue(context.Background(), RequestID, "12345")

	l.Info(ctx, "test info message", zap.String("key", "value"))
}

func TestLogger_Error(t *testing.T) {
	testLogger := zaptest.NewLogger(t)
	defer testLogger.Sync()

	l := &logger{
		serviceName: "test_service",
		logger:      testLogger,
	}

	ctx := context.WithValue(context.Background(), RequestID, "12345")

	l.Error(ctx, "test error message", zap.String("key", "value"))
}

func TestLogger_Info_WithoutRequestID(t *testing.T) {
	testLogger := zaptest.NewLogger(t)
	defer testLogger.Sync()

	l := &logger{
		serviceName: "test_service",
		logger:      testLogger,
	}

	ctx := context.Background()

	l.Info(ctx, "test info message", zap.String("key", "value"))
}

func TestLogger_Error_WithoutRequestID(t *testing.T) {
	testLogger := zaptest.NewLogger(t)
	defer testLogger.Sync()

	l := &logger{
		serviceName: "test_service",
		logger:      testLogger,
	}

	ctx := context.Background()

	l.Error(ctx, "test error message", zap.String("key", "value"))
}

func TestNew(t *testing.T) {
	l := New("test_service")

	assert.NotNil(t, l)
}

func TestGetLoggerFromCtx(t *testing.T) {
	testLogger := zaptest.NewLogger(t)
	defer testLogger.Sync()

	l := &logger{
		serviceName: "test_service",
		logger:      testLogger,
	}

	ctx := context.WithValue(context.Background(), LoggerKey, l)

	retrievedLogger := GetLoggerFromCtx(ctx)

	assert.Equal(t, l, retrievedLogger)
}
