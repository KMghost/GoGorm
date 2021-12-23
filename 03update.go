package main

import (
// driver "GoGorm/database"
// "errors"
// "fmt"
// "golang.org/x/crypto/bcrypt"
// "gorm.io/gorm"
// "time"
)

// type User driver.User

func main() {
	// db := driver.DB()

	// save
	// var user User
	// db.First(&user)
	//
	// user.Name = "jinzhu 2"
	// // user.Age = 100
	// // *user.Birthday = time.Now()
	// db.Save(&user)

	// ========================================================================
	// Update single column
	// 条件更新
	// var user User
	// db.Model(&User{}).Where("age = ?", 5).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

	// User 的 ID 是 `111`
	// user.ID = 35
	// db.Model(&user).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

	// 根据条件和 model 的值进行更新
	// user.ID  = 55
	// db.Model(&user).Where("age = ?", 60).Update("name", "hello")
	// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;

	// ========================================================================
	// Updates multiple columns
	// var user User
	// 根据 `struct` 更新属性，只会更新非零值的字段
	// user.ID = 35
	// db.Model(&user).Select("Name","Age","Active").Updates(User{Name: "hello", Age: 18, Active: false})
	// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

	// 根据 `map` 更新属性
	// user.ID = 35
	// db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	// ========================================================================
	// Update Selected Fields
	// var user User
	// 使用 Map 进行 Select
	// User's ID is `111`:
	// db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET name='hello' WHERE id=111;
	// user.ID = 36
	// db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
	// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

	// 使用 Struct 进行 Select（会 select 零值的字段）
	// user.ID = 37
	// db.Model(&user).Select("Name", "Age").Updates(User{Name: "jinzhu", Age: 234})
	// UPDATE users SET name='new_name', age=0 WHERE id=111;

	// Select 所有字段（查询包括零值字段的所有字段）
	// user.ID = 36
	// db.Model(&user).Select("*").Updates(User{Name: "jinzhu1", Age: 12, Active:false,ID: 36})

	// Select 除 Role 外的所有字段（包括零值字段的所有字段）
	// user.ID = 36
	// db.Model(&user).Select("*").Omit("ID").Updates(User{Name: "jinzhu", Age: 0})

	// ========================================================================
	// Update Hook

	// ========================================================================
	// Batch Updates
	// Update with struct
	// db.Model(User{}).Where("name = ?", "jinzhu").Updates(User{Name: "jinzhu2", Age: 18})
	// UPDATE users SET name='hello', age=18 WHERE role = 'admin';

	// Update with map
	// db.Table("users").Where("id IN ?", []int{35, 36}).Updates(map[string]interface{}{"name": "hello", "age": 18})
	// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);

	// ========================================================================
	// Block Global Updates
	// a := db.Model(&User{}).Update("name", "jinzhu").Error // gorm.ErrMissingWhereClause
	// fmt.Println(a)

	// db.Model(&User{}).Where("1 = 1").Update("name", "jinzhu")
	// UPDATE users SET `name` = "jinzhu" WHERE 1=1

	// db.Exec("UPDATE users SET name = ?", "jinzhu")
	// UPDATE users SET name = "jinzhu"

	// db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).Update("name", "jinzhu")
	// UPDATE users SET `name` = "jinzhu"

	// ========================================================================
	// SQL Expression
	// product 的 ID 是 `3`
	// var user User
	// db.Model(&user).Update("price", gorm.Expr("price * ? + ?", 2, 100))
	// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;

	// db.Model(&user).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})
	// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;

	// db.Model(&user).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
	// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = 3;

	// db.Model(&user).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
	// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = 3 AND quantity > 1;

	// ========================================================================
	// Returning Data From Modified Rows
	// var users []User
	// db.Debug().Model(&users).Clauses(clause.Returning{}).Where("name = ?", "jinzhu").Update(
	//     "age", gorm.Expr("age * ?", 2))
	// fmt.Printf("%+v",users)

	// db.Model(&users).Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"}, {Name: "salary"}}}).Where(
	//     "name = ?", "jinzhu").Update("age", gorm.Expr("age * ?", 2))

	// fmt.Printf("%+v",users)
	// result := db.Model(&User{ID:36,Name: "jinzhu2"}).Debug().Updates(map[string]interface{}{"name": "jinzhu1"})

	// result := db.Model(&User{ID: 36, Name: "jinzhu"}).Debug().Omit("Name").Updates(map[string]interface{}{
	//     "Name": "jinzhu2", "Age": 20,
	// })
	// fmt.Printf("%+v",result.Error)
}

// func (u *User) AfterUpdate(tx *gorm.DB) (err error) {
//     fmt.Printf("name: %+v",u)
//     if u.Name == "jinzhu" {
//         return errors.New("admin user not allowed to update")
//     }
//     return
// }

// func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
//     // if Role changed
//     if tx.Statement.Changed("Name") {
//         fmt.Println(1)
//         return errors.New("name not allowed to change")
//     }
//
//     if tx.Statement.Changed("Age") { // if Name or Role changed
//         fmt.Println(2)
//         tx.Statement.SetColumn("Age", 18)
//     }
//
//     // if any fields changed
//     if tx.Statement.Changed() {
//         fmt.Println(3)
//         tx.Statement.SetColumn("birthday", time.Now())
//     }
//     return nil
// }
