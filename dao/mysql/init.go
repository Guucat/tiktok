package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func Init() {

	dsn := "shengyi:123456@tsy@tcp(rm-2vc34w5spf5nm2992eo.mysql.cn-chengdu.rds.aliyuncs.com)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(1000)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(2000)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = db
}
