package util

import (
	"database/sql"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
	db  *sql.DB
)

func SQLQuery(sqlConnectionString string, sqlCommand string, args ...any) (r *sql.Rows, err error) {

	if rdb == nil {
		rdb = redis.NewClient(&redis.Options{
			Addr:     os.Getenv("REDISADDRESS"),
			Password: os.Getenv("REDISPASSWORD"),
			DB:       0,
		})
	}

	db, err := sql.Open("mysql", sqlConnectionString)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	rows, err := db.Query(sqlCommand, args...)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func SQLExec(sqlConnectionString string, sqlCommand string, args ...any) (cnt int64, err error) {
	db, err := sql.Open("mysql", sqlConnectionString)
	if err != nil {
		return -1, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	execResult, err := db.Exec(sqlCommand, args...)
	if err != nil {
		return -1, err
	}
	return execResult.RowsAffected()
}
