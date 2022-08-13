package model

import (
	"errors"
	"main/util"
)

type Kind struct {
	KindID    int    `db:"fKindID"`
	KindName  string `db:"fKindName"`
	Color     string `db:"fColor"`
	PhotoPath string `db:"fPhotoPath"`
}

func AllKind(sqlConnectionString string) (model []Kind, err error) {
	var kinds []Kind
	queryString := `SELECT fKindID, fKindName, fColor, fPhotoPath  FROM tAlbumKind `
	err = util.SQLQueryV2(&kinds, sqlConnectionString, true, queryString)

	if err != nil {
		return kinds, err
	}

	if kinds == nil {
		return kinds, errors.New("select not found")
	}
	return kinds, nil
}
