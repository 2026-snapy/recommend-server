package main

import (
	"log/slog"
	"net"
	"os"

	"github.com/2026-snapy/recommend-server/internal/config"
	"github.com/2026-snapy/recommend-server/internal/db"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	db, err := db.NewDB(cfg)
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		os.Exit(1)
	}

	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
        slog.Error("failed to listen", "error", err)
		os.Exit(1)
    }
	
	s := grpc.NewServer()
	
	slog.Info("Server listening on " + cfg.Addr)
	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve", "error", err)
		os.Exit(1)
    }
}