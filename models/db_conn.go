
package models


import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"fmt"
)



func InitDb() *gorm.DB{
	db, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil{
		fmt.Println("Error initializing db connection")
		panic(err)
	}
	//defer db.Close()
	return db
}