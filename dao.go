package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var db *sql.DB

func init() {

	setting :=GetSetting()


	connInfo := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",setting.MysqlUser,setting.MysqlPassword,setting.MysqlHost,setting.MysqlDB)

	fmt.Println(connInfo);

	db, _ = sql.Open("mysql",connInfo)
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func GetDB()  *sql.DB{

	return db;
}

func AddUserInfo(user *User) bool {

	_, err := db.Exec("insert into users(app_id,open_id,r_id) values(?,?,?)",user.Appid,user.OpenId,user.Rid)

	if err!=nil {

		CheckErr(err)
		return false;
	}

	return true;


}

func QueryUserInfo(app_id string,r_id string)  (user *User, er error) {

	rows, err := db.Query("select id,open_id,r_id from users where app_id=? and r_id=?",app_id,r_id)
	defer rows.Close()
	if err!=nil{

		return nil,err;
	}
	if rows.Next() {
		var id int64
		var rid *string
		var openId *string
		err = rows.Scan(&id,&openId,&rid)
		if err!=nil{

			return nil,err
		}

		userModel :=NewUser();
		userModel.OpenId=*openId;
		userModel.Rid=*rid;
		//userModel.Id=id;
		return userModel,nil;
	}

	return nil,err;
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