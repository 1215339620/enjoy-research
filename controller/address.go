package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app/dao/mysql"
)

// 查找所有省
func SearchProvince(c *gin.Context) {
	//直接去查数据库
	province, err := mysql.SearchProvince()
	if err != nil {
		zap.L().Error("查询省出错了", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	//成功返回省的信息
	ResponseSuccess(c, province)
}

// 查找所有市
func SearchCity(c *gin.Context) {
	province, _ := strconv.Atoi(c.Query("provinceid"))
	provinceid := uint(province)
	citys, err := mysql.SearchCity(provinceid)
	if err != nil {
		zap.L().Error("查询市失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, citys)

}

// 查找市下面的学校
func SearchSchool(c *gin.Context) {
	city, _ := strconv.Atoi(c.Query("cityid"))
	cityid := uint(city)
	citys, err := mysql.SearchSchool(cityid)
	if err != nil {
		zap.L().Error("查询学校失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, citys)
}

func SearchAcademies(c *gin.Context) {
	school, _ := strconv.Atoi(c.Query("schoolid"))
	schoolid := uint(school)
	academies, err := mysql.SearchAcademies(schoolid)
	if err != nil {
		zap.L().Error("查询专业失败", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, academies)
}

//查询所有的地址
func SeleteAllAddress(c *gin.Context) {
	err, address := mysql.SeleteAllAddress()
	if err != nil {
		zap.L().Error("查询所有地址", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, address)
}

//模糊查询学校
func SeleteLikeSchool(c *gin.Context) {
	//mysql.SeleteLikeSchool()

}
