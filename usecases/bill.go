package usecases


import (
    "fmt"
    "strconv"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/mahani-software-engineering/bms-server/db"
)

func CreateBill(w http.ResponseWriter, r *http.Request) {
    //params := mux.Vars(r)
    //customerId := params["id"]
    //if there are any invoices listed in the received request, 
    //first retrieve each of them and ensure that all orders in 
    //that invoice are updated to served if they are not yet paid (ie, paid=false).
    //NOTE: invoices in the request are to be a string of comma-separated invoice IDs
    //      split the string to get an array of the IDs and process for one at a time
    
    //if customer exists {
        //if (there are no existing orders where customerID is this customer's Id AND status is "served" AND paid=false){
            //if (there are no arders sent in the request){
                 //respond with "No bill candidates found. A bill must have orders or invoices served and not paid"
            //}esle if (user has made orders within this request sent ) {
                //get the total price of the items in the order list sent
                //create a new bill, auto create orders[], don't auto create user
                //ensure the new orders' "status" fields are set to "billed"
                //respond to client
            //}
        //}esle {
            //if (user has made orders within this request sent ) {
                //create a fresh instance of the bill entity, auto creare the new (additional) orders, get the new ID
                //ensure the new orders' "status" fields are set to "billed"
                //ensure the new orders' customerID fields are set to the customerId supplied
                //billAmount := 0  //to be accumulated
                //find the total price of all the newly received orders, add to billAmount
                //for each of the order records found existing, 
                    //link the order to this new bill (set this order's "BillID" field to this new bill's ID) 
                    //check what kind of item was ordered: product, service, package (from the order's "Item" field)
                    //retrieve details (name, price) of the item from the appropriate table
                    //multiply the known quantity of the current order in this iteration, calculate the "amount" of the invoice item
                    //add the amount to (accumulate) the "billAmount" variable
                    //modify the order's "status" field and set it to "billed"
                //end loop
                //set the value of the new Bill's "Amount" field to the accumulated billAmount.
                //respond to client
            //}else{
               //create a fresh instance of the bill entity, dont auto creare orders, get the new ID
               //billAmount := 0  //to be accumulated
               //for each of the order records found, 
                   //link the order to this new bill (set this order's "BillID" field to this new bill's ID) 
                   //check what kind of item was ordered: product, service, package (from the order's "Item" field)
                   //retrieve details (name, price) of the item from the appropriate table
                   //multiply the known quantity of the current order in this iteration, calculate the "amount" of the bill item
                   //add the amount to (accumulate) the "billAmount" variable
               //end loop
               //set the value of the new Bill's "Amount" field to the accumulated "billAmount".
               //respond to client
         //}
       //}
    //}else{
          //if (info for new customer provided){
                //if (user has made orders within this request sent ) {
                    //get the total price of the items in the order list sent
                    //create a new bill, auto create user and auto create orders inside the bill, 
                    //skip customer field of each new order created
                    //ensure the new orders' "status" fields are set to "billed"
                    //ensure the new orders' "CustomerID" fields are set to this new customer's ID
                    //respond to client
               //}else{
                    //create a new customer profile instead
                    //respond with "New customer registered, No bill candidates exist" 
               //}
         //}else{
              //if (user has made orders within this request sent ) {
                    //get the total price of the items in the order list sent
                    //create a new bill, auto create orders, dont auto create user
                    //ensure the new orders' "status" fields are set to "billed"
                    //ensure the new orders' "BillID" fields are set to this new bill's ID
                    //respond to client
               //}else{
                    //respond with "No bill candidates found, customer account not found"
               //}
         //}
  //}
}

func billExists (identifier uint) (bool, db.Bill, error) {
    //the identifier must be ID
    var bill db.Bill
    response := database.Where("id = ?", identifier).First(&bill)                   
    numberOfRowsFound := response.RowsAffected
    exists := numberOfRowsFound > 0
    return exists, bill, response.Error
}

func ReadBill(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)
    identf := params["id"]
    if identifier, err := strconv.Atoi(identf); err == nil {
        ok, bill, err := billExists(uint(identifier))
        if err != nil {
            respondToClient(w, 400, nil, err.Error())
        }
        
        if !ok {
            respondToClient(w, 404, nil, "Specified bill record not found")
        }
        
        respondToClient(w, 200, bill, "")
    }else{
        respondToClient(w, 403, nil, "Invalid bill identifier")
    }
}

func ReadAllBills(w http.ResponseWriter, r *http.Request) {
    var bills []db.Bill
    response := database.Find(&bills)
    numberOfRowsFound := response.RowsAffected
    exists := numberOfRowsFound > 0
    fmt.Println(numberOfRowsFound, "bills exist =", exists)
    msg := fmt.Sprintf("Found %d records", numberOfRowsFound)
    respondToClient(w, 200, bills, msg)
}



