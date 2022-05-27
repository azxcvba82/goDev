package model

import (
	"main/util"
)

type Kind struct {
	KindID    int    `int:"KindID"`
	KindName  string `string:"KindName"`
	Color     string `string:"Color"`
	PhotoPath string `string:"PhotoPath"`
}

func AllKind(sqlConnectionString string) (model []Kind, err error) {
	var kinds []Kind
	queryString := `SELECT *  FROM tAlbumKind `
	rows, err := util.SQLQuery(sqlConnectionString, queryString)

	if err != nil {
		return kinds, err
	}

	for rows.Next() {
		var obj Kind
		var fKindID int
		var fKindName string
		var fColor string
		var fPhotoPath string
		err = rows.Scan(&fKindID, &fKindName, &fColor, &fPhotoPath)

		util.CheckErr(err)
		obj.KindID = fKindID
		obj.KindName = fKindName
		obj.Color = fColor
		obj.PhotoPath = fPhotoPath
		kinds = append(kinds, obj)
	}
	return kinds, nil
}
