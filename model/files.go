package model

import "gorm.io/gorm"

type ShuoShuo struct {
	gorm.Model
	UserID   uint `json:"userID"` //用户手机号
	Content  string
	Files    []File    //图片
	LikeNum  int       //点赞数
	Comments []Comment //评论
}

type File struct {
	gorm.Model
	Name       string //文件名字
	Type       string //文件类型
	Size       string //大小
	Url        string //文件路径
	ShuoShuoID uint
}

type Filea struct {
	gorm.Model
	Name string //文件名字
	Type string //文件类型
	Size string //大小
}

//type File struct {
//	gorm.Model
//	UserPhone   string     //用户手机号
//	Description string     //文件描述
//	Images      []Images   //图片
//	LikeNum     int        //点赞数
//	Comments    []Comments //评论
//}
//type Images struct {
//	gorm.Model
//	ImageID string //图片ID
//	Name    string //图片名字
//	Size    string //大小
//}

type Comment struct {
	gorm.Model
	Replays    []Comment
	CommentID  uint   //评论父级ID  0
	Content    string //内容
	ShuoShuoID uint
}
