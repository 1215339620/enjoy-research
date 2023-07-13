package mysql

import (
	"gorm.io/gorm"
	"web_app/model"
)

// 注册
func Regist(u *model.User) (err error) {
	//存到数据库
	err = DB.Create(&u).Error
	return err
}

// 登录
func Login(user *model.User) (u *model.User, err error) {
	//查询数据库,判断用户是否存在
	//将查到的数据放到ub中,ub是一个指针类型，可以接受切片，即多个数据
	err = DB.Where("mobile = ?", user.Mobile).Take(&u).Error
	//用户不存在,返回一个错误:用户不存在
	if err == gorm.ErrRecordNotFound {
		return nil, ErrorUserNoExist
	}
	//数据库查询失败
	if err != nil {
		return nil, err
	}
	//判断密码是否正确
	if user.Password != u.Password {
		return nil, ErrorInvalidPassword
	}
	return u, err
}

// 用户回显数据
func ReturnDataMysql(mobile string) (user *model.User, err error) {
	//根据手机号向用户表查询数据
	err = DB.Where("mobile", mobile).Take(&user).Error
	return user, err

}

// 更新用户信息
func UpdateUser(u *model.User) (err error) {
	err = DB.Where("mobile", u.Mobile).Updates(u).Error
	return err
}

// 删除用户
func DeleteUser(mobile string) (err error) {
	DB.Where("mobile", mobile).Delete(&model.User{})
	return err
}

func HeadPortrait(file *model.HeadPortrait) (err error) {
	err = DB.Table("head_portraits").Create(file).Error
	return err
}
