package main

import (
	"log/slog"
	"net"
	"os"

	"github.com/2026-snapy/recommend-server/internal/config"
	"github.com/2026-snapy/recommend-server/internal/db"
	"github.com/2026-snapy/recommend-server/internal/friend"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
        slog.Error("failed to listen", "error", err)
		os.Exit(1)
    }

	db, err := db.NewDB(cfg)
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		os.Exit(1)
	}

	friendRepo := friend.NewFriendRepository(db)
	
	s := grpc.NewServer()
	
	slog.Info("Server listening on " + cfg.Addr)
	if err := s.Serve(lis); err != nil {
		slog.Error("failed to serve", "error", err)
		os.Exit(1)
    }
}