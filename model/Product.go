package model

import (
	"database/sql"
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
