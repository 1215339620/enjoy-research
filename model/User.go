package model

import (
	"gorm.io/gorm"
)

// 用于建数据库
type User struct {
	gorm.Model
	Account         string         `json:"account"`   //账号
	Password        string         `json:"password"`  //密码
	Mobile          string         `json:"mobile"`    //手机号
	Username        string         `json:"username"`  //用户名
	Sex             string         `json:"sex"`       //性别
	Address         string         `json:"address"`   //地址
	School          string         `json:"school"`    //所在院校
	Identity        string         `json:"identity"`  //身份
	Major           string         `json:"major"`     //专业
	Introduce       string         `json:"introduce"` //个人简介
	HeadPortrait    []HeadPortrait `json:"HeadPortrait"`
	ShuoShuos       []ShuoShuo
	TreeHoles       []TreeHole
	TreeHoleComment []TreeHoleComment
}

//用户头像表
type HeadPortrait struct {
	gorm.Model
	UserID uint   `json:"userid"`
	Name   string //文件名字
	Type   string //文件类型
	Size   string //大小
	Url    string //用户头像在服务器上的地址
}
type HeadPortraitParam struct {
	gorm.Model
	Name string //文件名字
	Type string //文件类型
	Size string //大小
	Url  string //用户头像在服务器上的地址
}

// 用于接收前端传来的参数
type ParameRegist struct {
	gorm.Model
	//Account   string `json:"account" binding:"required"`   //账号
	Password string `json:"password" binding:"required"` //密码
	Mobile   string `json:"mobile" binding:"required"`   //手机号
	Username string `json:"username" binding:"required"` //用户名
	//Sex       string `json:"sex" binding:"required"`       //性别
	//Address   string `json:"address" binding:"required"`   //地址
	School string `json:"school" binding:"required"` //所在院校
	//Identity  string `json:"identity" binding:"required"`  //身份
	//Major string `json:"major" binding:"required"` //专业
	//Introduce string `json:"introduce" binding:"required"` //个人简介
}

// 登录参数
type ParameLogin struct {
	Mobile   string `json:"mobile" binding:"required"`   //手机号
	Password string `json:"password" binding:"required"` //密码
}
