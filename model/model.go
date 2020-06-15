package model

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type Diary struct {
    gorm.Model
    Like    Like   `gorm:"foreignkey:ID;association_foreignkey:DiaryId"`
    UserID  uint   `gorm:column:user_id`
    Content string `gorm:column:content`
}

type Like struct {
    gorm.Model
    UserID   uint  `gorm:column:user_id`
    DiaryID  uint  `gorm:column:diary_id`
}

type User struct {
    gorm.Model
    Diary            Diary            `gorm:"foreignkey:ID;association_foreignkey:UserId"`
    Like             Like             `gorm:"foreignkey:ID;association_foreignkey:UserId"`
    FollowerRequest  FollowerRequest  `gorm:"foreignkey:ID;association_foreignkey:UserId"`
    Name             string           `gorm:"unique;not null`
    Password         string           `gorm:"unique;not null"`
    Token            int              `gorm:column:token`
}

type FollowerRequest struct {
    gorm.Model
    UserID        uint  `gorm:column:user_id`
    TargetUserID  uint  `gorm:column:target_user_id`
    Allow         bool  `gorm:column:allow`
}
