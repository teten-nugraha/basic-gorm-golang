package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Place struct {
	ID     int `gorm:"primary_key"`
	Name   string
	Town   Town
	TownId int `gorm:"ForeignKey:id"`
}

type Town struct {
	ID   int `gorm:"primary_key"`
	Name string
}

func main() {
	db, err := gorm.Open("mysql", "root:root12345@tcp(127.0.0.1:3306)/ormdemo?charset=utf8&parseTime=True")
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to Open")
	}
	log.Println("Connection Established")

	db.DropTableIfExists(&Place{}, &Town{})
	db.AutoMigrate(&Place{}, &Town{})
	db.Model(&Place{}).AddForeignKey("town_id", "towns(id)", "CASCADE", "CASCADE")

	t1 := Town{
		Name: "Pune",
	}
	t2 := Town{
		Name: "Mumbai",
	}
	t3 := Town{
		Name: "Hyderabad",
	}

	p1 := Place{
		Name: "Katraj",
		Town: t1,
	}
	p2 := Place{
		Name: "Thane",
		Town: t2,
	}
	p3 := Place{
		Name: "Secundarabad",
		Town: t3,
	}

	db.Save(&p1) //Saving one to one relationship
	db.Save(&p2)
	db.Save(&p3)

	fmt.Println("t1==>", t1, " p1==>", p1)
	fmt.Println("t2==>", t2, " p2s==>", p2)
	fmt.Println("t2==>", t3, " p2s==>", p3)

	// delete
	// db.Where("name=?", "Hyderabad").Delete(&Town{})

	// select
	places := Place{}
	towns := Town{}
	fmt.Println("Before Association ", places)
	db.Where("name=?", "Katraj").Find(&places)
	fmt.Println("After Association ", places)
	err = db.Model(&places).Association("town").Find(&places.Town).Error
	fmt.Println("After Association", towns, places)
	fmt.Println("After Association", towns, places, err)

	defer db.Close()
}
