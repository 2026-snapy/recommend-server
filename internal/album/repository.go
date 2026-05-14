package album

import (
	"time"

	"github.com/2026-snapy/recommend-server/internal/db"
	"github.com/jmoiron/sqlx"
)

type Cursor struct {
	PhotoCount int64 `db:"photo_count"`
	PublishedAt time.Time `db:"published_at"`
}

type DailyAlbumRepository struct {
	db *db.DB
}

func NewDailyAlbumRepository(db *db.DB) *DailyAlbumRepository {
	return &DailyAlbumRepository{db}
}

func (d *DailyAlbumRepository) CursorById(id int64) (Cursor, error) {
	var c Cursor
	err := d.db.Client.Get(
		&c,
		`SELECT photo_count, published_at
		FROM daily_albums WHERE id = ?`,
		id,
	)
	if err != nil {
		return Cursor{}, err
	}

	return c, nil
}

func (d *DailyAlbumRepository) IdsInUserIds(userIds []int64) ([]int64, error) {
	query, args, err:= sqlx.In(
		`SELECT id FROM daily_albums
		WHERE user_id IN (?)
		AND status = 'PUBLISHED'
		ORDER BY published_at DESC, photo_count DESC
		LIMIT 300`,
		userIds,
	)
	if err != nil {
		return []int64{}, err
	}

	query = d.db.Client.Rebind(query)

	var ids []int64
	err = d.db.Client.Select(&ids, query, args...)
	if err != nil {
		return []int64{}, err
	}

	return ids, nil
}

func (d *DailyAlbumRepository) IdsInUserIdsWithCursor(userIds []int64, cursorId int64, c Cursor) ([]int64, error) {
	query, args, err:= sqlx.In(
		`SELECT id FROM daily_albums
		WHERE user_id IN (?)
		AND status = 'PUBLISHED'
		AND (
		published_at < ?
		OR (published_at = ? AND photo_count < ?)
		OR (published_at = ? AND photo_count = ? AND id < ?)
		)
		ORDER BY published_at DESC, photo_count DESC
		LIMIT 300`,
		userIds,
		c.PublishedAt,
		c.PublishedAt,
		c.PhotoCount,
		c.PublishedAt,
		c.PhotoCount,
		cursorId,
	)
	if err != nil {
		return []int64{}, err
	}

	var ids []int64
	err = d.db.Client.Select(&ids, query, args...)
	if err != nil {
		return []int64{}, err
	}

	return ids, nil
}