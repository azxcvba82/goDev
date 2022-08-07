package util

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

func SQLQuery(sqlConnectionString string, sqlCommand string, args ...any) (r *sql.Rows, err error) {

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

func SQLQueryV2(model interface{}, sqlConnectionString string, useCache bool, sqlCommand string, args ...any) (err error) {

	db, err := sqlx.Open("mysql", sqlConnectionString)
	defer db.Close()

	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(time.Minute * 3)

	if useCache == false {
		if strings.Contains(reflect.ValueOf(model).Type().String(), "[]") {
			db.Select(model, sqlCommand, args...)
		} else {
			db.Get(model, sqlCommand, args...)
		}
		return err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDISADDRESS"),
		Password: os.Getenv("REDISPASSWORD"),
		DB:       0,
	})

	// NO implement cancel require
	ctx := context.Background()

	var cacheAvailable bool = true
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("redis conn err:" + err.Error())
		cacheAvailable = false
	}

	var sqlCommandWithArgs = sqlCommand
	for _, arg := range args {
		sqlCommandWithArgs += arg.(string)
	}

	md5Inst := md5.New()
	md5Inst.Write([]byte(sqlCommandWithArgs))
	md5Result := hex.EncodeToString(md5Inst.Sum([]byte("")))
	//_, err = rdb.Del(ctx, md5Result).Result()

	var intExist int64 = 0
	if cacheAvailable == true {
		intExist, err = rdb.Exists(ctx, md5Result).Result()
	}
	//fmt.Println(intExist)

	// key exist
	if intExist == 1 {

		jsonString, err := rdb.Get(ctx, md5Result).Result()
		if err != nil {
			fmt.Println("redis get err:" + err.Error())
		}
		json.Unmarshal([]byte(jsonString), model)
		fmt.Println("Redis READ: " + sqlCommand)

		return err

	} else {

		if strings.Contains(reflect.ValueOf(model).Type().String(), "[]") {
			db.Select(model, sqlCommand, args...)
		} else {
			db.Get(model, sqlCommand, args...)
		}
		s, err := json.Marshal(model)
		//fmt.Println(string(s))

		if cacheAvailable == true {
			var cacheDuration time.Duration = 86400 * time.Second
			_, err = rdb.Set(ctx, md5Result, string(s), cacheDuration).Result()
			if err != nil {
				fmt.Println("redis set err:" + err.Error())
			}
		}

		return err
	}

}

func SQLExec(sqlConnectionString string, withTransaction bool, sqlCommand string, args ...any) (id int64, cnt int64, err error) {
	db, err := sql.Open("mysql", sqlConnectionString)
	defer db.Close()

	if err != nil {
		return -1, -1, err
	}
	db.SetConnMaxLifetime(time.Minute * 3)

	var execResult sql.Result
	var insertId int64
	var rowsAffected int64
	if withTransaction == false {
		execResult, err = db.Exec(sqlCommand, args...)
	} else {
		ctx := context.Background()
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return -1, -1, err
		}
		defer tx.Rollback()
		execResult, err = tx.ExecContext(ctx, sqlCommand, args...)
		if err != nil {
			return -1, -1, err
		}

		insertId, err := execResult.LastInsertId()
		if err != nil {
			return -1, -1, err
		}

		rowsAffected, err := execResult.RowsAffected()
		if err != nil {
			return -1, -1, err
		}

		if err = tx.Commit(); err != nil {
			return insertId, rowsAffected, err
		}
	}
	if err != nil {
		return -1, -1, err
	}

	insertId, err = execResult.LastInsertId()
	if err != nil {
		return -1, -1, err
	}
	rowsAffected, err = execResult.RowsAffected()
	return insertId, rowsAffected, err
}
