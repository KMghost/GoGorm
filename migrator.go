package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID       int       `gorm:"primaryKey"`
	Name     string    `gorm:"type: varchar(100) comment '用户名' " `
	Age      int       `gorm:"type: int(20) comment '年龄'"`
	Birthday time.Time `gorm:"type: datetime comment '生日'"`
}
type Book struct {
	gorm.Model
	Name   string `json:"name" gorm:"type:varchar(20) not null comment '书名';"`
	Count  string `json:"count" gorm:"type:varchar(10) not null comment '价格';"`
	Author string `json:"author" gorm:"type:varchar(20) not null comment '作者';"`
	Type   string `json:"type" gorm:"type:varchar(20) not null default 1 comment '类型';"`
}

func main() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '用户表'").Migrator().CreateTable(&User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8 comment '图书表'").Migrator().CreateTable(&Book{})
}
