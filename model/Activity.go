package model

import (
	"database/sql"
	"main/util"
)

type Activity struct {
	Id        int    `int:"Id"`
	Launcher  string `string:"Launcher"`
	StartTime string `string:"StartTime"`
	EndTime   string `string:"EndTime"`
	Title     string `string:"Title"`
	PhotoPath string `string:"PhotoPath"`
}

func EventQuery(sqlConnectionString string, account string) (model []Activity, err error) {
	var activities []Activity
	nullfLauncher := new(sql.NullString)
	queryString := ""
	var rows *sql.Rows

	if account == "" {
		queryString = `SELECT *  FROM tActivity `
		rows, err = util.SQLQuery(sqlConnectionString, queryString)
	} else {
		queryString = `SELECT *  FROM tActivity WHERE fLauncher = ? `
		rows, err = util.SQLQuery(sqlConnectionString, queryString, account)
	}

	if err != nil {
		return activities, err
	}

	for rows.Next() {
		var obj Activity
		var fId int
		var fStartTime string
		var fEndTime string
		var fTitle string
		var fPhotoPath string
		err = rows.Scan(&fId, nullfLauncher, &fStartTime, &fEndTime, &fTitle, &fPhotoPath)

		util.CheckErr(err)
		if nullfLauncher.Valid {
			obj.Launcher = nullfLauncher.String
		}
		obj.Id = fId
		obj.StartTime = fStartTime
		obj.EndTime = fEndTime
		obj.Title = fTitle
		obj.PhotoPath = fPhotoPath
		activities = append(activities, obj)
	}
	return activities, nil
}
