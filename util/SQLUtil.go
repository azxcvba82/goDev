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
