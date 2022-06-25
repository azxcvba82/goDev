package model

import (
	"database/sql"
	"main/util"
)

type PlayList struct {
	PlayId    int    `int:"PlayId"`
	Account   string `string:"Account"`
	ProductId int    `int:"AlbumID"`
}

func GetPlayListByAccount(sqlConnectionString string, account string) (model []ProductSearch, err error) {
	playList := []ProductSearch{}
	nullfSinger := new(sql.NullString)
	nullfComposer := new(sql.NullString)
	queryString := `SELECT P.*, A.fAlbumName  FROM tPlayLists L
									INNER JOIN tProducts P ON L.fProductID = P.fProductID
									INNER JOIN tAlbum A ON P.fAlbumID = A.fAlbumID
									WHERE L.fAccount = ? `
	rows, err := util.SQLQuery(sqlConnectionString, queryString, account)

	if err != nil {
		return playList, err
	}

	counter := 0
	for rows.Next() {
		var obj ProductSearch
		var fProductID int
		var fProductName string
		var fAlbumID int
		var fSIPrice float32
		var fFilePath string
		var fPlayStart float32
		var fPlayEnd float32
		var fAlbumName string
		err = rows.Scan(&fProductID, &fAlbumID, &fProductName, nullfSinger, &fSIPrice, nullfComposer, &fFilePath, &fPlayStart, &fPlayEnd, &fAlbumName)

		util.CheckErr(err)
		if nullfSinger.Valid {
			obj.Singer = string(nullfSinger.String)
		}
		if nullfComposer.Valid {
			obj.Composer = string(nullfComposer.String)
		}
		obj.ProductID = fProductID
		obj.ProductName = fProductName
		obj.AlbumID = fAlbumID
		obj.SIPrice = fSIPrice
		obj.FilePath = fFilePath
		obj.PlayStart = fPlayStart
		obj.PlayEnd = fPlayEnd
		obj.AlbumName = fAlbumName
		playList = append(playList, obj)
		counter++
	}
	return playList, nil
}
