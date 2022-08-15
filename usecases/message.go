package usecases


import (
    "fmt"
    "time"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/mahani-software-engineering/bms-server/db"
)

func CreateMessage(w http.ResponseWriter, r *http.Request) {
    var message db.Message
    _ = json.NewDecoder(r.Body).Decode(&message)
    database.Create(&message)
    msg := fmt.Sprintf("Message sent")
    respondToClient(w, 201, message, msg)
}

func UpdateMessage(w http.ResponseWriter, r *http.Request) {
    //TODO: implement
}

func messageExists (identifier uint) (bool, db.Message, error) {
    //the identifier must be ID
    var message db.Message
    response := database.Where("id = ?", identifier).First(&message)                   
    numberOfRowsFound := response.RowsAffected
    exists := numberOfRowsFound > 0
    return exists, message, response.Error
}

func ReadMessage(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    identf := params["id"]
    if identifier, err := strconv.Atoi(identf); err == nil {
        ok, message, err := messageExists(uint(identifier))
        if err != nil {
            respondToClient(w, 400, nil, err.Error())
        }
        
        if !ok {
            respondToClient(w, 404, nil, "Specified message record not found")
        }
        
        respondToClient(w, 200, message, "")
    }else{
        respondToClient(w, 403, nil, "Invalid message identifier")
    }
}

func ReadAllMessages(w http.ResponseWriter, r *http.Request) {
    var messages []db.Message
    response := database.Find(&messages)
    numberOfRowsFound := response.RowsAffected
    msg := fmt.Sprintf("Found %d messages", numberOfRowsFound)
    respondToClient(w, 200, messages, msg)
}

func ReadMyMessages(w http.ResponseWriter, r *http.Request) {
    messages := []struct{ID uint `json:"id"`; DateTime time.Time `json:"date_time"`; Text string `json:"text"`; Fromm string `json:"fromm"`;}{}
    params := mux.Vars(r)
    identf := params["id"]
    fmt.Println("Read messages for user ID:",identf)
    if identifier, err := strconv.Atoi(identf); err == nil {
        response := database.Table("messages as m").Where("user_id = ?", identifier).Joins("JOIN users as u on u.id=m.created_by").Select(`m.id as id, m.created_at as date_time, m.text as text, CONCAT(u.firstname, " ",u.lastname) as fromm`).Find(&messages)
        numberOfRowsFound := response.RowsAffected
        msg := fmt.Sprintf("Found %d messages", numberOfRowsFound)
        respondToClient(w, 200, messages, msg)
    }else{
        respondToClient(w, 403, nil, "Invalid user identifier. System cannot read messages for anonymous user.")
    }  
}




