package model

import (
	"database/sql"
	"errors"
	"main/util"

	"gopkg.in/guregu/null.v4"
)

type Album struct {
	AlbumID    int      `db:"fAlbumID"`
	AlbumName  string   `db:"fAlbumName"`
	Maker      string   `db:"fMaker"`
	Account    string   `db:"fAccount"`
	Year       string   `db:"fYear"`
	Type       int      `db:"fType"`
	Status     int      `db:"fStatus"`
	ALPrice    float32  `db:"fALPrice"`
	CoverPath  string   `db:"fCoverPath"`
	Kinds      string   `db:"fKinds"`
	Discount   float32  `db:"fDiscount"`
	ActivityID null.Int `db:"fActivityID"`
}

func AllAlbum(sqlConnectionString string) (model []Album, err error) {
	var albums []Album
	//nullfActivityID := new(sql.NullInt32)
	queryString := `SELECT fAlbumID, fAlbumName, fMaker, fAccount, fYear, fType, fStatus, fALPrice, fCoverPath, fKinds, fDiscount, fActivityID  FROM tAlbum WHERE fStatus = 2 `
	//rows, err := util.SQLQuery(sqlConnectionString, queryString)
	err = util.SQLQueryV2(&albums, sqlConnectionString, queryString)

	if err != nil {
		return albums, err
	}

	// for rows.Next() {
	// 	var obj Album
	// 	var fAlbumID int
	// 	var fAlbumName string
	// 	var fMaker string
	// 	var fAccount string
	// 	var fYear string
	// 	var fType int
	// 	var fStatus int
	// 	var fALPrice float32
	// 	var fCoverPath string
	// 	var fKinds string
	// 	var fDiscount float32
	// 	err = rows.Scan(&fAlbumID, &fAlbumName, &fMaker, &fAccount, &fYear, &fType, &fStatus, &fALPrice, &fCoverPath, &fKinds, &fDiscount, nullfActivityID)

	// 	util.CheckErr(err)
	// 	if nullfActivityID.Valid {
	// 		obj.ActivityID = int(nullfActivityID.Int32)
	// 	}
	// 	obj.AlbumID = fAlbumID
	// 	obj.AlbumName = fAlbumName
	// 	obj.Maker = fMaker
	// 	obj.Account = fAccount
	// 	obj.Year = fYear
	// 	obj.Type = fType
	// 	obj.Status = fStatus
	// 	obj.ALPrice = fALPrice
	// 	obj.CoverPath = fCoverPath
	// 	obj.Kinds = fKinds
	// 	obj.Discount = fDiscount
	// 	albums = append(albums, obj)
	// }
	return albums, nil
}

func GetAlbumById(sqlConnectionString string, albumId string) (model Album, err error) {
	var album Album
	nullfActivityID := new(sql.NullInt32)
	queryString := `SELECT *  FROM tAlbum WHERE fAlbumID = ? LIMIT 1`
	rows, err := util.SQLQuery(sqlConnectionString, queryString, albumId)

	if err != nil {
		return album, err
	}

	counter := 0
	for rows.Next() {
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
			//album.ActivityID = int(nullfActivityID.Int32)
		}
		album.AlbumID = fAlbumID
		album.AlbumName = fAlbumName
		album.Maker = fMaker
		album.Account = fAccount
		album.Year = fYear
		album.Type = fType
		album.Status = fStatus
		album.ALPrice = fALPrice
		album.CoverPath = fCoverPath
		album.Kinds = fKinds
		album.Discount = fDiscount
		counter++
	}
	if counter == 0 {
		return album, errors.New("select id not found")
	}
	return album, nil
}

func GetAlbumsByKindId(sqlConnectionString string, kindId string) (model []Album, err error) {
	var albums []Album
	nullfActivityID := new(sql.NullInt32)
	queryString := `SELECT * FROM tAlbum where fStatus = 2 AND fkinds LIKE CONCAT('%',(SELECT fKindName FROM tAlbumKind WHERE fKindID = ?),'%') ORDER BY fYear DESC; `
	rows, err := util.SQLQuery(sqlConnectionString, queryString, kindId)

	if err != nil {
		return albums, err
	}

	counter := 0
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
			//obj.ActivityID = int(nullfActivityID.Int32)
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
		counter++
	}
	if counter == 0 {
		return albums, errors.New("select id not found")
	}
	return albums, nil
}
