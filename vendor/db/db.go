package db

import (
	"fmt"
	"database/sql"
	"comm"
	"config"
	"net/url"
)
var db *sql.DB
func init() {




	setting :=config.GetSetting()
	connInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=%s&parseTime=true",setting.MysqlUser,setting.MysqlPassword,setting.MysqlHost,setting.MysqlDB,url.QueryEscape("Asia/Beijing"))
	fmt.Println(connInfo);
	var err error;
	db, err = sql.Open("mysql",connInfo)
	if err!=nil{
		comm.CheckErr(err)
	}

	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}
