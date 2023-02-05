package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connetion DataBases
func DatabaseInit() {
	var err error
	// dsn := "root:@tcp(localhost:3306)/dumbmerch?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:m2NvLsZSp1iw38USN8Jj@tcp(containers-us-west-144.railway.app:7097)/railway?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Conneted To Databases")
}
