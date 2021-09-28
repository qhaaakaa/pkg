package migrations

import (
	"database/sql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golang-migrate/migrate/source/file"
	"log"
)

func main() {
	migrateInit()
}

func migrateInit() {
	db, err := sql.Open("mysql", "root:qh123456@tcp(127.0.0.1:3306)/gorm?multiStatements=true")
	if err != nil {
		log.Fatal(err)
		return
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	fileDir, err := (&file.File{}).Open("file://.")
	if err != nil {
		log.Fatal(err)
		return
	}
	m, err := migrate.NewWithInstance(
		"file",
		fileDir,
		"mysql",
		driver,
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	//向下
	err = m.Down()
	if err != nil {
		log.Fatal(err)
		return
	}
	//设置版本
	//m.Force(1)
	//向上
	err = m.Up()
	if err != nil {
		log.Fatal(err)
		return
	}
}
