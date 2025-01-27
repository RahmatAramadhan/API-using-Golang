package database

import (
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	dsn := "admin:admin123@tcp(127.0.0.1:3306)/mysql?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}
	fmt.Println("Koneksi ke database berhasil!")
}

func Migrate(models ...interface{}){
	err := DB.AutoMigrate(models...)
	if err != nil{
		log.Fatalf("Migrasi database berhasil!")
	}
}