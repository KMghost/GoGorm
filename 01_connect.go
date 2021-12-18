package main

// import (
//
// )

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

	// now := time.Now()
	// st := "Jinzhu"
	// database := database.DB()
	// User := database.Use
	// user := User{Name: &st, Age: 18, Birthday: &now}

	// ========================================================================
	// 添加所有字段

	// result := database.Create(&user) // 通过数据的指针来创建

	// id := user.ID             // 返回插入数据的主键
	// rerr := result.Error        // 返回 error
	// rro := result.RowsAffected // 返回插入记录的条数

	// fmt.Println(id)
	// fmt.Println(rerr)
	// fmt.Println(rro)

	// ========================================================================
	// 使用指定的字段创建记录

	// database.Select("Name","Age","CreatedAt").Create(&user)

	// ========================================================================
	// 除了 Omit 里面的字段，其他都存进数据库
	// database.Omit("Name", "Age", "CreatedAt").Create(&user)

	// ========================================================================
	// 批量插入
	// now := time.Now()
	// var users = []User{{Name: "jinzhu1",Birthday: &now},
	// 	{Name: "jinzhu2",Birthday: &now},
	// 	{Name: "jinzhu3",Birthday: &now}}
	// database.Create(&users)
	// for _, user := range users {
	// 	fmt.Println(user.ID) // 1,2,3
	// }

	// ========================================================================
	// 指定插入批数
	// database.CreateInBatches(users,1)

	// 通过 map 创建记录
	// database.Model(User{}).Create(map[string]interface{}{
	// 	"Name": clause.Expr{SQL: "?",Vars: []interface{}{23}},
	// 	"Location": clause.Expr{SQL: "ST_PointFromText(?)", },
	// })

	// ========================================================================
	// 关联创建
	// database.Create(&User{
	// 	Name: "jinzhu",
	// 	CreditCard: CreditCard{Number: "411111111111"},
	// })
	// u := User{}
	// u.ID  = 35
	// j := User{
	// 	Name:       "fe",
	// 	CreditCard: CreditCard{Number: "12"},
	// }
	// j.ID = 35
	// database.Clauses(clause.OnConflict{
	// 	UpdateAll: true,
	// }).Create(&j)
}
