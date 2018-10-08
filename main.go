package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserModel struct {
	Id      int    `gorm:"primary_key";"AUTO_INCREMENT"`
	Name    string `gorm:"size:255"`
	Address string `gorm:"type:varchar(100)"`
}

func main() {
	db, err := gorm.Open("mysql", "root:root12345@tcp(127.0.0.1:3306)/ormdemo?charset=utf8&parseTime=True")
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to Open")
	}
	log.Println("Connection Established")

	// db.Debug().DropTableIfExists(&UserModel{})

	// db.Debug().AutoMigrate(&UserModel{})

	// //Create
	// user := &UserModel{Name: "Jhon", Address: "New York"}
	// db.Create(user)

	// var users []UserModel = []UserModel{
	// 	UserModel{Name: "Ricky", Address: "Sydney"},
	// 	UserModel{Name: "Adam", Address: "Brisbane"},
	// 	UserModel{Name: "Justin", Address: "California"},
	// }

	// for _, user := range users {
	// 	db.Create(&user)
	// }

	// Update
	// user1 := &UserModel{Name: "Jhon", Address: "New York"}
	// db.Find(&user1)
	// user1.Address = "Dago Pojok"
	// db.Save(&user1)

	user := &UserModel{}
	db.Where("name in (?)", []string{"Jhon", "Adam"}).Find(&user)
	fmt.Println(user)
}
