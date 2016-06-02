package db

import (
	_ "github.com/go-sql-driver/mysql"
	"comm"
	"fmt"
)

type User struct {
	AppId string
	OpenId string
	Rid string
	Status int
}

func NewUser()  *User {

	return &User{}
}


func (self *User) Insert() bool {

	_, err := db.Exec("insert into users(app_id,open_id,r_id,status) values(?,?,?,?)",self.AppId,self.OpenId,self.Rid,self.Status)
	if err!=nil {

		comm.CheckErr(err)
		return false;
	}
	return true;

}

func (self *User)  QueryUserInfo(app_id string,r_id string)  (user *User, er error) {

	rows, err := db.Query("select id,open_id,r_id from users where app_id=? and r_id=?",app_id,r_id)
	defer rows.Close()
	if err!=nil{
		fmt.Println("error=",err)
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

func (self *User)  IsExistUser(app_id string,r_id string) bool {
	rows, err := db.Query("select count(*) cn from users where app_id=? and r_id=?",app_id,r_id)
	defer rows.Close()
	comm.CheckErr(err)

	if rows.Next() {
		var cn int

		err = rows.Scan(&cn)
		comm.CheckErr(err)

		if cn>0{

			return true;
		}
	}

	return false;
}