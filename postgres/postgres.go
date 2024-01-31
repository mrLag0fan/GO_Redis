package postgres

import (
	errors "GO_Redis/error"
	my_user "GO_Redis/user"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "159357"
	dbname   = "go-user-db"
)

var db *sql.DB

func init() {
	var err error

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("The database is connected")
}

func GetByUUID(id uuid.UUID) *my_user.User {
	row := db.QueryRow(`SELECT * FROM "go-user" WHERE id = $1`, id)

	result := &my_user.User{}
	err := row.Scan(&result.Name, &result.Surname, &result.Email)
	errors.CheckError(err)
	return result
}

func Post(userID uuid.UUID, u my_user.User) {
	_, err := db.Exec(`INSERT INTO "go-user" VALUES ($1, $2, $3, $4)`,
		userID,
		u.Name,
		u.Surname,
		u.Email)
	errors.CheckError(err)
}
