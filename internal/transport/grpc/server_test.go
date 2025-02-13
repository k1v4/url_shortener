package grpc

//func TestServer_StartStop(t *testing.T) {
//	ctx := context.Background()
//
//	shortenerLogger := logger.New(logger.ServiceName)
//	ctx = context.WithValue(ctx, logger.LoggerKey, shortenerLogger)
//
//	// Создаем временный gRPC сервер
//	grpcPort := 50052
//	restPort := 8081
//	mockService := &mock_grpc.MockILinksService{} // Мок сервиса
//	server, err := NewServer(ctx, grpcPort, restPort, mockService)
//	require.NoError(t, err, "Failed to create server")
//
//	// Канал для отслеживания ошибок
//	errChan := make(chan error, 1)
//
//	// Запускаем серверы в отдельной горутине
//	go func() {
//		errChan <- server.Start(ctx)
//	}()
//
//	// Даем серверам время на запуск
//	time.Sleep(2 * time.Second)
//
//	// Останавливаем серверы
//	err = server.Stop(ctx)
//	assert.NoError(t, err, "Server should stop without errors")
//
//	// Проверяем, что серверы завершились без ошибок
//	select {
//	case err := <-errChan:
//		if err != nil && !errors.Is(err, http.ErrServerClosed) {
//			assert.NoError(t, err, "Server should stop without errors")
//		}
//	case <-ctx.Done():
//		t.Fatal("Test timed out")
//	}
//}
