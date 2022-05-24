package model

import (
	"database/sql"
	"main/util"
)

type Album struct {
	AlbumID    int     `int:"AlbumID"`
	AlbumName  string  `string:"AlbumName"`
	Maker      string  `string:"Maker"`
	Account    string  `string:"Account"`
	Year       string  `string:"Year"`
	Type       int     `int:"Type"`
	Status     int     `int:"Status"`
	ALPrice    float32 `float32:"ALPrice"`
	CoverPath  string  `string:"CoverPath"`
	Kinds      string  `string:"Kinds"`
	Discount   float32 `string:"Discount"`
	ActivityID int     `int:"ActivityID"`
}

func AllAlbum(sqlConnectionString string) (model []Album, err error) {
	var albums []Album
	nullfActivityID := new(sql.NullInt32)
	queryString := `SELECT *  FROM tAlbum WHERE fStatus = 2 `
	rows, err := util.SQLQuery(sqlConnectionString, queryString)

	if err != nil {
		return albums, err
	}

	for rows.Next() {
		var obj Album
		var fAlbumID int
		var fAlbumName string
		var fMaker string
		var fAccount string
		var fYear string
		var fType int
		var fStatus int
		var fALPrice float32
		var fCoverPath string
		var fKinds string
		var fDiscount float32
		err = rows.Scan(&fAlbumID, &fAlbumName, &fMaker, &fAccount, &fYear, &fType, &fStatus, &fALPrice, &fCoverPath, &fKinds, &fDiscount, nullfActivityID)

		util.CheckErr(err)
		if nullfActivityID.Valid {
			obj.ActivityID = int(nullfActivityID.Int32)
		}
		obj.AlbumID = fAlbumID
		obj.AlbumName = fAlbumName
		obj.Maker = fMaker
		obj.Account = fAccount
		obj.Year = fYear
		obj.Type = fType
		obj.Status = fStatus
		obj.ALPrice = fALPrice
		obj.CoverPath = fCoverPath
		obj.Kinds = fKinds
		obj.Discount = fDiscount
		albums = append(albums, obj)
	}
	return albums, nil
}
