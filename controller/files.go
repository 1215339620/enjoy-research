package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/dao/mysql"
	"web_app/model"
)

func UploadFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		zap.L().Error("上传失败", zap.Error(err))
		ResponseError(c, CodeUploadError)
		return
	}
	// 处理文件
	files := form.File["files"]
	for _, file := range files {
		err := c.SaveUploadedFile(file, "./static/"+file.Filename)
		if err != nil {
			fmt.Println(err)
		}
		f := &model.Filea{
			Name: file.Filename,
			Type: file.Header.Get("Content-Type"),
			Size: strconv.Itoa(int(file.Size)),
		}
		err = mysql.UploadFiles(f)
		if err != nil {
			zap.L().Error("插入数据库失败", zap.Error(err))
			ResponseError(c, CodeServerBusy)
			return
		}
	}

	ResponseSuccess(c, CodeSuccess)
}

func CreateShuoShuo(c *gin.Context) {
	shuoshuo := new(model.ShuoShuo)
	err := c.ShouldBindJSON(&shuoshuo)
	fmt.Println(shuoshuo.UserID)
	if err != nil {
		zap.L().Error("发表失败", zap.Error(err))
		ResponseError(c, CodeCreateError)
		return
	}
	err = mysql.CreateShuoShuo(shuoshuo)
	if err != nil {
		zap.L().Error("插入数据库失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
	//var FileIds []uint
	//var Files []model.File
	//mysql.DB.Create(model.ShuoShuo{Content: "考研", Files: Files, LikeNum: 0})
}
func MyShuoShuo(c *gin.Context) {
	user, _ := strconv.Atoi(c.Query("userid"))
	userid := uint(user)
	err, shuos := mysql.MyShuoShuo(userid)
	if err != nil {
		zap.L().Error("查询发表资料失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, shuos)
	//mysql.DB.Model().Preload("ShuoShuos")
}
