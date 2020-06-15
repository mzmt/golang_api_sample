package model

import (
    "log"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

// TODO env設定

func InitializeDB() {
    db := ConnectDB()
    MigrateDB(db)
}

func ConnectDB() *gorm.DB {
    db, err := gorm.Open("postgres", "dbname=golang_api sslmode=disable")
    if err != nil {
        panic(err.Error())
    }
    return db
}

func MigrateDB(db *gorm.DB) *gorm.DB {
    log.Println("migration start")
    db.AutoMigrate(&User{}, &Diary{})
    // db.Model(&Diary{}).AddForeignKey("diary_id", "projects(id)", "CASCADE", "CASCADE")
    return db
}
