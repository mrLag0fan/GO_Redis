package go_redis

import (
	"GO_Redis/error"
	"GO_Redis/postgres"
	"GO_Redis/user"
	"GO_Redis/util"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Post(u user.User) {
	mappedUser := util.StructTOMap(u)
	ctx := context.Background()
	userID := uuid.New()
	err := redisClient.HMSet(ctx, userID.String(), mappedUser).Err()
	error.CheckError(err)
	postgres.Post(userID, u)
}

func GetByUUID(id uuid.UUID) *user.User {
	ctx := context.Background()
	u, err := redisClient.HGetAll(ctx, id.String()).Result()
	if errors.Is(err, redis.Nil) {
		postgres.GetByUUID(id)
	} else {
		error.CheckError(err)
	}
	return util.MapToUser(u)
}
