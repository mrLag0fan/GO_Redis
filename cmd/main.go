package main

import (
	"GO_Redis/internal/config"
	"GO_Redis/internal/entity"
	redis "GO_Redis/internal/redis"
	errors "GO_Redis/pkg/error"
	"GO_Redis/pkg/store"
	"fmt"
	"log"

	"github.com/google/uuid"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config, err := config.NewFromEnv()
	if err != nil {
		log.Fatalf("Can`t read congif from ENV: %v", err)
	}

	_, err = store.NewStore(config.Postgres)
	if err != nil {
		log.Fatalf("Can`t connect to PostgreSQL: %v", err)
	}

	u := entity.User{
		Name:    "oleh",
		Surname: "shalapsky",
		Email:   "shalapsky.oleg@gmail.com",
		Age:     20,
	}

	redis.Post(u)
	parse, err := uuid.Parse("a8ca47ac-e966-42bc-a8f2-3c158aa9a5d4")
	errors.CheckError(err)
	fmt.Println(*redis.GetByUUID(parse))
}
