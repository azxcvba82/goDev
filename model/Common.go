package model

import (
	"encoding/json"
	"main/util"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SMTPConfig struct {
	APIKey string `json:"apiKey"`
	From   string `json:"from"`
}

func GetSysConfig(sqlConnectionString string, key string) (value string, err error) {
	var obj string
	queryString := `SELECT fValue FROM tSysConfig WHERE fKey = ? LIMIT 1 `
	err = util.SQLQueryV2(&obj, sqlConnectionString, true, queryString, key)

	if err != nil {
		return obj, err
	}

	return obj, nil
}

func SentMail(sqlConnectionString string, address string, subject string, msg string) (res string, err error) {

	var obj SMTPConfig
	stmpConfig, err := GetSysConfig(sqlConnectionString, "SMTPConfig")
	err = json.Unmarshal([]byte(stmpConfig), &obj)

	from := mail.NewEmail("platform user", obj.From)
	to := mail.NewEmail("no-reply", address)
	message := mail.NewSingleEmail(from, subject, to, msg, msg)
	client := sendgrid.NewSendClient(obj.APIKey)
	response, err := client.Send(message)

	if err != nil {
		return response.Body, nil
	} else {
		return response.Body, err
	}
}
