package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"wallet/internal/app"

	"go.uber.org/zap"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	bootstrap, err := app.Bootstrap(ctx)
	if err != nil {
		panic(err)
	}
	defer bootstrap.Log.Sync()

	startHTTPServer(ctx, bootstrap)
}

func startHTTPServer(ctx context.Context, bootstrap *app.BootstrapResult) {
	addr := fmt.Sprintf(":%d", bootstrap.Cfg.HTTP.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      bootstrap.Engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	bootstrap.Log.Info("starting HTTP server", zap.String("addr", addr))
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			bootstrap.Log.Fatal("listen", zap.Error(err))
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bootstrap.Log.Info("shutting down server")
	if err := srv.Shutdown(ctx); err != nil {
		bootstrap.Log.Error("server shutdown failed", zap.Error(err))
	}
}
