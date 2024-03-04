package postgres

import (
	my_user "GO_Redis/internal/model"
	"GO_Redis/internal/store"
	"database/sql"
	"github.com/google/uuid"
	"log"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(store store.Store) *PostgresRepo {
	return &PostgresRepo{db: store.DB}
}

func (repo *PostgresRepo) GetByUUID(id uuid.UUID) *my_user.User {
	row := repo.db.QueryRow(`SELECT * FROM "go-user" WHERE id = $1`, id)

	result := &my_user.User{}
	err := row.Scan(&result.Name, &result.Surname, &result.Email)
	if err != nil {
		log.Fatalf("Can`t select user with id %v from postgres database: %v", id, err)
	}
	return result
}

func (repo *PostgresRepo) Post(userID uuid.UUID, u my_user.User) {
	_, err := repo.db.Exec(`INSERT INTO "go-user" VALUES ($1, $2, $3, $4)`,
		userID,
		u.Name,
		u.Surname,
		u.Email)
	if err != nil {
		log.Fatalf("Can`t insert user into postgres database: %v", err)
	}
}
