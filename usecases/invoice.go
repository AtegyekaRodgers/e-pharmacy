package usecases


import (
    "fmt"
    "strconv"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/mahani-software-engineering/bms-server/db"
)

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
    //params := mux.Vars(r)
    //customerId := params["id"]
    //if customer exists {
        //if (there are no existing orders where customerID is this customer's Id AND status is "pending"){
            //if (there are no arders sent in the request){
                 //respond with "No pending orders to add to invoice"
            //}esle if (user has made orders within this request sent ) {
                //get the total price of the items in the order list sent
                //create a new invoice, auto create orders[], don't auto create user
                //ensure the new orders' "status" fields are set to "invoiced"
                //respond to client
            //}
        //}esle {
            //if (user has made orders within this request sent ) {
                //create a fresh instance of the invoice entity, auto creare the new (additional) orders, get the new ID
                //set the customerID field to the customerId supplied
                //invoiceAmount := 0  //to be accumulated
                //for each of the order records found, 
                    //link the order to this new invoice (set this order's "InvoiceID" field to this new invoice's ID) 
                    //check what kind of item was ordered: product, service, package (from the order's "Item" field)
                    //retrieve details (name, price) of the item from the appropriate table
                    //multiply the known quantity of the current order in this iteration, calculate the "amount" of the invoice item
                    //accumulate the "invoinceAmount" variable
                    //modify the order's "status" field and set it to "invoiced"
                //end loop
                //set the value of the new Invice's "Amount" field to the accumulated invoice amount.
                //respond to client
            //}else{
               //create a fresh instance of the invoice entity, dont auto creare orders, get the new ID
               //invoiceAmount := 0  //to be accumulated
               //for each of the order records found, 
                   //link the order to this new invoice (set this order's "InvoiceID" field to this new invoice's ID) 
                   //check what kind of item was ordered: product, service, package (from the order's "Item" field)
                   //retrieve details (name, price) of the item from the appropriate table
                   //multiply the known quantity of the current order in this iteration, calculate the "amount" of the invoice item
                   //accumulate the "invoince amount" variable
               //end loop
               //set the value of the new Invice's "Amount" field to the accumulated invoice amount.
               //respond to client
         //}
       //}
    //}else{
          //if (info for new customer provided){
                //if (user has made orders within this request sent ) {
                    //get the total price of the items in the order list sent
                    //create a new invoice, auto create user and auto create orders inside the invoice, 
                    //skip customer field of each new order created
                    //ensure the new orders' "status" fields are set to "invoiced"
                    //ensure the new orders' "CustomerID" fields are set to this new customer's ID
                    //respond to client
               //}else{
                    //create a new customer profile instead
                    //respond with "New customer registered, no pending orders to create invoice" 
               //}
         //}else{
              //if (user has made orders within this request sent ) {
                    //get the total price of the items in the order list sent
                    //create a new invoice, auto create orders, dont auto create user
                    //ensure the new orders' "status" fields are set to "invoiced"
                    //ensure the new orders' "InvoiceID" fields are set to this new invoice's ID
                    //respond to client
               //}else{
                    //respond with "No pending orders to add to invoice, customer account not found"
               //}
         //}
  //}
}

func invoiceExists (identifier uint) (bool, db.Invoice, error) {
    //the identifier must be ID
    var invoice db.Invoice
    response := database.Where("id = ?", identifier).First(&invoice)                   
    numberOfRowsFound := response.RowsAffected
    exists := numberOfRowsFound > 0
    return exists, invoice, response.Error
}

func ReadInvoice(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)
    identf := params["id"]
    if identifier, err := strconv.Atoi(identf); err == nil {
        ok, invoice, err := invoiceExists(uint(identifier))
        if err != nil {
            respondToClient(w, 400, nil, err.Error())
        }
        
        if !ok {
            respondToClient(w, 404, nil, "Specified invoice record not found")
        }
        
        respondToClient(w, 200, invoice, "")
    }else{
        respondToClient(w, 403, nil, "Invalid invoice identifier")
    }
}

func ReadAllInvoices(w http.ResponseWriter, r *http.Request) {
    var invoices []db.Invoice
    response := database.Find(&invoices)
    numberOfRowsFound := response.RowsAffected
    exists := numberOfRowsFound > 0
    fmt.Println(numberOfRowsFound, "invoices exist =", exists)
    msg := fmt.Sprintf("Found %d records", numberOfRowsFound)
    respondToClient(w, 200, invoices, msg)
}









