package mysql

import (
	"web_app/model"
)

func CreateTable() (err error) {
	//err = DB.AutoMigrate(&model.User{})
	//DB.AutoMigrate(&model.Province{})
	//DB.AutoMigrate(&model.City{})
	//DB.AutoMigrate(&model.School{})
	//DB.AutoMigrate(&model.Academy{})
	//DB.AutoMigrate(&model.Major{})
	//DB.AutoMigrate(&model.Subject{})
	//DB.AutoMigrate(&model.MajorSubject{})
	//DB.AutoMigrate(&model.File{})
	//DB.AutoMigrate(&model.ShuoShuo{})
	//DB.AutoMigrate(&model.Comment{})
	DB.AutoMigrate(&model.HeadPortrait{})
	DB.AutoMigrate(&model.TreeHole{})
	DB.AutoMigrate(&model.TreeHoleComment{})
	return
}
