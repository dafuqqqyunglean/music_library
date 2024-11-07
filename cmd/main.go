package main

import (
	"context"
	"fmt"
	"github.com/dafuqqqyunglean/music_library/config"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	prvLogger, _ := zap.NewProduction()
	defer prvLogger.Sync()
	logger := prvLogger.Sugar()

	fmt.Println(logger.Level())

	prvCtx := context.Background()
	ctx, cancel := context.WithCancel(prvCtx)
	defer cancel()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Fatalf("error occured while reading config: %s", err.Error())
	}

	app := NewApp(ctx, logger, cfg)

	app.InitService()

	if err = app.Run(); err != nil {
		logger.Errorf(err.Error())
		return
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	<-sigChan

	if err = app.Shutdown(ctx); err != nil {
		logger.Errorf(err.Error())
		return
	}
}
