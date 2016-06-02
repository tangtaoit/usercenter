package main

import (
	"log"
)

type UserDto struct {

	//公开的用户唯一ID
	OpenId string `json:"open_id"`
	//用户关联ID
	Rid string `json:"r_id"`
	//用户凭证
	Token string `json:"token"`

}


func NewUserDto()  *UserDto{

	return &UserDto{}
}

func init()  {

	log.Println("12333");

}








