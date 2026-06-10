package friend

import "github.com/2026-snapy/recommend-server/internal/db"

type FriendRepository struct {
	db *db.DB
}

func NewFriendRepository(db *db.DB) *FriendRepository {
	return &FriendRepository{db}
}

func (f *FriendRepository) IdsByUserId(userId int64) ([]int64, error) {
	var ids []int64
	err := f.db.Client.Select(
		&ids,
		`SELECT CASE WHEN user_a_id = ?
		THEN user_b_id ELSE user_a_id END
		FROM friends WHERE user_a_id = ? OR user_b_id = ?`,
		userId,
		userId,
		userId,
	)
	if err != nil {
		return []int64{}, err
	}

	return ids, nil
}