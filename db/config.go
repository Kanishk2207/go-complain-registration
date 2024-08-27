// db/config.go
package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
    dsn := "mysql:abc@tcp(0.0.0.0:3306)/bugsmirror" 
    return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
