package main

// import (
//     driver "GoGorm/database"
//     "gorm.io/gorm/clause"
// )

func main() {
	// db := driver.DB()
	// user := driver.User{}
	// user2 := driver.User{}

	// 获取第一条记录（主键升序）
	// db.First(&user2)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取一条记录，没有指定排序字段
	// db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	// db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	// result := db.First(&user, 1)
	// rowsa := result.RowsAffected // 返回找到的记录数
	// res_err := result.Error        // returns error or nil

	// 检查 ErrRecordNotFound 错误
	// r_err := errors.Is(result.Error, gorm.ErrRecordNotFound)

	// ========================================================================
	// 用主键索引
	// var users []driver.User
	// db.First(&user, 35)
	// // SELECT * FROM users WHERE id = 10;
	//
	// db.First(&user, "35")
	// // SELECT * FROM users WHERE id = 10;
	//
	// db.Find(&users, []int{35, 36, 37})
	// // SELECT * FROM users WHERE id IN (1,2,3);
	//
	// // 如果主键是字符串（例如像 uuid），查询将被写成这样
	// db.First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	// // SELECT * FROM users WHERE id = "1b74413f-f3b8-409f-ac47-e8c062e3472a";

	// ========================================================================
	// 检索全部对象
	// var users []driver.User
	// result := db.Find(&users)
	// // fmt.Printf("%+v",result)
	// res := result.RowsAffected // 返回找到的记录数，相当于 `len(users)`
	// ss := result.Error        // returns error

	// ========================================================================
	// String 条件
	// var user driver.User
	// var users []driver.User
	// // 获取第一条匹配的记录
	// db.Where("name = ?", "jinzhu").First(&user)
	// // SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
	//
	// // 获取全部匹配的记录
	// db.Where("name <> ?", "jinzhu").Find(&users)
	// // SELECT * FROM users WHERE name <> 'jinzhu';
	//
	// // IN
	// db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
	// // SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');
	//
	// // LIKE
	// db.Where("name LIKE ?", "%jin%").Find(&users)
	// // SELECT * FROM users WHERE name LIKE '%jin%';
	//
	// // AND
	// db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	// // SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

	// Time
	// var layout string = "2006-01-02 15:04:05"
	// var timeStr string = "2021-12-17 22:44:23"

	// timeObj1, _ := time.Parse(layout, timeStr)   // utc time
	// fmt.Println(timeObj1)

	// lastWeek, _ := time.ParseInLocation(layout, timeStr, time.Local)  // CST time
	// db.Where("birthday = ?", lastWeek).Find(&users)
	// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

	// BETWEEN
	// db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
	// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';

	// for _,data := range users{
	//     fmt.Printf("%v\n",data)
	// }

	// ========================================================================
	// Map 条件
	// Struct
	// var user driver.User
	// var users []driver.User
	// db.Where(&driver.User{Name: "jinzhu", Age: 20}).First(&user)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

	// Map
	// db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
	// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

	// 主键切片条件
	// db.Where([]int64{37, 38, 22}).Find(&users)
	// SELECT * FROM users WHERE id IN (20, 21, 22);

	// fmt.Printf("%v\n",user)
	// for i, data := range users{
	//     fmt.Printf("%d  %v\n",i,data)
	// }

	// ========================================================================
	// 结构体查询 (前面结构体的值会传给后面的参数，如果没有则为空, 大小写不敏感)
	// var users []driver.User
	// db.Where(&driver.User{Name: "jinzhu"}, "name", "Age").Find(&users)
	// // SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
	//
	// // db.Where(&driver.User{Age: 20}, "Age").Find(&users)
	// // SELECT * FROM users WHERE age = 20;

	// for i, data := range users{
	//     fmt.Printf("%d  %v\n",i,data)
	// }

	// ========================================================================
	// Scan
	// var result []Result
	// // result.Name = "jinzhu"
	// db.Table("users").Select("name", "age").Where("age = ?", "0").Debug().Scan(&result)
	// fmt.Printf("%+v",result)

	// ========================================================================
	// Distinct
	// var results []driver.User
	// db.Distinct("name", "age").Order("name, age desc").Debug().Find(&results)
	// for i, data := range results{
	//     fmt.Printf("%d  %v\n",i,data)
	// }

	// ========================================================================
	// order by
	// var user driver.User
	// db.Clauses(clause.OrderBy{
	//     Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	// }).Debug().Find(&user)

}

type Result struct {
	Name string
	Age  int
}
