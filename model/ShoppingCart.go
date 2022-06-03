package model

import (
	"database/sql"
	"errors"
	"main/util"
)

type PurchaseItem struct {
	PurchaseItemID int     `int:"PurchaseItemID"`
	Customer       string  `string:"Customer"`
	ProductID      int     `int:"ProductID"`
	Date           string  `string:"Date"`
	Price          float32 `string:"Price"`
	Quanity        int     `int:"Quanity"`
	IsAlbum        int     `int:"IsAlbum"`
	Discount       float32 `float32:"Discount"`
}

type ShoppingCart struct {
	CartID   int     `int:"CartID"`
	Customer string  `string:"Customer"`
	Date     string  `string:"Date"`
	Price    float32 `string:"Price"`
	Type     int     `int:"Type"`
}

type ShoppingCartList struct {
	PurchaseItemID int     `int:"PurchaseItemID"`
	Customer       string  `string:"Customer"`
	ProductID      int     `int:"ProductID"`
	Date           string  `string:"Date"`
	Price          float32 `string:"Price"`
	Quanity        int     `int:"Quanity"`
	IsAlbum        int     `int:"IsAlbum"`
	Discount       float32 `float32:"Discount"`
	Type           int     `int:"Type"`
}

func GetShoppingCartByAccount(sqlConnectionString string, account string) (model []ShoppingCartList, err error) {
	var shoppingCarts []ShoppingCartList
	nullfQuanity := new(sql.NullInt32)
	nullfDiscount := new(sql.NullFloat64)
	queryString := ` SELECT P.*, S.fType  FROM tPurchaseItem P
	INNER JOIN tShoppingCart S ON P.fPurchaseItemID = S.fCartID
	WHERE P.fCustomer = ? AND S.fType = 0 `
	rows, err := util.SQLQuery(sqlConnectionString, queryString, account)

	if err != nil {
		return shoppingCarts, err
	}

	counter := 0
	for rows.Next() {
		var obj ShoppingCartList
		var fPurchaseItemID int
		var fCustomer string
		var fProductID int
		var fDate string
		var fPrice float32
		var fisAlbum int
		var fType int
		err = rows.Scan(&fPurchaseItemID, &fCustomer, &fProductID, &fDate, &fPrice, nullfQuanity, &fisAlbum, nullfDiscount, &fType)

		if nullfQuanity.Valid {
			obj.Quanity = int(nullfQuanity.Int32)
		}

		if nullfDiscount.Valid {
			obj.Discount = float32(nullfDiscount.Float64)
		}

		util.CheckErr(err)
		obj.PurchaseItemID = fPurchaseItemID
		obj.Customer = fCustomer
		obj.ProductID = fProductID
		obj.Date = fDate
		obj.Price = fPrice
		obj.IsAlbum = fisAlbum
		obj.Type = fType
		shoppingCarts = append(shoppingCarts, obj)
		counter++
	}
	if counter == 0 {
		return shoppingCarts, errors.New("select id not found")
	}
	return shoppingCarts, nil
}
