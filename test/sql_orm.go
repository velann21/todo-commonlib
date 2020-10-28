package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Author struct {
	Name  string
	Email string
}

type Blog struct {
	ID      int
	Author  Author `gorm:"embedded;embeddedPrefix:author_"`
	Upvotes int32
}


func main() {
	db, err := gorm.Open(mysql.Open("root:Siar@123@tcp(localhost:3306)/velanormusers?"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Blog{})
	if err != nil{
		return
	}






	//// Migrate the schema
	//db.AutoMigrate(&Product{})
	//
	//// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//
	//// Read
	//var product Product
	//db.First(&product, 1) // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42
	//
	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - delete product
	//db.Delete(&product, 1)
}