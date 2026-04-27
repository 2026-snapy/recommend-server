package feed

import (
	feedpb "github.com/2026-snapy/proto/go/feed"
	"github.com/2026-snapy/recommend-server/internal/album"
	"github.com/2026-snapy/recommend-server/internal/friend"
)

type FeedService struct {
	friendRepo     *friend.FriendRepository
	dailyAlbumRepo *album.DailyAlbumRepository
}

func NewFeedService(friendRepo *friend.FriendRepository, dailyAlbumRepo *album.DailyAlbumRepository) *FeedService {
	return &FeedService{
		friendRepo:     friendRepo,
		dailyAlbumRepo: dailyAlbumRepo,
	}
}

func (f *FeedService) RecommendFeed(req *feedpb.RecommendRequest) (*feedpb.RecommendResponse, error) {
	friendIds, err := f.friendRepo.IdsByUserId(req.GetUserId())
	if err != nil {
		return nil, err
	}

	if len(friendIds) == 0 {
		return &feedpb.RecommendResponse{
			AlbumIds: []int64{},
		}, nil
	}

	size := int(req.GetSize())

	var ids []int64
	if req.Cursor == nil {
		ids, err = f.dailyAlbumRepo.IdsInUserIds(friendIds)
	} else {
		cursorId := *req.Cursor
		c, cErr := f.dailyAlbumRepo.CursorById(cursorId)
		if cErr != nil {
			return nil, cErr
		}
		ids, err = f.dailyAlbumRepo.IdsInUserIdsWithCursor(friendIds, cursorId, c)
	}
	if err != nil {
		return nil, err
	}

	hasNext := len(ids) > size
	if hasNext {
		ids = ids[:size]
	}

	resp := &feedpb.RecommendResponse{
		AlbumIds: ids,
		HasNext:  hasNext,
	}
	if hasNext && len(ids) > 0 {
		next := ids[len(ids)-1]
		resp.NextCursor = &next
	}

	return resp, nil
}
