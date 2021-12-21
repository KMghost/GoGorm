package main

// import (
//     driver "GoGorm/database"
//     "fmt"
//     "gorm.io/gorm"
// )
// type User driver.User
func main() {
	//     db := driver.DB()
	// var user []User
	// ========================================================================
	// 子查询
	// db.Where("age > (?)", db.Table("users").Select("AVG(age)")).Debug().Find(&user)

	// subQuery := db.Select("AVG(age)").Where("name LIKE ?", "name%").Table("users")
	// db.Select("AVG(age) as avgage").Group("name").Having("AVG(age) > (?)", subQuery).Debug().Find(&user)

	// ========================================================================
	// Named Argument
	// db.Where("name = @name And age = @age", map[string]interface{}{"name": "jinzhu","age":20}).Debug().First(&user)

	// ========================================================================
	// FirstOrInit
	// db.FirstOrInit(&user, User{Name: "non_existing"})
	// db.Where(User{Name: "Jinzhu"}).FirstOrInit(&user)
	// db.Debug().FirstOrInit(&user, map[string]interface{}{"name": "av"})

	// Attrs
	// db.Debug().Where(User{Name: "12"}).Attrs(User{Age: 12}).FirstOrInit(&user)

	// Assign
	// db.Where(User{Name: "non_existing"}).Assign(User{Age: 12}).FirstOrInit(&user)
	// db.Where(User{Name: "Jinzhu"}).Assign(User{Age: 12}).FirstOrInit(&user)

	// ========================================================================
	// FirstOrCreate
	// db.FirstOrCreate(&user, User{Name: "non_existing"})

	// Attrs
	// db.Where(User{Name: "non_existing"}).Attrs(User{Age: 12}).FirstOrCreate(&user)

	// Assign
	// db.Where(User{Name: "non_existing"}).Assign(User{Age: 12}).FirstOrCreate(&user)

	// ========================================================================
	// Optimizer Hints
	// db.Clauses(hints.New("MAX_EXECUTION_TIME(10000)")).Find(&user)

	// ========================================================================
	// Index Hints
	// db.Clauses(hints.UseIndex("idx_user_name")).Find(&user)

	// ========================================================================
	// Iteration

	// rows, _ := db.Model(&User{}).Where("name = ?", "jinzhu").Rows()
	//
	// defer rows.Close()
	//
	// for rows.Next() {
	//     var user User
	//     // ScanRows is a method of `gorm.DB`, it can be used to scan a row into a struct
	//     db.ScanRows(rows, &user)
	//     fmt.Printf("%+v\n", user)
	//     // do something
	// }

	// ========================================================================
	// FindInBatches
	//
	// batch size 100
	// var results []User
	// var count int = 0
	// result := db.FindInBatches(&results, 100, func(tx *gorm.DB, batch int) error {
	//     count++
	//     for index, _ := range results {
	//         // batch processing found records
	//         fmt.Printf("%+v\n",results[index].Name)
	//         // results[index].Age = 60
	//     }
	//
	//     // tx.Save(&results)
	//
	//     // fmt.Println("RowsAffected: ", tx.RowsAffected) // number of records in this batch
	//
	//     // fmt.Println("batch: ", batch) // Batch 1, 2, 3
	//
	//     // returns error will stop future batches
	//     return nil
	// })
	// fmt.Println(count)
	// fmt.Printf("back：%v\n",result.Error)
	// fmt.Printf("%+v\n",result.RowsAffected)
	// result.Error // returned error
	// result.RowsAffected // processed records count in all batches
	// fmt.Printf("%+v",user)
	// if cap(user) == 0 {
	//     fmt.Printf("%+v\n", user)
	// } else {
	//     for i, data := range user {
	//         fmt.Printf("%d  %v\n", i, data)
	//     }
	// }

	// ========================================================================
	// Pluck
	// var ages []int64
	// var users []map [string] interface{}
	// var users []User
	// db.Model(&User{}).Pluck("age", &ages)

	// var names []string
	// db.Model(&User{}).Pluck("name", &names)

	// db.Table("deleted_users").Pluck("name", &names)
	//
	// // Distinct Pluck
	// db.Model(&User{}).Distinct().Pluck("Name", &names)
	// // SELECT DISTINCT `name` FROM `users`

	//
	// // Requesting more than one column, use `Scan` or `Find` like this:
	// db.Model(&users).Select("name", "age").Scan(&users)
	// db.Select("name", "age").Find(&users)
	// fmt.Println(users)

	// ========================================================================
	// Scopes
	// var users []User
	// db.Scopes(AmountGreaterThan1000, PaidWithCreditCard).Find(&users)
	// Find all credit card orders and amount greater than 1000

	// db.Scopes(AmountGreaterThan1000, PaidWithCod).Find(&users)
	// Find all COD orders and amount greater than 1000

	// db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"jinzhu", "fe"})).Find(&users)
	// Find all paid, shipped orders that amount greater than 1000

	// for i, data := range users {
	//     fmt.Printf("%d  %v\n", i, data)
	// }

	// ========================================================================
	// Count
	// var count int64
	// db.Model(&User{}).Where("name = ?", "jinzhu").Or("name = ?", "fe").Count(&count)
	// SELECT count(1) FROM users WHERE name = 'jinzhu' OR name = 'jinzhu 2'

	// db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
	// SELECT count(1) FROM users WHERE name = 'jinzhu'; (count)

	// db.Table("users").Count(&count)
	// SELECT count(1) FROM deleted_users;

	// Count with Distinct
	// db.Model(&User{}).Distinct("name").Count(&count)
	// SELECT COUNT(DISTINCT(`name`)) FROM `users`

	// db.Table("users").Select("count(distinct(name))").Count(&count)
	// SELECT count(distinct(name)) FROM deleted_users

	// Count with Group
	// users := []User{
	//     {Name: "name1"},
	//     {Name: "name2"},
	//     {Name: "name3"},
	//     {Name: "name3"},
	// }

	// db.Model(&User{}).Group("name").Count(&count)
	// count // => 3
	// fmt.Println(count)
}

// func (u *User) AfterFind(tx *gorm.DB) error {
//     if u.Name == "" {
//         // u.as = "user"
//     }
//
//     return errors.New("test")
// }

// func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
//     return db.Where("age > ?", 20)
// }
// func PaidWithCreditCard(db *gorm.DB) *gorm.DB {
//     return db.Where("name = ?", "Jinzhu")
// }
//
// func PaidWithCod(db *gorm.DB) *gorm.DB {
//     return db.Where("name = ?", "fe")
// }
//
// func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
//     return func (db *gorm.DB) *gorm.DB {
//         return db.Where("name IN (?)", status)
//     }
// }
