package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web_app/controller"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "./static")
	v1 := r.Group("/api/v1")
	v1.POST("/regist", controller.Regist)
	v1.POST("/login", controller.Login)
	//应用jwt做是否登录验证
	//v1.Use(middlewares.JWTAuthMiddleware())
	u := v1.Group("/user")
	{
		//回显用户信息
		u.GET("/returndata", controller.ReturnData)
		//修改用户信息
		u.POST("/updateuser", controller.UpdateUser)
		//用户注销
		u.DELETE("/deleteuser/:mobile", controller.DeleteUser)
		//用户头像上传
		u.POST("/headportrait", controller.HeadPortrait)
	}
	file := v1.Group("/file")
	{
		//用户上传文件
		file.POST("/uploadfile", controller.UploadFiles)
		//发表资料
		file.POST("/createshuoshuo", controller.CreateShuoShuo)
		//我发布的资料
		file.GET("/myshuoshuo", controller.MyShuoShuo)

	}
	treehole := v1.Group("/treehole")
	{
		//发表树洞
		treehole.POST("/createtreehole", controller.CreateTreeHole)
		//查询所有的树洞
		treehole.GET("/seletealltreehole", controller.SeleteAllTreeHole)
		//获取树洞
		treehole.GET("/gettreehole", controller.GetTreeHole)
	}
	address := v1.Group("/address")
	{
		//查找所有省
		address.GET("/searchprovince", controller.SearchProvince)
		//查找所有市
		address.GET("/searchcity", controller.SearchCity)
		//查找市下面的学校
		address.GET("/searchschool", controller.SearchSchool)
		//查找学校下面的专业
		address.GET("searchacademies", controller.SearchAcademies)
		//查询所有的地址
		address.GET("/seletealladdress", controller.SeleteAllAddress)
		//模糊查询学校
		address.GET("seletelikeschool", controller.SeleteLikeSchool)

	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
