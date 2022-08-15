package usecases


import (
    "fmt"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/mahani-software-engineering/bms-server/db"
)

func CreateDoctor(w http.ResponseWriter, r *http.Request) {
    var doctor db.Doctor
    _ = json.NewDecoder(r.Body).Decode(&doctor)
    database.Create(&doctor)
    msg := fmt.Sprintf("New doctor recorded")
    respondToClient(w, 201, doctor, msg)
}

func UpdateDoctor(w http.ResponseWriter, r *http.Request) {
    //TODO: implement
}

func doctorExists (identifier uint) (bool, db.Doctor, error) {
    //the identifier can be ID, phone, email, username
    var doctor db.Doctor
    response := database.Where("id = ?", identifier).First(&doctor)                   
    numberOfRowsFound := response.RowsAffected
    exists := numberOfRowsFound > 0
    return exists, doctor, response.Error
}

func ReadDoctor(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)
    identf := params["id"]
    if identifier, err := strconv.Atoi(identf); err == nil {
        ok, doctor, err := doctorExists(uint(identifier))
        if err != nil {
            respondToClient(w, 400, nil, err.Error())
        }
        
        if !ok {
            respondToClient(w, 404, nil, "Specified doctor record not found")
        }
        
        respondToClient(w, 200, doctor, "")
    }else{
        respondToClient(w, 403, nil, "Invalid doctor identifier")
    }
}

func ReadAllDoctors(w http.ResponseWriter, r *http.Request) {
    var doctors []db.Doctor
    response := database.Find(&doctors)
    numberOfRowsFound := response.RowsAffected
    msg := fmt.Sprintf("Found %d doctors", numberOfRowsFound)
    respondToClient(w, 200, doctors, msg)
}



