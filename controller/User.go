package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/dao/mysql"
	"web_app/middlewares/jwt"
	"web_app/model"
	"web_app/serve/User"
)

// 用户注册
func Regist(c *gin.Context) {
	//获取注册表单的参数
	u := new(model.ParameRegist)
	//看字段数据是否都有
	if err := c.ShouldBindJSON(&u); err != nil {
		//不符合要求，则不让提交
		zap.L().Error("注册参数：", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	user := &model.User{
		Mobile:   u.Mobile,
		Password: u.Password,
		Username: u.Username,
		School:   u.School,
	}
	//参数正确,存到数据库里面
	err := mysql.Regist(user)
	if err != nil {
		zap.L().Error("服务繁忙", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "注册成功")
}

// 用户登录
func Login(c *gin.Context) {
	u := new(model.ParameLogin)
	if err := c.ShouldBindJSON(&u); err != nil {
		zap.L().Error("参数错误", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//业务逻辑处理
	user, err := User.Login(u)
	if err != nil {
		zap.L().Error("login.UserLogin failed", zap.String("username", u.Mobile), zap.Error(err))
		//判断出错类型-用户不存在,查不到
		if errors.Is(err, mysql.ErrorUserNoExist) {
			ResponseError(c, CodeUserNotExist)
			return
		} else if errors.Is(err, mysql.ErrorInvalidPassword) {
			ResponseError(c, CodeInvalidPassword)
		} else {
			//其他问题返回服务繁忙
			ResponseError(c, CodeServerBusy)
			return
		}
	}
	//生成JWT
	token, err := jwt.GenToken(int64(int(user.ID)), user.Username)
	if err != nil {
		return
	}

	//查询没有问题则返回用户部分数据给前端
	ResponseSuccess(c, gin.H{
		"userdata": user,
		"token":    token,
	})
}

// 回显用户数据
func ReturnData(c *gin.Context) {
	mobile := c.Query("mobile")
	//拿到参数查数据库
	user, err := mysql.ReturnDataMysql(mobile)
	if err != nil {
		zap.L().Error("回显查数据失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, user)
}

// 用户修改信息
func UpdateUser(c *gin.Context) {
	u := new(model.User)
	if err := c.ShouldBindJSON(&u); err != nil {
		//不符合要求，则不让提交
		zap.L().Error("参数绑定失败：", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	//没有问题修改数据库
	err := mysql.UpdateUser(u)
	if err != nil {
		zap.L().Error("用户更新数据失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "更新成功")
}

// 用户注销
func DeleteUser(c *gin.Context) {
	mobile := c.Param("mobile")
	//修改数据库信息
	err := mysql.DeleteUser(mobile)
	if err != nil {
		zap.L().Error("服务繁忙", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, "删除成功")
}

/**
用户头像上传
*/
func HeadPortrait(c *gin.Context) {
	file, err := c.FormFile("headPortrait")
	value := c.Query("userid")
	user, _ := strconv.Atoi(value)
	userid := uint(user)
	if err != nil {
		zap.L().Error("头像上传失败", zap.Error(err))
		ResponseError(c, CodeUploadError)
		return
	}
	file.Filename = value + file.Filename
	//把文件存到本地
	c.SaveUploadedFile(file, "./static/headportrait/"+file.Filename)
	//用户id会在用户修改信息后自动放入表中
	portrait := model.HeadPortrait{
		UserID: userid,
		Name:   file.Filename,
		Type:   file.Header.Get("Content-Type"),
		Size:   strconv.Itoa(int(file.Size)),
		Url:    "http://localhost:8080/static/headportrait/" + file.Filename,
	}
	//把数据存到数据库
	err = mysql.HeadPortrait(&portrait)
	if err != nil {
		zap.L().Error("数据库插入头像数据失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	type Url struct {
		Url string
	}

	url := Url{
		Url: "http://localhost:8080/static/headportrait/" + file.Filename,
	}
	ResponseSuccess(c, url)
}
