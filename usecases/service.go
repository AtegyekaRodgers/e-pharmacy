package usecases


import (
    "fmt"
    "strconv"
    "net/http"
    //"encoding/json"
    "github.com/gorilla/mux"
    "github.com/mahani-software-engineering/bms-server/db"
)



func CreateService(w http.ResponseWriter, r *http.Request) {

}

func UpdateService(w http.ResponseWriter, r *http.Request) {

}

func serviceExists (identifier uint) (bool, db.Service, error) {
    //the identifier can be ID, phone, email, username
    var service db.Service
    response := database.Where("id = ?", identifier).First(&service)                   
    numberOfRowsFound := response.RowsAffected
    exists := numberOfRowsFound > 0
    return exists, service, response.Error
}

func ReadService(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)
    identf := params["id"]
    if identifier, err := strconv.Atoi(identf); err == nil {
        ok, service, err := serviceExists(uint(identifier))
        if err != nil {
            respondToClient(w, 400, nil, err.Error())
        }
        
        if !ok {
            respondToClient(w, 404, nil, "Specified service record not found")
        }
        
        respondToClient(w, 200, service, "")
    }else{
        respondToClient(w, 403, nil, "Invalid service identifier")
    }
}

func ReadAllServices(w http.ResponseWriter, r *http.Request) {
    var services []db.Service
    response := database.Find(&services)
    numberOfRowsFound := response.RowsAffected
    msg := fmt.Sprintf("Found %d records", numberOfRowsFound)
    respondToClient(w, 200, services, msg)
}

