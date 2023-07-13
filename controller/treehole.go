package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web_app/dao/mysql"
	"web_app/model"
)

func CreateTreeHole(c *gin.Context) {
	treehole := new(model.TreeHole)
	err := c.ShouldBindJSON(&treehole)
	if err != nil {
		zap.L().Error("发表失败", zap.Error(err))
		ResponseError(c, CodeCreateError)
		return
	}
	err = mysql.CreateTreeHole(treehole)
	if err != nil {
		zap.L().Error("插入数据库失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, CodeSuccess)
}

//查询所有的树洞，并做分页查询
func SeleteAllTreeHole(c *gin.Context) {

}

func GetTreeHole(c *gin.Context) {
	treeholes, err := mysql.GetTreeHole()
	if err != nil {
		zap.L().Error("查询树洞错误", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, treeholes)

}
