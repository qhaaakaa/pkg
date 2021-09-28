##gorMigrate 数据迁移

### 基于gorm框架的扩展，在原有基础上增加了回滚，版本控制等功能
```go
package main

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
    
	// 创建gormigrate
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
	})

	//回滚
	//if err = m.RollbackLast(); err != nil {
	//	log.Fatalf("Could not RollbackLast: %v", err)
	//}

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
```