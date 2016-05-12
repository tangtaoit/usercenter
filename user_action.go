package main

import (
	"net/http"
	"github.com/gorilla/mux"
)


//绑定用户信息
func BindUserInfo(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if isOk:=appIsOk(w,r);!isOk{
		return;
	}

	vars := mux.Vars(r);
	app_id := vars["app_id"];
	var resultUser = NewUser();
	ReadJson(r.Body,&resultUser)

	if isok:=IsExistUser(app_id,resultUser.Rid);isok {
		ResponseError(w,http.StatusBadRequest,"用户信息已存在!");
		return;
	}

	openId :=GenerOpenId();
	authBackend := InitJWTAuthenticationBackend();
	token,erro := authBackend.GenerateToken(openId);
	CheckErr(erro)


	user :=NewUser();
	user.Token=token;
	user.Rid=resultUser.Rid;
	user.OpenId=openId;
	user.Appid=app_id;


	if AddUserInfo(user) {
		user.Appid="";
		WriteJson(w,user)
		return;
	}

}

//获取认证信息
func GetUserInfo(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if isOk:=appIsOk(w,r);!isOk{
		return;
	}
	vars := mux.Vars(r);
	app_id := vars["app_id"];

	r_id :=r.FormValue("r_id");

	if user:= QueryUserInfo(app_id,r_id);user!=nil{

		WriteJson(w,user)
		return;
	}
}



func appIsOk(w http.ResponseWriter,r *http.Request) bool {
	vars := mux.Vars(r);
	app_id := vars["app_id"];
	app_key := vars["app_key"];

	if err:=AuthApp(app_id,app_key);err!=nil{
		ResponseError(w,http.StatusUnauthorized,"appid和appkey不合法!")
		return false;
	}

	return true;
}
