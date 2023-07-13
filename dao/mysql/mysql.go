package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"web_app/settings"
)

var DB *gorm.DB

// Init 初始化MySQL连接
func Init(cfg *settings.MySQLConfig) (err error) {
	// "user:password@tcp(host:port)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	DB, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{})
	if err != nil {
		return
	} else {
		sqlDB, _ := DB.DB()
		sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
		sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	}

	//自动建表

	//u := user.User{
	//	Model: gorm.Model{
	//		ID: 11,
	//	},
	//	UserName: "luosen",
	//	File: []file.File{
	//		{Name: "file1"},
	//		{Name: "file2"},
	//	},
	//}
	//DB.Create(&u)
	return
}

// 自动建表方法
func creatTable(dst interface{}) {
	if !DB.Migrator().HasTable(dst) {
		err := DB.AutoMigrate(dst)
		if err != nil {
			return
		}
		if DB.Migrator().HasTable(dst) {
			fmt.Println("表创建成功")
		} else {
			fmt.Println("表创建失败")
		}
	} else {
		fmt.Println("表已存在")
	}
}
