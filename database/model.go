package database

import (
	"time"
)

type User struct {
	// gorm.Model
	ID       int `gorm:"primaryKey"`
	Name     string
	Age      int
	Active   bool
	Birthday *time.Time
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

// type CreditCard struct {
//     gorm.Model
//     Number string
//     UserID uint
// }

// type User struct {
// 	gorm.Model
// 	Name       string
// 	CreditCard CreditCard
// }
