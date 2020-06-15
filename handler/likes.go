package handler

import (
    "strconv"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/mzmt/golang_api_sample/model"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func CreateLike(w http.ResponseWriter, r *http.Request) {
    if !auth(w, r) {
        return
    }
    db := model.ConnectDB()
    current_user := GetUserObject(w, r)
    if current_user == nil { return }
    diary_id_int, _ := strconv.Atoi(mux.Vars(r)["id"])
    diary_id := uint(diary_id_int)
    // current_user_id_int := int(current_user.ID)

    if isAuther(db, current_user.ID, diary_id_int) {
        respondError(w, 400, "Can not like to your diary")
        return
    }
    if isAlreadlyLike(db, current_user.ID, diary_id_int) {
        respondError(w, 400, "Already liked")
        return
    }

    like := model.Like{UserID: current_user.ID, DiaryID: diary_id}
    db.Create(&like)

    respondJSON(w, http.StatusOK, like)
}

func isAuther(db *gorm.DB, user_id uint, diary_id int) (bool) {
    diary := model.Diary{}
    db.Find(&diary, diary_id)
    if diary.UserID == user_id {
        return true
    } else {
        return false
    }
}

func isAlreadlyLike(db *gorm.DB, user_id uint, diary_id int) (bool) {
    like := model.Like{}
    if err := db.Where("user_id = ?", user_id).Where("diary_id = ?", diary_id).First(&like).Error; err != nil {
        return false
    }
    return true
}
