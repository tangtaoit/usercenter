package main

import (
	"net/http"
	"comm"
	"db"
	"github.com/tangtaoit/util"
	"io/ioutil"
	"strings"
	"log"
	"fmt"
)

//绑定用户信息
func BindUserInfo(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	appId,appKey,basesign,isOk:=AppIsOk(w,r);
	if!isOk{
		return;
	}

	sign := r.Header.Get("sign")
	signs :=strings.Split(sign,".")
	if len(signs)!=2 {
		util.ResponseError(w,http.StatusBadRequest,"非法请求!")
		return
	}

	bodyBytes,err := ioutil.ReadAll(r.Body)
	if err!=nil{
		util.ResponseError(w,http.StatusBadRequest,"参数有误!")
		return;
	}

	var resultUser = NewUserDto();
	util.CheckErr(util.ReadJsonByByte(bodyBytes,&resultUser))

	if resultUser.Rid=="" {

		util.ResponseError(w,http.StatusBadRequest,"关联的ID不能为空!");
		return;
	}

	var signMap map[string]interface{}
	util.CheckErr(util.ReadJsonByByte(bodyBytes,&signMap))

	wantSign := util.SignWithBaseSign(signMap,appKey,basesign,nil)
	gotSign :=signs[1];
	if wantSign!=gotSign {
		log.Println("wantSign: ",wantSign,"gotSign: ",gotSign)
		util.ResponseError(w,http.StatusBadRequest,"签名不匹配!")
		return
	}

	authBackend := InitJWTAuthenticationBackend();

	user := db.NewUser()
	if user,err:= user.QueryUserInfo(appId,resultUser.Rid);user!=nil{

		fmt.Println("error=",err)
		resultUser.Token,_ =authBackend.GenerateToken(user.OpenId,appId)
		resultUser.OpenId=user.OpenId
		util.WriteJson(w,resultUser)
		return;
	}

	openId :=comm.GenerUUId();

	token,erro := authBackend.GenerateToken(openId,appId);
	comm.CheckErr(erro)
	resultUser.Token =token
	resultUser.OpenId=openId

	user =db.NewUser();
	user.Rid=resultUser.Rid;
	user.OpenId=openId
	user.AppId=appId;
	user.Status=0

	if user.Insert() {
		util.WriteJson(w,resultUser)
		return;
	}else{
		util.ResponseError(w,http.StatusBadRequest,"用户添加失败!")
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
		token,erro := authBackend.GenerateToken(user.OpenId,appId);
		comm.CheckErr(erro)

		userDto :=NewUserDto()
		userDto.Token=token
		comm.WriteJson(w,user)
		return;
	}else{
		comm.ResponseError(w,http.StatusBadRequest,"用户不存在!")
	}
}


