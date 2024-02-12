package postgres

import (
	my_user "GO_Redis/internal/entity"
	"GO_Redis/pkg/database"
	errors "GO_Redis/pkg/error"
	"github.com/google/uuid"
)

func GetByUUID(id uuid.UUID) *my_user.User {
	row := database.DB.QueryRow(`SELECT * FROM "go-user" WHERE id = $1`, id)

	result := &my_user.User{}
	err := row.Scan(&result.Name, &result.Surname, &result.Email)
	errors.CheckError(err)
	return result
}

func Post(userID uuid.UUID, u my_user.User) {
	_, err := database.DB.Exec(`INSERT INTO "go-user" VALUES ($1, $2, $3, $4)`,
		userID,
		u.Name,
		u.Surname,
		u.Email)
	errors.CheckError(err)
}
