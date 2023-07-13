package mysql

import (
	"gorm.io/gorm"
	"web_app/model"
)

func CreateShuoShuo(shuoshuo *model.ShuoShuo) (err error) {
	err = DB.Create(shuoshuo).Error
	return
}
func UploadFiles(file *model.Filea) (err error) {
	err = DB.Table("files").Create(file).Error
	return
}
func MyShuoShuo(userid uint) (err error, ShuoShuos []model.ShuoShuo) {
	err = DB.Model(&model.User{Model: gorm.Model{ID: userid}}).Association("ShuoShuos").Find(&ShuoShuos)
	return
}
