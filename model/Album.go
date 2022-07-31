package model

import (
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
	queryString := `SELECT fAlbumID, fAlbumName, fMaker, fAccount, fYear, fType, fStatus, fALPrice, fCoverPath, fKinds, fDiscount, fActivityID  FROM tAlbum WHERE fStatus = 2 `
	err = util.SQLQueryV2(&albums, sqlConnectionString, true, queryString)

	if err != nil {
		return albums, err
	}

	return albums, nil
}

func GetAlbumById(sqlConnectionString string, albumId string) (model Album, err error) {
	var album Album
	queryString := `SELECT fAlbumID, fAlbumName, fMaker, fAccount, fYear, fType, fStatus, fALPrice, fCoverPath, fKinds, fDiscount, fActivityID  FROM tAlbum WHERE fAlbumID = ? LIMIT 1`
	err = util.SQLQueryV2(&album, sqlConnectionString, true, queryString, albumId)

	if err != nil {
		return album, err
	}

	var albumForCheck Album
	if album == albumForCheck {
		return album, errors.New("select id not found")
	}
	return album, nil
}

func GetAlbumsByKindId(sqlConnectionString string, kindId string) (model []Album, err error) {
	var albums []Album
	queryString := `SELECT fAlbumID, fAlbumName, fMaker, fAccount, fYear, fType, fStatus, fALPrice, fCoverPath, fKinds, fDiscount, fActivityID FROM tAlbum where fStatus = 2 AND fkinds LIKE CONCAT('%',(SELECT fKindName FROM tAlbumKind WHERE fKindID = ?),'%') ORDER BY fYear DESC; `
	err = util.SQLQueryV2(&albums, sqlConnectionString, true, queryString, kindId)

	if err != nil {
		return albums, err
	}

	if albums == nil {
		return albums, errors.New("select id not found")
	}
	return albums, nil
}
