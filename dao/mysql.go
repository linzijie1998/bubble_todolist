package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:123456@tcp(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Close() bool {
	err := DB.Close()
	if err != nil {
		panic(err)
		return false
	}
	return true
}
