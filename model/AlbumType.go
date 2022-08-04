package model

import (
	"errors"
	"main/util"
)

type AlbumType struct {
	TypeID   int    `db:"fTypeID"`
	TypeName string `db:"fTypeName"`
}

func AllAlbumType(sqlConnectionString string) (model []AlbumType, err error) {
	var albumTypes []AlbumType
	queryString := `SELECT fTypeID, fTypeName FROM tAlbumType  `
	err = util.SQLQueryV2(&albumTypes, sqlConnectionString, true, queryString)

	if err != nil {
		return albumTypes, err
	}

	if albumTypes == nil {
		return albumTypes, errors.New("select id not found")
	}
	return albumTypes, nil
}
