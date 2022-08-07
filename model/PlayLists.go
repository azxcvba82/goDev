package model

import (
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
	queryString := `SELECT P.fProductID, P.fProductName, P.fAlbumID, P.fSinger, P.fSIPrice, P.fComposer, P.fFilePath, P.fPlayStart, P.fPlayEnd, 
									A.fAlbumName, A.fCoverPath  FROM tPlayLists L
									INNER JOIN tProducts P ON L.fProductID = P.fProductID
									INNER JOIN tAlbum A ON P.fAlbumID = A.fAlbumID
									WHERE L.fAccount = ? `
	err = util.SQLQueryV2(&playList, sqlConnectionString, false, queryString, account)

	if err != nil {
		return playList, err
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
