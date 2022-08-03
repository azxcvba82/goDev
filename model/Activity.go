package model

import (
	"main/util"

	"gopkg.in/guregu/null.v4"
)

type Activity struct {
	Id        int         `db:"fId"`
	Launcher  null.String `db:"fLauncher"`
	StartTime string      `db:"fStartTime"`
	EndTime   string      `db:"fEndTime"`
	Title     string      `db:"fTitle"`
	PhotoPath string      `db:"fPhotoPath"`
}

func EventQuery(sqlConnectionString string, account string) (model []Activity, err error) {
	var activities []Activity
	queryString := ""

	if account == "" {
		queryString = `SELECT fId, fLauncher, fStartTime, fEndTime, fTitle, fPhotoPath  FROM tActivity `
		err = util.SQLQueryV2(&activities, sqlConnectionString, true, queryString)
	} else {
		queryString = `SELECT fId, fLauncher, fStartTime, fEndTime, fTitle, fPhotoPath  FROM tActivity WHERE fLauncher = ? `
		err = util.SQLQueryV2(&activities, sqlConnectionString, true, queryString, account)
	}

	if err != nil {
		return activities, err
	}

	return activities, nil
}
