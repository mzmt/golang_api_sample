package model

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Diary struct {
    gorm.Model
    UserID  uint   `gorm:column:user_id`
    Content string `gorm:column:content`
}

type Like struct {
    gorm.Model
}

type User struct {
    gorm.Model
    Diary   Diary    `gorm:"foreignkey:ID;association_foreignkey:UserId"`
    Name    string   `gorm:column:name`
    Pasword string   `gorm:column:password`
    Token   int      `gorm:column:token`
}

type FollowerRequest struct {
    gorm.Model
}
