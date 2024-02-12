package go_redis

import (
	"GO_Redis/internal/entity"
	"GO_Redis/internal/postgres"
	"GO_Redis/internal/util"
	"GO_Redis/pkg/database"
	error2 "GO_Redis/pkg/error"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func Post(u entity.User) {
	mappedUser := util.StructTOMap(u)
	ctx := context.Background()
	userID := uuid.New()
	err := database.RedisClient.HMSet(ctx, userID.String(), mappedUser).Err()
	error2.CheckError(err)
	postgres.Post(userID, u)
}

func GetByUUID(id uuid.UUID) *entity.User {
	ctx := context.Background()
	u, err := database.RedisClient.HGetAll(ctx, id.String()).Result()
	if errors.Is(err, redis.Nil) {
		postgres.GetByUUID(id)
	} else {
		error2.CheckError(err)
	}
	return util.MapToUser(u)
}
