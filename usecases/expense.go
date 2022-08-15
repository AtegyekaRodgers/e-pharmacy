package usecases

import (
    "fmt"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/mahani-software-engineering/bms-server/db"
)

func CreateExpense(w http.ResponseWriter, r *http.Request) {
    expenses := struct{Expenses []db.Expense }{}
    _ = json.NewDecoder(r.Body).Decode(&expenses)
    database.Create(&expenses.Expenses)
    msg := fmt.Sprintf("New expense recorded")
    respondToClient(w, 201, expenses, msg)
}

func UpdateExpense(w http.ResponseWriter, r *http.Request) {
    //TODO: implement
}

func expenseExists (identifier uint) (bool, db.Expense, error) {
    //the identifier must be ID
    var expense db.Expense
    response := database.Where("id = ?", identifier).First(&expense)                   
    numberOfRowsFound := response.RowsAffected
    exists := numberOfRowsFound > 0
    return exists, expense, response.Error
}

func ReadExpense(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    identf := params["id"]
    if identifier, err := strconv.Atoi(identf); err == nil {
        ok, expense, err := expenseExists(uint(identifier))
        if err != nil {
            respondToClient(w, 400, nil, err.Error())
        }
        
        if !ok {
            respondToClient(w, 404, nil, "Specified expense record not found")
        }
        
        respondToClient(w, 200, expense, "")
    }else{
        respondToClient(w, 403, nil, "Invalid expense identifier")
    }
}

func ReadAllExpenses(w http.ResponseWriter, r *http.Request) {
    var expenses []db.Expense
    response := database.Find(&expenses)
    numberOfRowsFound := response.RowsAffected
    exists := numberOfRowsFound > 0
    fmt.Println(numberOfRowsFound, "expenses exist =", exists)
    msg := fmt.Sprintf("Found %d expenses", numberOfRowsFound)
    respondToClient(w, 200, expenses, msg)
}

func ReadExpensesTotalAmount(w http.ResponseWriter, r *http.Request) {
    var expenses []db.Expense
    _ = database.Find(&expenses)
    var totalAmount uint
    for _, exp := range expenses {
        totalAmount = totalAmount + exp.Amount
    }
    respondToClient(w, 200, totalAmount, "")
}






