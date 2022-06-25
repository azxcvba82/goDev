package model

import (
	"database/sql"
	"errors"
	"main/util"
	"strconv"
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
	queryString := `SELECT P.*, A.fAlbumName, A.fCoverPath  FROM tPlayLists L
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
		var fCoverPath string
		err = rows.Scan(&fProductID, &fAlbumID, &fProductName, nullfSinger, &fSIPrice, nullfComposer, &fFilePath, &fPlayStart, &fPlayEnd, &fAlbumName, &fCoverPath)


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
		obj.CoverPath = fCoverPath

		playList = append(playList, obj)
		counter++
	}
	return playList, nil
}

func UserAddPlayLists(sqlConnectionString string, account string, productId string) (result int64, err error) {
	prod, err := GetProductById(util.GetSQLConnectStringRead(), productId)

	if err != nil {
		outputErr := errors.New("product not found")
		return -1, outputErr
	}

	prods, err := GetPlayListByAccount(util.GetSQLConnectStringRead(), account)

	_ = prod

	num, _ := strconv.Atoi(productId)
	for i := 0; i < len(prods); i++ {
		if prods[i].ProductID == num {
			return -1, errors.New("product exist")
		}
	}

	queryString := `INSERT INTO tPlayLists (fAccount, fProductID) VALUES (?,?)`
	row, err := util.SQLExec(sqlConnectionString, queryString, account, productId)
	if err != nil {
		return -1, err
	}

	return row, nil
}

func UserDeletePlayLists(sqlConnectionString string, account string, productId string) (result int64, err error) {
	prods, err := GetPlayListByAccount(util.GetSQLConnectStringRead(), account)

	exist := false

	num, _ := strconv.Atoi(productId)
	for i := 0; i < len(prods); i++ {
		if prods[i].ProductID == num {
			exist = true
			break
		}
	}

	if exist == true {
		queryString := `DELETE FROM tPlayLists WHERE fAccount = ? AND fProductID = ? `
		row, err := util.SQLExec(sqlConnectionString, queryString, account, productId)
		if err != nil {
			return -1, err
		}
		return row, nil
	} else {
		return -1, errors.New("product not exist in account")
	}

}
