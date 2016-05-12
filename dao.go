package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {

	db, _ = sql.Open("mysql", "tangtao:123456@tcp(172.30.121.158:3306)/sampledb?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func AddUserInfo(user *User) bool {

	_, err := db.Exec("insert into users(app_id,open_id,r_id,token) values(?,?,?,?)",user.Appid,user.OpenId,user.Rid,user.Token)

	if err!=nil {

		CheckErr(err)
		return false;
	}

	return true;


}

func QueryUserInfo(app_id string,r_id string)  *User {

	rows, err := db.Query("select id,open_id,r_id,token from users where app_id=? and r_id=?",app_id,r_id)
	defer rows.Close()
	CheckErr(err)
	if rows.Next() {
		var id int64
		var rid *string
		var openId *string
		var token *string
		err = rows.Scan(&id,&openId,&rid,&token)
		CheckErr(err)

		userModel :=NewUser();
		userModel.Token=*token;
		userModel.OpenId=*openId;
		userModel.Rid=*rid;
		//userModel.Id=id;
		return userModel;
	}

	return nil;
}

func IsExistUser(app_id string,r_id string) bool {
	rows, err := db.Query("select count(*) cn from users where app_id=? and r_id=?",app_id,r_id)
	defer rows.Close()
	CheckErr(err)

	if rows.Next() {
		var cn int

		err = rows.Scan(&cn)
		CheckErr(err)

		if cn>0{

			return true;
		}
	}

	return false;
}