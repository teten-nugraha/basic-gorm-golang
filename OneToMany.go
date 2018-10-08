package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Customer struct {
	CustomerID   int `gorm:"primary_key"`
	CustomerName string
	Contacts     []Contact `gorm:"ForeignKey:CustID"`
}

type Contact struct {
	ContactID   int `gorm:"primary_key"`
	CountryCode int
	MobileNo    uint
	CustId      int
}

func main() {
	db, err := gorm.Open("mysql", "root:root12345@tcp(127.0.0.1:3306)/ormdemo?charset=utf8&parseTime=True")
	defer db.Close()
	if err != nil {
		log.Println("Connection Failed to Open")
	}
	log.Println("Connection Established")

	db.DropTableIfExists(&Customer{}, &Contact{})
	db.AutoMigrate(&Customer{}, &Contact{})
	db.Model(&Contact{}).AddForeignKey("cust_id", "customers(customer_id)", "CASCADE", "CASCADE")

	Custs1 := Customer{CustomerName: "John", Contacts: []Contact{
		{CountryCode: 91, MobileNo: 956112},
		{CountryCode: 91, MobileNo: 997555}}}

	Custs2 := Customer{CustomerName: "Martin", Contacts: []Contact{
		{CountryCode: 90, MobileNo: 808988},
		{CountryCode: 90, MobileNo: 909699}}}

	Custs3 := Customer{CustomerName: "Raym", Contacts: []Contact{
		{CountryCode: 75, MobileNo: 798088},
		{CountryCode: 75, MobileNo: 965755}}}

	Custs4 := Customer{CustomerName: "Stoke", Contacts: []Contact{
		{CountryCode: 80, MobileNo: 805510},
		{CountryCode: 80, MobileNo: 758863}}}

	db.Create(&Custs1)
	db.Create(&Custs2)
	db.Create(&Custs3)
	db.Create(&Custs4)

	customers := &Customer{}
	contacts := &Contact{}

	db.Debug().Where("customer_name=?", "Martin").Preload("Contacts").Find(&customers) //db.Debug().Where("customer_name=?","John").Preload("Contacts").Find(&customers)
	fmt.Println("Customers", customers)
	fmt.Println("Contacts", contacts)

}
