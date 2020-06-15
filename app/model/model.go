package model

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Diary struct {
    gorm.Model
    Content string `gorm:column:content`
}

type Like struct {
    gorm.Model
}

type User struct {
    gorm.Model
    Name    string  `gorm:column:name`
    Pasword string  `gorm:column:password`
    Token   string  `gorm:column:token`
}
