package main

import (
	errors "GO_Redis/error"
	redis "GO_Redis/redis"
	"GO_Redis/user"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	u := user.User{
		Name:    "oleh",
		Surname: "shalapsky",
		Email:   "shalapsky.oleg@gmail.com",
		Age:     20,
	}

	redis.Post(u)
	parse, err := uuid.Parse("edae679b-196e-4815-bc97-fe5008d6140f")
	errors.CheckError(err)
	fmt.Println(*redis.GetByUUID(parse))
}
