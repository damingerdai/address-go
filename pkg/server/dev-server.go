package server

import (
	"log/slog"
	"net/http"
)

type DevServer struct {
	server *http.Server
}

func NewDevServer(server *http.Server) *DevServer {
	return &DevServer{server: server}
}

func (ds *DevServer) Run() error {
	if err := ds.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("listen error", "msg", err.Error())
		return err
	}
	return nil
}
