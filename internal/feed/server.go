package feed

import (
	"context"

	feedpb "github.com/2026-snapy/proto/go/feed"
)

type server struct {
	feedpb.UnimplementedRecommendServiceServer
	svc *FeedService
}

func NewServer(svc *FeedService) *server {
	return &server{svc: svc}
}

func (s *server) RecommendFeed(ctx context.Context, req *feedpb.RecommendRequest) (*feedpb.RecommendResponse, error) {
	return s.svc.RecommendFeed(req)
}
