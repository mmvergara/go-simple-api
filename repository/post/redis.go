package post

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mmvergara/go-simple-api/model"
	"github.com/redis/go-redis/v9"
)
type RedisRepo struct {
	Client *redis.Client
}

func postIDKey(id uint64) string {
	return fmt.Sprintf("post:%d", id)
}

func (r *RedisRepo) Insert(ctx context.Context, post model.Post) error {
	postJson, err := json.Marshal(post)
	if err != nil {
		return err
	}

	key := postIDKey(post.PostID)


	res := r.Client.SetNX(ctx, key, string(postJson), 0)

	if err:= res.Err(); err != nil {
		return fmt.Errorf("failed to set post %w", err)
	}

	return nil
}

func (r *RedisRepo) FindById(ctx context.Context, id uint64) (model.Post, error) {
	key := postIDKey(id)

}