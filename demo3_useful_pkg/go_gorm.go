package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	UserID   int       `gorm: "column: user_id"`
	UserName string    `gorm: "column: user_name"`
	Age      int       `gorm: "column: age"`
	Role     string    `gorm: "column: role"`
	Birthday time.Time `gorm: "column: birthday"`
}

type AgeCount struct {
	Age   int   `json:"age"`
	Total int64 `json:"total"`
}

/**
type User struct {
    UserID   int       `json: "user_id"`
    UserName string    `json: "user_name"`
    Age      int       `json: "age"`
    Birthday time.Time `json: "birthday"`
}
*/

func openDB() (db *gorm.DB, err error) {
	return gorm.Open("mysql", "root:R0otAwc10ud@tcp(172.16.2.100:8090)/test?charset=utf8&parseTime=True&loc=Local")
}

func testCreateTable() {
	db, err := openDB()
	if err != nil {
		fmt.Println("Open db error", err)
		return
	}
	defer db.Close()

	// db.LogMode(true)

	fmt.Println("Init.")
	ret := db.HasTable(&Product{})
	fmt.Println("->	Has table with Model Product result:", ret)

	fmt.Println("Create table With Model Product.")
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Product{})

	ret = db.HasTable(&Product{})
	fmt.Println("->	Has table With Model Product result:", ret)

	ret = db.HasTable("product")
	fmt.Println("->	Has table with name product result:", ret)

	ret = db.HasTable("products")
	fmt.Println("->	Has table with name products result:", ret)

	fmt.Println("Drop table with Model Product.")
	db.DropTable(&Product{})

	ret = db.HasTable(&Product{})
	fmt.Println("->	Has table with Model Product result:", ret)

	fmt.Println("Set SingularTable true.")
	db.SingularTable(true)

	fmt.Println("Create table With Model Product.")
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Product{})

	ret = db.HasTable(&Product{})
	fmt.Println("->	Has table With Model Product result:", ret)

	ret = db.HasTable("product")
	fmt.Println("->	Has table with name product result:", ret)

	ret = db.HasTable("products")
	fmt.Println("->	Has table with name products result:", ret)

	fmt.Println("Drop table with Model Product.")
	db.DropTable(&Product{})

	ret = db.HasTable(&Product{})
	fmt.Println("->	Has table with Model Product result:", ret)
}

func testQueryData() {
	db, err := openDB()
	if err != nil {
		fmt.Println("Open mysql error,", err)
		return
	}
	db.SingularTable(true)
	// db.LogMode(true)

	// 表和数据已经手动创建好
	// user_id	user_name	age role 	birthday
	// 1		Dinghh		20	admin	2019-07-15 19:05:39
	// 2		ChenJianhua	19	user	2019-07-15 19:05:39
	// 3		Weike		19	user	2019-07-16 09:29:11
	// 4		ZhangXu		18	user	2019-07-16 09:29:27
	// 5		Yance		19	user	2019-07-16 09:29:37
	user1 := User{UserName: "Aaron", Age: 18, Role: "admin", Birthday: time.Now()}
	db.Create(&user1)

	// Select ALL
	fmt.Println("Get all users.")
	var users_all []User
	db.Find(&users_all) // SELECT * FROM users;
	fmt.Println("->	model result:", users_all)
	users_all_json, _ := json.Marshal(users_all)
	fmt.Println("->	json result:", string(users_all_json))
	fmt.Println("------------------------------------")

	// Select Condition
	fmt.Println("Get user with id")
	var user User
	db.Where("user_id = ?", 1).Find(&user) // SELECT * FROM users WHERE user_id = 10;
	user_json, _ := json.Marshal(user)
	fmt.Println("->	json result:", string(user_json))
	fmt.Println("------------------------------------")

	// In
	fmt.Println("Get users with multi names")
	var users_in []User
	db.Where("user_name in (?)", []string{"Dinghh", "Yance"}).Find(&users_in)
	users_in_json, _ := json.Marshal(users_in)
	fmt.Println("->	json result:", string(users_in_json))
	fmt.Println("------------------------------------")

	// Order
	fmt.Println("Get users order by age")
	var users_order []User
	db.Where("role = ?", "user").Order("age desc").Find(&users_order)
	users_order_json, _ := json.Marshal(users_order)
	fmt.Println("->	json result:", string(users_order_json))
	fmt.Println("------------------------------------")

	// Limit and offset
	fmt.Println("Get users with limit and offset")
	var users_limit []User
	db.Where("role = ?", "user").Order("age asc").Limit(2).Offset(1).Find(&users_limit)
	users_limit_json, _ := json.Marshal(users_limit)
	fmt.Println("-> json result:", string(users_limit_json))
	fmt.Println("------------------------------------")

	// Count
	fmt.Println("Get users count")
	var users_count []User
	var count int
	db.Find(&users_count).Count(&count)
	fmt.Println("->	All users count:", count)
	fmt.Println("------------------------------------")

	var users_userrole_count []User
	var userrole_count int
	db.Where("role = ?", "user").Find(&users_userrole_count).Count(&userrole_count)
	fmt.Println("->	User role count:", userrole_count)
	fmt.Println("------------------------------------")

	// Group
	fmt.Println("Get all age count")
	var ageCounts []AgeCount
	db.Table("user").Select("age, count(age) as total").Group("age").Having("count(age) > ?", 1).Scan(&ageCounts)
	ageCounts_json, _ := json.Marshal(ageCounts)
	fmt.Println("->	json result:", string(ageCounts_json))

	// Not Equal
	// db.Where("user_name <> ?", "jinzhu").Find(&users)

	// LIKE
	// db.Where("user_name LIKE ?", "%jin%").Find(&users)

	// AND
	// db.Where("user_name = ? AND age >= ?", "jinzhu", "22").Find(&users)

	// Time
	// db.Where("birthday > ?", lastWeek).Find(&users)
	// db.Where("birthday BETWEEN ? AND ?", lastWeek, today).Find(&users)

	// Or
	// db.Where("user_name = ?", "admin").Or("age = ?", 18).Find(&users)

	// Muti Query
	// db.Where("user_name <> ?","jinzhu").Where("age >= ? and role <> ?",20,"admin").Find(&users)
	// db.Where("role = ?", "admin").Or("role = ?", "super_admin").Not("name = ?", "jinzhu").Find(&users)

	// Select
	fmt.Println("Select with field")
	var users_fields []User
	db.Where("user_name = ?", "Dinghh").Select("user_name, age").Find(&users_fields)
	users_fields_json, _ := json.Marshal(users_fields)
	fmt.Println("->	json result:", string(users_fields_json))

	db.Where("user_name = ?", "Aaron").Delete(User{})

	// Raw SQL
	fmt.Println("Execute raw sql")
	var users_raw []User
	db.Raw("select * from user where age = ?", 20).Scan(&users_raw)
	users_raw_json, _ := json.Marshal(users_raw)
	fmt.Println("->	json result:", string(users_raw_json))

	/**
	rows, _ := db.Raw("select user_name, age from users where age = ?", 18).Rows()
	defer rows.Close()
	for rows.Next() {
		var user_name1 string
		var age1 int
	    rows.Scan(&user_name1, &age1)
		fmt.Println("-> result: username, age", user_name1, age1)
	}
	*/

	// Transaction
	tx := db.Begin()
	user_add1 := User{UserID: 20, UserName: "Test1", Age: 21, Role: "admin", Birthday: time.Now()}
	user_add2 := User{UserID: 20, UserName: "Test2", Age: 21, Role: "admin", Birthday: time.Now()}

	if err1 := tx.Create(&user_add1).Error; err1 != nil {
		fmt.Println("Create user1 error", err1)
		tx.Rollback()
		return
	}
	// user_add2 will create failed since userid duplicate then rollback that means user_add1 will not created
	if err2 := tx.Create(&user_add2).Error; err2 != nil {
		fmt.Println("Create user2 error", err2)
		tx.Rollback()
		return
	}
	tx.Commit()
}

func main() {
	testCreateTable()
	fmt.Println("==============================")
	testQueryData()
}
