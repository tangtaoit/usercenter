package main

import (
	"log"
)

type UserDto struct {

	//公开的用户唯一ID
	OpenId string `json:"open_id,omitempty"`
	//用户关联ID
	Rid string `json:"r_id,omitempty"`
	//APPID
	AppId string `json:"app_id,omitempty"`
	//用户凭证
	Token string `json:"token,omitempty"`

}


func NewUserDto()  *UserDto{

	return &UserDto{}
}

func init()  {

	log.Println("12333");

}








