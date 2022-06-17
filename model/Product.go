package model

import (
	"database/sql"
	"errors"
	"main/util"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ProductID   int     `int:"productID"`
	ProductName string  `string:"productName"`
	AlbumID     int     `int:"albumID"`
	Singer      string  `string:"singer"`
	SIPrice     float32 `float32:"siPrice"`
	Composer    string  `string:"composer"`
	FilePath    string  `string:"filePath"`
	PlayStart   float32 `float32:"playStart"`
	PlayEnd     float32 `float32:"playEnd"`
}

type ProductSearch struct {
	ProductID   int     `int:"productID"`
	ProductName string  `string:"productName"`
	AlbumID     int     `int:"albumID"`
	Singer      string  `string:"singer"`
	SIPrice     float32 `float32:"siPrice"`
	Composer    string  `string:"composer"`
	FilePath    string  `string:"filePath"`
	PlayStart   float32 `float32:"playStart"`
	PlayEnd     float32 `float32:"playEnd"`
	AlbumName   string  `string:"AlbumName"`
	CoverPath   string  `string:"CoverPath"`
	ActivityID  int     `int:"ActivityID"`
	Discount    float32 `float32:"Discount"`
}

type SearchOption struct {
	ProductName string `string:"productName"`
	AlbumName   string `string:"AlbumName"`
	Singer      string `string:"Singer"`
	Group       string `string:"Group"`
	Composer    string `string:"Composer"`
	Type        string `string:"Type"`
}

func GetProductsByAlbumId(sqlConnectionString string, albumId string) (model []Product, err error) {
	var products []Product
	nullfSinger := new(sql.NullString)
	nullfComposer := new(sql.NullString)
	queryString := `SELECT * FROM tProducts where fAlbumID = ? `
	rows, err := util.SQLQuery(sqlConnectionString, queryString, albumId)

	if err != nil {
		return products, err
	}

	counter := 0
	for rows.Next() {
		var obj Product
		var fProductID int
		var fProductName string
		var fAlbumID int
		var fSIPrice float32
		var fFilePath string
		var fPlayStart float32
		var fPlayEnd float32
		err = rows.Scan(&fProductID, &fAlbumID, &fProductName, nullfSinger, &fSIPrice, nullfComposer, &fFilePath, &fPlayStart, &fPlayEnd)

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
		products = append(products, obj)
		counter++
	}
	if counter == 0 {
		return products, errors.New("select id not found")
	}
	return products, nil
}

func GetProductsByProductName(sqlConnectionString string, c echo.Context) (model []ProductSearch, err error) {
	name := c.QueryParam("name")
	albumName := c.QueryParam("albumName")
	singer := c.QueryParam("singer")
	group := c.QueryParam("group")
	composer := c.QueryParam("composer")
	albumType := c.QueryParam("type")
	var option SearchOption
	option.ProductName = name
	option.AlbumName = albumName
	option.Singer = singer
	option.Group = group
	option.Composer = composer
	option.Type = albumType

	sb := ""
	var querySlice []any
	if option.ProductName != "" {
		sb += " fProductName LIKE CONCAT('%',?,'%') "
		querySlice = append(querySlice, option.ProductName)
	}

	if option.AlbumName != "" {
		if len(sb) == 0 {
			sb += ` A.fAlbumName LIKE CONCAT('%',?,'%')`
		} else {
			sb += ` AND A.fAlbumName LIKE CONCAT('%',?,'%')`
		}
		querySlice = append(querySlice, option.AlbumName)
	}

	var products []ProductSearch
	nullfSinger := new(sql.NullString)
	nullfComposer := new(sql.NullString)
	nullfActivityID := new(sql.NullInt32)
	queryString := `SELECT P.*, A.fAlbumName, A.fCoverPath, A.fActivityID, A.fDiscount FROM tProducts P
	INNER JOIN godev.tAlbum A ON P.fAlbumID = A.fAlbumID WHERE ` + sb
	rows, err := util.SQLQuery(sqlConnectionString, queryString, querySlice...)

	if err != nil {
		return products, err
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
		var fDiscount float32
		err = rows.Scan(&fProductID, &fAlbumID, &fProductName, nullfSinger, &fSIPrice, nullfComposer, &fFilePath, &fPlayStart, &fPlayEnd, &fAlbumName, &fCoverPath, nullfActivityID, &fDiscount)

		util.CheckErr(err)
		if nullfSinger.Valid {
			obj.Singer = string(nullfSinger.String)
		}
		if nullfComposer.Valid {
			obj.Composer = string(nullfComposer.String)
		}
		if nullfActivityID.Valid {
			obj.ActivityID = int(nullfActivityID.Int32)
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
		obj.Discount = fDiscount
		products = append(products, obj)
		counter++
	}
	if counter == 0 {
		return products, errors.New("select id not found")
	}
	return products, nil
}
