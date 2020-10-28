package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	Age          uint8
	Birthday     time.Time
}


func InsertUser(){
	user := User{Name: "velan", Age: 28, Birthday: time.Now()}

	db, err := gorm.Open(mysql.Open("root:Siar@123@tcp(localhost:3306)/velanormusers?"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&User{})
	if err != nil{
		return
	}
	result := db.Create(&user)

	fmt.Println(user.ID)
	if result.Error != nil{
		fmt.Println("err")
		return
	}

	fmt.Println(result.RowsAffected)
	//user.ID             // returns inserted data's primary key
	//result.Error        // returns error
	//result.RowsAffected // returns inserted records count

}

func InsertWithSelectField(){
	db, err := gorm.Open(mysql.Open("root:Siar@123@tcp(localhost:3306)/velanormusers?"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&User{})
	if err != nil{
		return
	}
	user := User{Email: "velann21@gmail.com"}
	result := db.Select("Email").Create(&user)
    fmt.Println(result)


}
func main() {
	InsertUser()
	InsertWithSelectField()
}






