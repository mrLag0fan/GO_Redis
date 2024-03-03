package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "159357"
	dbname   = "go-user-db"
)

var DB *sql.DB
var RedisClient *redis.Client

func init() {
	var wg sync.WaitGroup
	go initRedis(&wg)
	go initPostgres(&wg)
	wg.Add(2)
	wg.Wait()
}

func initRedis(wg *sync.WaitGroup) {
	defer wg.Done()
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func initPostgres(wg *sync.WaitGroup) {
	defer wg.Done()
	var err error
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("The database is connected")
}
