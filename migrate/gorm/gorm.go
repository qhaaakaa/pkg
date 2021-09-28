package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type Book struct {
	Id        string `gorm:"column:id; PRIMARY_KEY"`
	Title     string `gorm:"column:title;index:idx_title"`
	Content   string `gorm:"column:content"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	db, err := gorm.Open(mysql.Open("root:qh123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//创建迁移
	db.AutoMigrate(&Book{})

	// 可以通过Set设置附加参数，下面设置表的存储引擎为InnoDB
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Book{})

	//删除表
	//db.Migrator().DropTable(&Book{})
}