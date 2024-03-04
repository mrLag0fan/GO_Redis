package main

import (
	"GO_Redis/internal/config"
	"GO_Redis/internal/model"
	"GO_Redis/internal/postgres"
	redis "GO_Redis/internal/redis"
	"GO_Redis/internal/store"
	"fmt"
	"github.com/google/uuid"
	"log"
)

func main() {
	cfg, err := config.NewFromEnv()
	if err != nil {
		log.Fatalf("Can`t read congif from ENV: %v", err)
	}
	newStore, err := store.NewStore(*cfg)
	if err != nil {
		log.Fatalf("Can`t create store: %v", err)
	}
	u := model.User{
		Name:    "oleh",
		Surname: "shalapsky",
		Email:   "shalapsky.oleg@gmail.com",
		Age:     20,
	}
	postgresRepo := postgres.NewPostgresRepo(*newStore)
	redisRepo := redis.NewRedisRepo(*newStore, *postgresRepo)

	redisRepo.Post(u)
	parse, err := uuid.Parse("a8ca47ac-e966-42bc-a8f2-3c158aa9a5d4")
	if err != nil {
		log.Fatalf("Can`t create uuid: %v", err)
	}
	fmt.Println(*redisRepo.GetByUUID(parse))
}
