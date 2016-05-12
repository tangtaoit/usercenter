package main

import (
	"log"
)

type User struct {

	//用户ID
	Id int64 `json:"id,omitempty"`
	//公开的用户唯一ID
	OpenId string `json:"open_id,omitempty"`
	//用户关联ID
	Rid string `json:"r_id,omitempty"`
	//APPID
	Appid string `json:"app_id,omitempty"`
	//用户凭证
	Token string `json:"token,omitempty"`

}


func NewUser()  *User{

	return &User{}
}

func init()  {

	log.Println("12333");

}




type ResultError struct {

	ErrCode int `json:"err_code"`
	ErrMsg string `json:"err_msg"`

}



