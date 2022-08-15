package usecases


import (
    "fmt"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/mahani-software-engineering/bms-server/db"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product db.Product
    _ = json.NewDecoder(r.Body).Decode(&product)
    database.Create(&product)
    msg := fmt.Sprintf("New product recorded")
    respondToClient(w, 201, product, msg)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

}

func productExists (identifier uint) (bool, db.Product, error) {
    //the identifier can be ID, phone, email, username
    var product db.Product
    response := database.Where("id = ?", identifier).First(&product)                   
    numberOfRowsFound := response.RowsAffected
    exists := numberOfRowsFound > 0
    return exists, product, response.Error
}

func ReadProduct(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)
    identf := params["id"]
    if identifier, err := strconv.Atoi(identf); err == nil {
        ok, product, err := productExists(uint(identifier))
        if err != nil {
            respondToClient(w, 400, nil, err.Error())
        }
        
        if !ok {
            respondToClient(w, 404, nil, "Specified product record not found")
        }
        
        respondToClient(w, 200, product, "")
    }else{
        respondToClient(w, 403, nil, "Invalid product identifier")
    }
}

func ReadAllProducts(w http.ResponseWriter, r *http.Request) {
    var products []db.Product
    response := database.Find(&products)
    numberOfRowsFound := response.RowsAffected
    msg := fmt.Sprintf("Found %d records", numberOfRowsFound)
    respondToClient(w, 200, products, msg)
}



