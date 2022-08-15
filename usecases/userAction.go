package usecases


import (
    "net/http"
    "github.com/mahani-software-engineering/bms-server/db"
)

func newActionRecord(uid uint, an, descr, entity, option string) {
    var action db.UserAction
    action.ActionNumber = an
    action.OnEntity = entity
    action.SpecificEntity = option
    action.Description = descr
    if(uid > 0){
      action.UserID = uid
      database.Create(&action)
    }else{
      database.Omit("UserID").Create(&action)
    }
}


func ReadUserAction(w http.ResponseWriter, r *http.Request) {

}

func ReadAllUserActions(w http.ResponseWriter, r *http.Request) {

}


