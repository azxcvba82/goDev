package model

import (
	"errors"
	"main/util"
)

type AlbumType struct {
	TypeID   int    `int:"TypeID"`
	TypeName string `string:"TypeName"`
}

func AllAlbumType(sqlConnectionString string) (model []AlbumType, err error) {
	var albumTypes []AlbumType
	queryString := `SELECT * FROM tAlbumType  `
	rows, err := util.SQLQuery(sqlConnectionString, queryString)

	if err != nil {
		return albumTypes, err
	}

	counter := 0
	for rows.Next() {
		var obj AlbumType
		var fTypeID int
		var fTypeName string
		err = rows.Scan(&fTypeID, &fTypeName)

		util.CheckErr(err)
		obj.TypeID = fTypeID
		obj.TypeName = fTypeName
		albumTypes = append(albumTypes, obj)
		counter++
	}
	if counter == 0 {
		return albumTypes, errors.New("select id not found")
	}
	return albumTypes, nil
}
