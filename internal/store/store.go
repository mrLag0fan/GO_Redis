package store

import (
	"GO_Redis/internal/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

type Store struct {
	DB          *sql.DB
	RedisClient *redis.Client
}

func NewStore(cfg config.Config) (*Store, error) {
	var err error
	store := Store{}
	store.RedisClient = initRedis(cfg)
	store.DB, err = initPostgres(cfg)
	if err != nil {
		return nil, err
	}
	return &store, err
}

func initRedis(cfg config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
}

func initPostgres(cfg config.Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Dbname)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("The database is connected")
	return db, err
}
