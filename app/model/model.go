package model

import (
)

type Diary struct {
    Content string `gorm:column:content`
}

type Like struct {
}

type User struct {
    Name string `gorm:column:name`
}
