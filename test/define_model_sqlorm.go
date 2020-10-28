package main

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NormalModel(){
	type User struct {
		ID           uint
		Name         string
		Email        *string
		Age          uint8
		Birthday     *time.Time
		MemberNumber sql.NullString
		ActivedAt    sql.NullTime
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}

	db, err := gorm.Open(mysql.Open("root:Siar@123@tcp(localhost:3306)/velanormusers?"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Blog{})
	if err != nil{
		return
	}

}

func ModelWithGoOrmModel(){
	type User struct {
		Name string
		gorm.Model
	}
	db, err := gorm.Open(mysql.Open("root:Siar@123@tcp(localhost:3306)/velanormusers?"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Blog{})
	if err != nil{
		return
	}

}

func ModelWithFieldLevelPermissions(){
	type User struct {
		Name string `gorm:"<-:create"` // allow read and create
		//Name string `gorm:"<-:update"` // allow read and update
		//Name string `gorm:"<-"`        // allow read and write (create and update)
		//Name string `gorm:"<-:false"`  // allow read, disable write permission
		//Name string `gorm:"->"`        // readonly (disable write permission unless it configured )
		//Name string `gorm:"->;<-:create"` // allow read and create
		//Name string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
		//Name string `gorm:"-"`  // ignore this field when write and read
	}

	db, err := gorm.Open(mysql.Open("root:Siar@123@tcp(localhost:3306)/velanormusers?"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Blog{})
	if err != nil{
		return
	}
}

func ModelWithAutoTimeUpdate(){
	type User struct {
		CreatedAt time.Time // Set to current time if it is zero on creating
		UpdatedAt int       // Set to current unix seconds on updaing or if it is zero on creating
		Updated   int64 `gorm:"autoUpdateTime:nano"` // Use unix nano seconds as updating time
		//Updated   int64 `gorm:"autoUpdateTime:milli"`// Use unix milli seconds as updating time
		Created   int64 `gorm:"autoCreateTime"`      // Use unix seconds as creating time
	}

	db, err := gorm.Open(mysql.Open("root:Siar@123@tcp(localhost:3306)/velanormusers?"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Blog{})
	if err != nil{
		return
	}
}

func ModelEmbeddedStruct() {
	type Author struct {
		Name  string
		Email string
	}

	type Blog struct {
		ID      int
		Author  Author `gorm:"embedded"`
		Upvotes int32
	}

	// equals
	//type Blog struct {
	//	ID    int64
	//	Name  string
	//	Email string
	//	Upvotes  int32
	//}

	db, err := gorm.Open(mysql.Open("root:Siar@123@tcp(localhost:3306)/velanormusers?"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Blog{})
	if err != nil{
		return
	}


}

func EmbeddedPrefixStruct(){
	type Blog struct {
		ID      int
		Author  Author `gorm:"embedded;embeddedPrefix:author_"`
		Upvotes int32
	}

	db, err := gorm.Open(mysql.Open("root:Siar@123@tcp(localhost:3306)/velanormusers?"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Blog{})
	if err != nil{
		return
	}

    dbb, err := db.DB()
	dbb.Close()


}