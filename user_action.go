package main

import (
	"net/http"
	"comm"
	"db"
)

//绑定用户信息
func BindUserInfo(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	appId,_,_,isOk:=AppIsOk(w,r);
	if!isOk{
		return;
	}
	var resultUser = NewUserDto();
	comm.CheckErr(comm.ReadJson(r.Body,&resultUser))

	if resultUser.Rid=="" {

		comm.ResponseError(w,http.StatusBadRequest,"关联的ID不能为空!");
		return;
	}

	authBackend := InitJWTAuthenticationBackend();

	user := db.NewUser()
	if user,_:= user.QueryUserInfo(appId,resultUser.Rid);user!=nil{

		resultUser.Token,_ =authBackend.GenerateToken(user.OpenId)
		resultUser.OpenId=user.OpenId
		comm.WriteJson(w,resultUser)
		return;
	}

	openId :=comm.GenerUUId();

	token,erro := authBackend.GenerateToken(openId);
	comm.CheckErr(erro)
	resultUser.Token =token
	resultUser.OpenId=openId

	user =db.NewUser();
	user.Rid=resultUser.Rid;
	user.OpenId=openId
	user.AppId=appId;
	user.Status=0

	if user.Insert() {
		comm.WriteJson(w,resultUser)
		return;
	}else{
		comm.ResponseError(w,http.StatusBadRequest,"用户添加失败!")
		return
	}

}

//获取认证信息
func GetUserInfo(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	appId,_,_,isOk:=AppIsOk(w,r);
	if !isOk{
		return;
	}
	r_id :=r.FormValue("r_id");
	user :=db.NewUser()
	if user,_:= user.QueryUserInfo(appId,r_id);user!=nil{

		authBackend := InitJWTAuthenticationBackend();
		token,erro := authBackend.GenerateToken(user.OpenId);
		comm.CheckErr(erro)

		userDto :=NewUserDto()
		userDto.Token=token
		userDto.AppId=appId
		comm.WriteJson(w,user)
		return;
	}else{
		comm.ResponseError(w,http.StatusBadRequest,"用户不存在!")
	}
}


