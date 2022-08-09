package model

import (
	"errors"
	"main/util"

	"gopkg.in/guregu/null.v4"
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
	PurchaseItemID int        `db:"fPurchaseItemID"`
	Customer       string     `db:"fCustomer"`
	ProductID      int        `db:"fProductID"`
	Date           string     `db:"fDate"`
	Price          float32    `db:"fPrice"`
	Quanity        null.Int   `db:"fQuanity"`
	IsAlbum        int        `db:"fisAlbum"`
	Discount       null.Float `db:"fDiscount"`
	Type           int        `db:"fType"`
	CoverPath      string     `db:"fCoverPath"`
	ProductName    string     `db:"fProductName"`
	AlbumName      string     `db:"fAlbumName"`
	ALPrice        float32    `db:"fALPrice"`
	ALDiscount     null.Float `db:"fALDiscount"`
}

func GetShoppingCartByAccount(sqlConnectionString string, account string) (model []ShoppingCartList, err error) {
	var shoppingCarts []ShoppingCartList
	queryString := ` SELECT P.fPurchaseItemID, P.fCustomer, P.fProductID, P.fDate, P.fPrice, P.fQuanity, P.fisAlbum, P.fDiscount,   
	S.fType, A.fCoverPath, I.fProductName, A.fAlbumName, A.fALPrice, A.fDiscount AS fALDiscount  FROM tPurchaseItem P
	INNER JOIN tShoppingCart S ON P.fPurchaseItemID = S.fCartID
	LEFT JOIN tProducts I ON P.fProductID = I.fProductID
	LEFT JOIN tAlbum A ON I.fAlbumID = A.fAlbumID
	WHERE P.fCustomer = ? AND S.fType = 0 `
	err = util.SQLQueryV2(&shoppingCarts, sqlConnectionString, true, queryString, account)

	if err != nil {
		return shoppingCarts, err
	}

	if shoppingCarts == nil {
		return shoppingCarts, errors.New("select id not found")
	}
	return shoppingCarts, nil
}
