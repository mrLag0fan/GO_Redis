package main

import (
	"GO_Redis/internal/entity"
	redis "GO_Redis/internal/redis"
	errors "GO_Redis/pkg/error"
	"fmt"
	"github.com/google/uuid"
)

func main() {
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
