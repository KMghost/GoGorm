package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// type User struct {
//     gorm.Model
//     ID       int `gorm:"primaryKey"`
//     Name     *string
//     Age      int
//     Birthday *time.Time
// }

// ========================================================================
// 创建钩子
// func (u User) BeforeCreate(tx *gorm.DB) (err error) {
//     if *u.Name == "Jinzhu" {
//         fmt.Println("good"+*u.Name)
//         *u.Name = "zhujin"
//     } else
//     {
//         fmt.Println("bad: "+*u.Name)
//     }
//     return
// }

func main() {
	// dsn := "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// _, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	// now := time.Now()
	// st := "Jinzhu"
	// user := User{Name: &st, Age: 18, Birthday: &now}

	// ========================================================================
	// 添加所有字段

	// result := db.Create(&user) // 通过数据的指针来创建

	// id := user.ID             // 返回插入数据的主键
	// rerr := result.Error        // 返回 error
	// rro := result.RowsAffected // 返回插入记录的条数

	// fmt.Println(id)
	// fmt.Println(rerr)
	// fmt.Println(rro)

	// ========================================================================
	// 使用指定的字段创建记录

	// db.Select("Name","Age","CreatedAt").Create(&user)

	// ========================================================================
	// 除了 Omit 里面的字段，其他都存进数据库
	// db.Omit("Name", "Age", "CreatedAt").Create(&user)

	// ========================================================================
	// 批量插入
	// now := time.Now()
	// var users = []User{{Name: "jinzhu1",Birthday: &now},
	// 	{Name: "jinzhu2",Birthday: &now},
	// 	{Name: "jinzhu3",Birthday: &now}}
	// db.Create(&users)
	// for _, user := range users {
	// 	fmt.Println(user.ID) // 1,2,3
	// }

	// ========================================================================
	// 指定插入批数
	// db.CreateInBatches(users,1)

	// 通过 map 创建记录
	// db.Model(User{}).Create(map[string]interface{}{
	// 	"Name": clause.Expr{SQL: "?",Vars: []interface{}{23}},
	// 	"Location": clause.Expr{SQL: "ST_PointFromText(?)", },
	// })

	// ========================================================================
	// 关联创建
	// db.Create(&User{
	// 	Name: "jinzhu",
	// 	CreditCard: CreditCard{Number: "411111111111"},
	// })
	// u := User{}
	// u.ID  = 35
	j := User{
		Name:       "fe",
		CreditCard: CreditCard{Number: "12"},
	}
	j.ID = 35
	db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&j)
}

// 通过自定义类型创建记录
// type Location struct {
// 	X, Y int
// }
//
// type User struct {
// 	Name     string
// 	Location Location
// }

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard
}
