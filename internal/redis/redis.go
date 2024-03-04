package go_redis

import (
	"GO_Redis/internal/model"
	"GO_Redis/internal/postgres"
	"GO_Redis/internal/store"
	"GO_Redis/internal/util"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"log"
)

type RedisRepo struct {
	client   *redis.Client
	postgres postgres.PostgresRepo
}

func NewRedisRepo(store store.Store, repo postgres.PostgresRepo) *RedisRepo {
	return &RedisRepo{
		client:   store.RedisClient,
		postgres: repo,
	}
}

func (repo *RedisRepo) Post(u model.User) {
	mappedUser := util.StructTOMap(u)
	ctx := context.Background()
	userID := uuid.New()
	err := repo.client.HMSet(ctx, userID.String(), mappedUser).Err()
	if err != nil {
		log.Fatalf("Can`t insert user into redis database: %v", err)
	}
	repo.postgres.Post(userID, u)
}

func (repo *RedisRepo) GetByUUID(id uuid.UUID) *model.User {
	ctx := context.Background()
	u, err := repo.client.HGetAll(ctx, id.String()).Result()
	if errors.Is(err, redis.Nil) {
		repo.postgres.GetByUUID(id)
	}
	if err != nil {
		log.Fatalf("Can`t select user with id %v from redis database: %v", id, err)
	}
	return util.MapToUser(u)
}
