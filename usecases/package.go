package usecases


import (
    "fmt"
    "strconv"
    "net/http"
    //"encoding/json"
    "github.com/gorilla/mux"
    "github.com/mahani-software-engineering/bms-server/db"
)



func CreatePackage(w http.ResponseWriter, r *http.Request) {

}

func UpdatePackage(w http.ResponseWriter, r *http.Request) {

}

func packageExists (identifier uint) (bool, db.Package, error) {
    //the identifier can be ID, phone, email, username
    var pkage db.Package
    response := database.Where("id = ?", identifier).First(&pkage)                   
    numberOfRowsFound := response.RowsAffected
    exists := numberOfRowsFound > 0
    return exists, pkage, response.Error
}

func ReadPackage(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)
    identf := params["id"]
    if identifier, err := strconv.Atoi(identf); err == nil {
        ok, pkage, err := packageExists(uint(identifier))
        if err != nil {
            respondToClient(w, 400, nil, err.Error())
        }
        
        if !ok {
            respondToClient(w, 404, nil, "Specified package record not found")
        }
        
        respondToClient(w, 200, pkage, "")
    }else{
        respondToClient(w, 403, nil, "Invalid package identifier")
    }
}

func ReadAllPackages(w http.ResponseWriter, r *http.Request) {
    var packages []db.Package
    response := database.Find(&packages)
    numberOfRowsFound := response.RowsAffected
    msg := fmt.Sprintf("Found %d records", numberOfRowsFound)
    respondToClient(w, 200, packages, msg)
}
