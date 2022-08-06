package model

import (
	"errors"
	"main/util"
	"strconv"

	"github.com/labstack/echo/v4"
	"gopkg.in/guregu/null.v4"
)

type Product struct {
	ProductID   int         `db:"fProductID"`
	ProductName string      `db:"fProductName"`
	AlbumID     int         `db:"fAlbumID"`
	Singer      null.String `db:"fSinger"`
	SIPrice     float32     `db:"fSIPrice"`
	Composer    null.String `db:"fComposer"`
	FilePath    string      `db:"fFilePath"`
	PlayStart   float32     `db:"fPlayStart"`
	PlayEnd     float32     `db:"fPlayEnd"`
}

type ProductSearch struct {
	ProductID   int         `db:"fProductID"`
	ProductName string      `db:"fProductName"`
	AlbumID     int         `db:"fAlbumID"`
	Singer      null.String `db:"fSinger"`
	SIPrice     float32     `db:"fSIPrice"`
	Composer    null.String `db:"fComposer"`
	FilePath    string      `db:"fFilePath"`
	PlayStart   float32     `db:"fPlayStart"`
	PlayEnd     float32     `db:"fPlayEnd"`
	AlbumName   string      `db:"fAlbumName"`
	CoverPath   string      `db:"fCoverPath"`
	ActivityID  null.Int    `db:"fActivityID"`
	Discount    float32     `db:"fDiscount"`
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
	queryString := `SELECT fProductID, fProductName, fAlbumID, fSinger, fSIPrice, fComposer, fFilePath, fPlayStart, fPlayEnd FROM tProducts where fAlbumID = ? `
	err = util.SQLQueryV2(&products, sqlConnectionString, true, queryString, albumId)

	if err != nil {
		return products, err
	}

	if products == nil {
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

	if option.Singer != "" {
		if len(sb) == 0 {
			sb += ` fSinger LIKE CONCAT('%',?,'%')`
		} else {
			sb += ` AND fSinger LIKE CONCAT('%',?,'%')`
		}
		querySlice = append(querySlice, option.Singer)
	}

	if option.Group != "" {
		if len(sb) == 0 {
			sb += ` A.fMaker LIKE CONCAT('%',?,'%')`
		} else {
			sb += ` AND A.fMaker LIKE CONCAT('%',?,'%')`
		}
		querySlice = append(querySlice, option.Group)
	}

	if option.Composer != "" {
		if len(sb) == 0 {
			sb += ` fComposer LIKE CONCAT('%',?,'%')`
		} else {
			sb += ` AND fComposer LIKE CONCAT('%',?,'%')`
		}
		querySlice = append(querySlice, option.Composer)
	}

	if option.Type != "" {
		if len(sb) == 0 {
			sb += ` A.fType = ? `
		} else {
			sb += ` AND A.fType  = ? `
		}
		typeId, _ := strconv.Atoi(option.Type)
		querySlice = append(querySlice, typeId)
	}

	var products []ProductSearch
	queryString := `SELECT P.fProductID, P.fProductName, P.fAlbumID, P.fSinger, P.fSIPrice, P.fComposer, P.fFilePath, P.fPlayStart, P.fPlayEnd, 
	A.fAlbumName, A.fCoverPath, A.fActivityID, A.fDiscount FROM tProducts P
	INNER JOIN godev.tAlbum A ON P.fAlbumID = A.fAlbumID WHERE ` + sb
	err = util.SQLQueryV2(&products, sqlConnectionString, true, queryString, querySlice...)

	if err != nil {
		return products, err
	}

	if products == nil {
		return products, errors.New("select id not found")
	}
	return products, nil
}

func GetProductById(sqlConnectionString string, id string) (model Product, err error) {
	var obj Product
	queryString := `SELECT fProductID, fProductName, fAlbumID, fSinger, fSIPrice, fComposer, fFilePath, fPlayStart, fPlayEnd FROM tProducts where fProductID = ? `
	err = util.SQLQueryV2(&obj, sqlConnectionString, true, queryString, id)

	if err != nil {
		return obj, err
	}

	var objForCheck Product
	if obj == objForCheck {
		return obj, errors.New("select id not found")
	}
	return obj, nil
}
