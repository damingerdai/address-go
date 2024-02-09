package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type ProdServer struct {
	server *http.Server
}

func NewPordServer(server *http.Server) *ProdServer {
	return &ProdServer{server}
}

func (ps *ProdServer) Run() error {
	ctx := context.Background()
	errorChan := make(chan error)
	go func() {
		if err := ps.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("listen error", "msg", err.Error())
			errorChan <- err
		}
	}()

	// https://github.com/hlandao/service/issues/10
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		slog.Info("Shutdown server...\n")
		childCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		if err := ps.server.Shutdown(childCtx); err != nil {
			slog.Error("Server shutdown", "msg", err.Error())
		}
		select {
		case <-childCtx.Done():
			slog.Error("5 second timeout")
		default:
			slog.Error("Server exiting")
		}
	case err := <-errorChan:
		slog.Error("fail to run server:", "msg", err.Error())
	}

	return nil
}
