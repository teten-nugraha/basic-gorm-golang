package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserL struct {
	ID int `gorm:"primary_key"`
	Uname string
	Languages []Language `gorm:"many2many:user_languages";"ForeignKey:UserId"`
}

type Language struct {
	ID int `gorm:"primary_key"`
	Name string
}

type UserLanguages struct {
	ID int `gorm:"primary_key"`
	Name string
}

func main () {
	db, err := gorm.Open("mysql", "root:root12345@tcp(127.0.0.1:3306)/ormdemo?charset=utf8&parseTime=True")
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to Open")
	}
	log.Println("Connection Established")

	defer db.Close()
	db.DropTableIfExists(&UserLanguages{}, &Language{}, &UserL{})
	db.AutoMigrate(&UserL{}, &Language{}, &UserLanguages{})
	
	//All foreign keys need to define here
	db.Model(UserLanguages{}).AddForeignKey("user_l_id", "user_ls(id)", "CASCADE", "CASCADE")
	db.Model(UserLanguages{}).AddForeignKey("language_id", "languages(id)", "CASCADE", "CASCADE")
	
	langs := []Language{{Name:"English"},{Name:"French"}}

	user1 := UserL{Uname:"Jhon", Languages: langs}
	user2 := UserL{Uname:"Martin", Languages: langs}
	user3 := UserL{Uname:"Ray", Languages: langs}

	db.Save(&user1)
	db.Save(&user2)
	db.Save(&user3)

	fmt.Println("After saving records")
	fmt.Println("User1 ", &user1)
	fmt.Println("User2 ", &user2)
	fmt.Println("User3 ", &user3)

	// fetching
	user := &UserL{}
	db.Debug().Where("uname=?","Ray").Find(&user)
	err = db.Debug().Model(&user).Association("Languages").Find(&user.Languages).Error
	fmt.Println("User is now coming ", user, err)
}