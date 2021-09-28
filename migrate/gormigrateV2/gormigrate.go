package main

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main()  {
	gormInit()
}

func gormInit() {
	db, err := gorm.Open(mysql.Open("root:qh123456@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		// create persons table
		{
			ID: "201608301400",
			Migrate: func(tx *gorm.DB) error {
				// it's a good pratice to copy the struct inside the function,
				// so side effects are prevented if the original struct changes during the time
				type Person struct {
					gorm.Model
					Name string `gorm:"type:varchar(50);default name;not null"`
				}
				return tx.AutoMigrate(&Person{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("people")
			},
		},
		// add age column to persons
		{
			ID: "201608301415",
			Migrate: func(tx *gorm.DB) error {
				// when table already exists, it just adds fields as columns
				type Person struct {
					Age int `gorm:"type:int(10);default 1;not null"`
				}
				return tx.AutoMigrate(&Person{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropColumn("people", "age")
			},
		},
		// add pets table
		{
			ID: "201608301430",
			Migrate: func(tx *gorm.DB) error {
				type Pet struct {
					gorm.Model
					Name     string
					PersonID int
				}
				return tx.AutoMigrate(&Pet{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("pets")
			},
		},
		{
			ID: "201608301431",
			Migrate: func(tx *gorm.DB) error {
				err := tx.Exec(`CREATE TABLE user (
					  id int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
					  username varchar(30) NOT NULL COMMENT '账号',
					  password varchar(100) NOT NULL COMMENT '密码',
					  createtime int(10) NOT NULL DEFAULT '0' COMMENT '创建时间',
					  PRIMARY KEY (id)
					) ENGINE=InnoDB DEFAULT CHARSET=utf8;`).Error
				return err
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("pets")
			},
		},
	})

	//回滚
	if err = m.RollbackLast(); err != nil {
		log.Fatalf("Could not RollbackLast: %v", err)
	}
	//迁移到某个版本
	//if err = m.MigrateTo("201608301430"); err != nil {
	//	log.Fatalf("Could not migrate: %v", err)
	//}
	//回滚到某个版本
	//if err = m.RollbackTo("201608301415"); err != nil {
	//	log.Fatalf("Could not migrate: %v", err)
	//}

	//迁移
	//if err = m.Migrate(); err != nil {
	//	log.Fatalf("Could not migrate: %v", err)
	//}

	log.Printf("Migration did run successfully")
}
