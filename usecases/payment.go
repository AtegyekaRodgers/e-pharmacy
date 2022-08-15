package usecases


import (
    "fmt"
    "strconv"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/mahani-software-engineering/bms-server/db"
)

func CreatePayment(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    dcase := params["id"]
    
    type ReceivedOrder struct {
        Item     string  `json:"item" gorm:"size:8"`
        ItemID   uint    `json:"itemid"`
        Quantity uint    `json:"quantity"`
        Status   string  `json:"status"`
    }
    type ReceivedInvoice struct {
        Orders []db.OrderOrBooking `json:"orders"`
    }
    type ReceivedBill struct {
        Invoices string              `json:"invoices"`
        Orders   []db.OrderOrBooking `json:"orders"`
    }
    type ReceivedPayment struct {
        Order   ReceivedOrder   `json:"order"`
        Invoice ReceivedInvoice `json:"invoice"`
        Bill    ReceivedBill    `json:"bill"`
        Amount     uint         `json:"amount"`
        Item       string       `json:"item"`       //order, booking, invoice, bill
        ItemID     uint         `json:"itemid"`
        Instalment bool         `json:"instalment"` //true/false
        Customer   db.User      `json:"customer"`
        CustomerID uint         `json:"customerid"`
        CreatedBy  uint         `json:"createdby"`
    }
    switch (dcase) {
        case "5":
            //customer does not exist, no new customer info given, an order in the request, item=order
            var recvdPmt ReceivedPayment
            var pmt db.Payment
	        _ = json.NewDecoder(r.Body).Decode(&recvdPmt)
	        
	        //create an order first
	        orderNew := db.OrderOrBooking{
	            Item:      recvdPmt.Order.Item,
	            ItemID:    recvdPmt.Order.ItemID,
	            Quantity:  recvdPmt.Order.Quantity,
	            Status:    recvdPmt.Order.Status,
	            Paid:      true,
	            CreatedBy: recvdPmt.CreatedBy,
	            //Customer: {}, //ignored on creation
	        }
	        
	        database.Omit("Customer","CustomerID","VisitID","InvoiceID","BillID").Create(&orderNew)
	        //ID of the newly created order = orderNew.ID
	        pmt.ItemID = orderNew.ID
	        pmt.Item = recvdPmt.Item
	        pmt.Amount = recvdPmt.Amount
	        pmt.CreatedBy = recvdPmt.CreatedBy
	        database.Omit("Customer","CustomerID").Create(&pmt)
	        if recvdPmt.Order.Item == "product" && orderNew.Status == "served" {
	            balanceStock(recvdPmt.Order.ItemID, "remove", recvdPmt.Order.Quantity, recvdPmt.CreatedBy)
	        }
	        //finished.
	        
	        specificActionDetails := fmt.Sprintf("Unregistered customer oredered and paid: %s", orderNew.Item)
	        newActionRecord(pmt.CreatedBy, "ACN0002/5", "Payment received, customer not registered, new customer info not given, customer ordered for an item, item=order ", "Payment",specificActionDetails)     
	        respondToClient(w, 201, pmt, "Payment recorded successfully")
	        //done.
        case "6":
            //customer does not exist, no new customer info given, an order in the request, item=invoice
            var recvdPmt ReceivedPayment
            var pmt db.Payment
	        _ = json.NewDecoder(r.Body).Decode(&recvdPmt)
	        
	        //create an order first
	        invoiceNew := db.Invoice{
	            Amount:    recvdPmt.Amount,
	            Status:    "paid",
	            CreatedBy: recvdPmt.CreatedBy,
	            Orders:    recvdPmt.Invoice.Orders,
	            //Customer: {}, //ignored on creation
	        }
	        //invoiceOrders := &invoiceNew.Orders
            for k, order := range invoiceNew.Orders {
                order.Paid = true
                order.CreatedBy = recvdPmt.CreatedBy
                invoiceNew.Orders[k] = order
                if order.Item == "product" && order.Status == "served" {
	                balanceStock(order.ItemID, "remove", order.Quantity, recvdPmt.CreatedBy)
	            }
            }
	        database.Omit("Customer","CustomerID","Orders.Customer", "Orders.CustomerID","Orders.VisitID","Orders.BillID").Create(&invoiceNew)
	        //ID of the newly created invoice = invoiceNew.ID
	        pmt.ItemID = invoiceNew.ID
	        pmt.Item   = recvdPmt.Item
	        pmt.Amount = recvdPmt.Amount
	        pmt.CreatedBy = recvdPmt.CreatedBy
	        database.Omit("Customer","CustomerID").Create(&pmt)
	        //finished.
	        
	        specificActionDetails := fmt.Sprintf("Unregistered customer oredered and paid: %s", invoiceNew.ID)
	        newActionRecord(pmt.CreatedBy, "ACN0002/6", "Payment received, customer not registered, new customer info not given, customer ordered for an item, item=invoice ", "Payment", specificActionDetails)     
	        respondToClient(w, 201, pmt, "Payment recorded successfully")
	        //done.
        case "7":
            //customer does not exist, no new customer info given, an order in the request, item=bill
            var recvdPmt ReceivedPayment
            var pmt db.Payment
	        _ = json.NewDecoder(r.Body).Decode(&recvdPmt)
	        _ = json.NewDecoder(r.Body).Decode(&pmt)
	        
	        //create an order first
	        billNew := db.Bill{
	            Amount:    recvdPmt.Amount,
	            Status:    "paid",
	            CreatedBy: recvdPmt.CreatedBy,
	            Invoices:  recvdPmt.Bill.Invoices,
	            Orders:    recvdPmt.Bill.Orders,
	            //Customer: {}, //ignored on creation
	        }
	        //billOrders := &billNew.Orders
            for k, order := range billNew.Orders {
                order.Status = "billed"
                order.Paid = true
                order.CreatedBy = recvdPmt.CreatedBy
                billNew.Orders[k] = order
                if order.Item == "product" {
	                balanceStock(order.ItemID, "remove", order.Quantity, recvdPmt.CreatedBy)
	            }
            }
	        database.Omit("Customer","CustomerID","Orders.Customer","Orders.CustomerID","Orders.VisitID","Orders.InvoiceID").Create(&billNew)
	        //ID of the newly created bill = billNew.ID
	        pmt.ItemID = billNew.ID
	        pmt.Item   = recvdPmt.Item
	        pmt.Amount = recvdPmt.Amount
	        pmt.CreatedBy = recvdPmt.CreatedBy
	        database.Omit("Customer","CustomerID").Create(&pmt)
	        //finished.
	        
	        specificActionDetails := fmt.Sprintf("Unregistered customer oredered and paid: %s", billNew.ID)
	        newActionRecord(pmt.CreatedBy, "ACN0002/7", "Payment received, customer not registered, new customer info not given, customer ordered for an item, item=bill ", "Payment", specificActionDetails)     
	        respondToClient(w, 201, pmt, "Payment recorded successfully")
	        //done.
        case "8":
            //customer does not exist, new customer info given, no an order in the request, item=none
            var pmt db.Payment
	        _ = json.NewDecoder(r.Body).Decode(&pmt)
	        if pmt.Amount > 0 {
	            database.Create(&pmt)
	            specificActionDetails := fmt.Sprintf("New customer registered and paid advance for future orders: payment ID = %s", pmt.ID)
	            newActionRecord(pmt.CreatedBy, "ACN0002/8", "Payment received, customer was not registered, new customer info given, no new order by customer", "Payment", specificActionDetails)     
	            respondToClient(w, 201, pmt, "Payment recorded successfully")
	        }else{
	            var custmr db.User
	            database.Create(&pmt.Customer)
	            specificActionDetails := fmt.Sprintf("New customer registered: customer ID = %s", custmr.ID)
	            newActionRecord(custmr.ID, "ACN0002/8", "New customer was registered through create payment endpoint, no new order by customer", "Customer", specificActionDetails)     
	            respondToClient(w, 201, custmr, "Customer registered successfully")
	        }
	        //done.
        case "13":
            //customer does not exist, new customer info given, an order in the request, item=order
            var recvdPmt ReceivedPayment
            var pmt db.Payment
	        _ = json.NewDecoder(r.Body).Decode(&recvdPmt)
	        
	        //create an order first
	        orderNew := db.OrderOrBooking{
	            Item:      recvdPmt.Order.Item,
	            ItemID:    recvdPmt.Order.ItemID,
	            Quantity:  recvdPmt.Order.Quantity,
	            Status:    recvdPmt.Order.Status,
	            Paid:      true,
	            CreatedBy: recvdPmt.CreatedBy,
	            //Customer: {}, //ignored on creation
	        }
	        database.Omit("Customer","CustomerID","VisitID","InvoiceID","BillID").Create(&orderNew)
	        //ID of the newly created order = orderNew.ID
	        pmt.ItemID = orderNew.ID
	        pmt.Item   = recvdPmt.Item
	        pmt.Amount = recvdPmt.Amount
	        pmt.Customer = recvdPmt.Customer
	        pmt.CreatedBy = recvdPmt.CreatedBy
	        database.Create(&pmt)
	        //find the ID of the newly auto created customer, and use it to update orderNew's "customerid" field
	        ok, custmr, err := userExists(pmt.Customer.Phone)
	        if err != nil {
	            fmt.Println("Error occured while trying to get the ID of the customer that was auto-created by creating payment.")
	        }
	        
	        if ok {
	            orderNew.CustomerID = custmr.ID
	        }
	        //now ready to update the orderNew in the database
	        database.Save(&orderNew)
	        if orderNew.Item == "product" && orderNew.Status == "served" {
                balanceStock(orderNew.ItemID, "remove", orderNew.Quantity, recvdPmt.CreatedBy)
            }
	        //finished.
	        
	        specificActionDetails := fmt.Sprintf("Customer: %s phone: %s",custmr.ID,custmr.Phone)
	        newActionRecord(pmt.CreatedBy, "ACN0002/13", "Payment received, customer not registered, new customer info given, customer ordered for an item, item=order ", "Payment", specificActionDetails)     
	        respondToClient(w, 201, pmt, "Payment recorded successfully")
	        //done.
        case "14":
            //customer does not exist, new customer info given, an order in the request, item=invoice
            var recvdPmt ReceivedPayment
            var pmt db.Payment
	        _ = json.NewDecoder(r.Body).Decode(&recvdPmt)
	        
	        //create an order first
	        invoiceNew := db.Invoice{
	            Amount:    recvdPmt.Amount,
	            Status:    "paid",
	            CreatedBy: recvdPmt.CreatedBy,
	            Orders:    recvdPmt.Invoice.Orders,
	            //Customer: {}, //ignored on creation
	        }
	        //invoiceOrders := &invoiceNew.Orders
            for k, order := range invoiceNew.Orders {
                order.Paid = true
                order.CreatedBy = recvdPmt.CreatedBy
                invoiceNew.Orders[k] = order
                if order.Item == "product" && order.Status == "served" {
                    balanceStock(order.ItemID, "remove", order.Quantity, recvdPmt.CreatedBy)
                }
            }
	        database.Omit("Customer","CustomerID","Orders.Customer","Orders.CustomerID","Orders.VisitID","Orders.BillID").Create(&invoiceNew)
	        //ID of the newly created invoice = invoiceNew.ID
	        pmt.ItemID = invoiceNew.ID
	        pmt.Item   = recvdPmt.Item
	        pmt.Amount = recvdPmt.Amount
	        pmt.Customer = recvdPmt.Customer
	        pmt.CreatedBy = recvdPmt.CreatedBy
	        database.Create(&pmt)
	        //find the ID of the newly auto created customer, and use it to update invoiceNew's "customerid" field
	        invoiceNew.CustomerID = pmt.Customer.ID
	            
	        for k, order := range invoiceNew.Orders {
                order.CustomerID = pmt.Customer.ID
                order.InvoiceID = invoiceNew.ID
                invoiceNew.Orders[k] = order
            }
            database.Table("order_or_bookings").Where("invoice_id = ?", invoiceNew.ID).Update("customer_id", pmt.Customer.ID)
            
	        //now ready to update the invoiceNew in the database
	        database.Save(&invoiceNew)
            //finished.
	        
	        specificActionDetails := fmt.Sprintf("Customer: %s phone: %s",pmt.Customer.ID, pmt.Customer.Phone)
	        newActionRecord(pmt.CreatedBy, "ACN0002/14", "Payment received, customer not registered, new customer info given, customer ordered for an item, item=invoice ", "Payment", specificActionDetails)     
	        respondToClient(w, 201, pmt, "Payment recorded successfully")
	        //done.
        case "15":
            //customer does not exist, new customer info given, an order in the request, item=bill
            var recvdPmt ReceivedPayment
            var pmt db.Payment
	        _ = json.NewDecoder(r.Body).Decode(&recvdPmt)
	        
	        //create an order first
	        billNew := db.Bill{
	            Amount:    recvdPmt.Amount,
	            Status:    "paid",
	            CreatedBy: recvdPmt.CreatedBy,
	            Invoices:  recvdPmt.Bill.Invoices,
	            Orders:    recvdPmt.Bill.Orders,
	            //Customer: {}, //ignored on creation
	        }
	        //billOrders := &billNew.Orders
            for k, order := range billNew.Orders {
                order.Status = "billed"
                order.Paid = true
                order.CreatedBy = recvdPmt.CreatedBy
                billNew.Orders[k] = order
                if order.Item == "product" {
                    balanceStock(order.ItemID, "remove", order.Quantity, recvdPmt.CreatedBy)
                }
            }
            
	        database.Omit("Customer","CustomerID","Orders.Customer","Orders.CustomerID","Orders.VisitID","Orders.InvoiceID").Create(&billNew)
	        //ID of the newly created bill = billNew.ID
	        pmt.ItemID = billNew.ID
	        pmt.Item   = recvdPmt.Item
	        pmt.Amount = recvdPmt.Amount
	        pmt.Customer = recvdPmt.Customer
	        pmt.CreatedBy = recvdPmt.CreatedBy
	        database.Create(&pmt)
	        //find the ID of the newly auto created customer, and use it to update billNew's "customerid" field
	        billNew.CustomerID = pmt.Customer.ID
	        
	        for k, order := range billNew.Orders {
                order.BillID = billNew.ID
                order.CustomerID = pmt.Customer.ID
                billNew.Orders[k] = order
            }
            database.Table("order_or_bookings").Where("bill_id = ?", billNew.ID).Update("customer_id", pmt.Customer.ID)
            
	        //now ready to update the billNew in the database
	        database.Save(&billNew)
            //finished.
	        
	        specificActionDetails := fmt.Sprintf("Customer: %s phone: %s",pmt.Customer.ID, pmt.Customer.Phone)
	        newActionRecord(pmt.CreatedBy, "ACN0002/15", "Payment received, customer not registered, new customer info given, customer ordered for an item, item=bill ", "Payment", specificActionDetails)     
	        respondToClient(w, 201, pmt, "Payment recorded successfully")
	        //done.
        case "17":
            //customer exists, no new customer info given, no order in the request, item=order
            var pmt db.Payment
	        _ = json.NewDecoder(r.Body).Decode(&pmt)
	        
	        ok, custmr, err := userExists(fmt.Sprintf("%s",pmt.CustomerID))
	        if err != nil {
	            fmt.Println("Error occured while trying to confirm existace of specified custmer on receiving payment.")
	        }
	        
	        if !ok {
	            respondToClient(w, 428, pmt, "Specified customer not found. Please register first")
	        }
	        
	        ook, order, errr := orderExists(pmt.ItemID)
	        if errr != nil {
	            fmt.Printf("Can't find the specified %s item.\n",pmt.Item)
	        }
	        
	        if !ook {
	            msge := fmt.Sprintf("Can't find the specified %s item.", pmt.Item)
	            respondToClient(w, 428, pmt, msge)
	        }
	        
	        order.Paid = true
	        order.CustomerID = pmt.CustomerID
	        database.Save(&order)
	        if order.Item == "product" && order.Status == "served" {
                balanceStock(order.ItemID, "remove", order.Quantity, pmt.CreatedBy)
            }
	        
	        database.Omit("Customer").Create(&pmt)
	        //finished.
	        
	        specificActionDetails := fmt.Sprintf("Registered customer (id: %s) paid for earlier order: %s payment ID = %s", custmr.ID, order.ItemID, pmt.ID)
	        newActionRecord(pmt.CreatedBy, "ACN0002/17", "Payment received, from registered customer, no new order by customer", "Payment", specificActionDetails)     
	        respondToClient(w, 201, pmt, "Payment recorded successfully")
	        //done.
        case "21":
            //customer exists, no new customer info given, an order in the request, item=order
            var recvdPmt ReceivedPayment
            var pmt db.Payment
	        _ = json.NewDecoder(r.Body).Decode(&recvdPmt)
	        
	        //create an order first
	        orderNew := db.OrderOrBooking{
	            Item:      recvdPmt.Order.Item,
	            ItemID:    recvdPmt.Order.ItemID,
	            Quantity:  recvdPmt.Order.Quantity,
	            Status:    recvdPmt.Order.Status,
	            Paid:      true,
	            CreatedBy: recvdPmt.CreatedBy,
	            CustomerID: recvdPmt.CustomerID,
	            //Customer: {}, //ignored on creation
	        }
	        database.Omit("Customer","VisitID","InvoiceID","BillID").Create(&orderNew)
	        //ID of the newly created order = orderNew.ID
	        pmt.ItemID = orderNew.ID
	        pmt.Item   = recvdPmt.Item
	        pmt.Amount = recvdPmt.Amount
	        pmt.CustomerID = recvdPmt.CustomerID
	        pmt.CreatedBy = recvdPmt.CreatedBy
	        database.Omit("Customer").Create(&pmt)
	        if orderNew.Item == "product" && orderNew.Status == "served" {
                balanceStock(orderNew.ItemID, "remove", orderNew.Quantity, recvdPmt.CreatedBy)
            }
	        //finished.
	        
	        specificActionDetails := fmt.Sprintf("Customer: %d, Ordered: %d",pmt.CustomerID, orderNew.ID)
	        newActionRecord(pmt.CreatedBy, "ACN0002/21", "Payment received, customer is registered, new customer info not needed, customer ordered for an item, item=order", "Payment", specificActionDetails)     
	        respondToClient(w, 201, pmt, "Payment recorded successfully")
	        //done.
        case "22":
            //customer exists, no new customer info given, an order in the request, item=invoice
            var recvdPmt ReceivedPayment
            var pmt db.Payment
	        _ = json.NewDecoder(r.Body).Decode(&recvdPmt)
	        
	        //create an order first
	        invoiceNew := db.Invoice{
	            Amount:    recvdPmt.Amount,
	            Status:    "paid",
	            CreatedBy: recvdPmt.CreatedBy,
	            Orders:    recvdPmt.Invoice.Orders,
	            CustomerID: recvdPmt.CustomerID,
	            //Customer: {}, //ignored on creation
	        }
	        //invoiceOrders := &invoiceNew.Orders
            for k, order := range invoiceNew.Orders {
                order.Paid = true
                order.CustomerID = recvdPmt.CustomerID
                order.CreatedBy = recvdPmt.CreatedBy
                invoiceNew.Orders[k] = order
                if order.Item == "product" && order.Status == "served" {
                    balanceStock(order.ItemID, "remove", order.Quantity, recvdPmt.CreatedBy)
                }
            }
	        database.Omit("Customer","Orders.Customer","Orders.VisitID","Orders.BillID").Create(&invoiceNew)
	        //ID of the newly created invoice = invoiceNew.ID
	        pmt.ItemID = invoiceNew.ID
	        pmt.Item   = recvdPmt.Item
	        pmt.Amount = recvdPmt.Amount
	        pmt.CustomerID = recvdPmt.CustomerID
	        pmt.CreatedBy = recvdPmt.CreatedBy
	        database.Omit("Customer").Create(&pmt)
	        //finished.
	        
	        specificActionDetails := fmt.Sprintf("Customer: %s, Ordered: %s",pmt.CustomerID, invoiceNew.ID)
	        newActionRecord(pmt.CreatedBy, "ACN0002/22", "Payment received, customer is registered, new customer info not needed, customer ordered for an item, item=invoice", "Payment", specificActionDetails)     
	        respondToClient(w, 201, pmt, "Payment recorded successfully")
	        //done.
        case "23":
            //customer exists, no new customer info given, an order in the request, item=bill
            var recvdPmt ReceivedPayment
            var pmt db.Payment
	        _ = json.NewDecoder(r.Body).Decode(&recvdPmt)
	        
	        //create an order first
	        billNew := db.Bill{
	            Amount:    recvdPmt.Amount,
	            Status:    "paid",
	            CreatedBy: recvdPmt.CreatedBy,
	            Orders:    recvdPmt.Bill.Orders,
	            Invoices:  recvdPmt.Bill.Invoices,
	            CustomerID: recvdPmt.CustomerID,
	            //Customer: {}, //ignored on creation
	        }
	        //billOrders := &billNew.Orders
            for k, order := range billNew.Orders {
                order.Status = "billed"
                order.Paid = true
                order.CustomerID = recvdPmt.CustomerID
                order.CreatedBy = recvdPmt.CreatedBy
                billNew.Orders[k] = order
                if order.Item == "product" {
                    balanceStock(order.ItemID, "remove", order.Quantity, recvdPmt.CreatedBy)
                }
            }
	        database.Omit("Customer","Orders.Customer","Orders.VisitID","Orders.InvoiceID").Create(&billNew)
	        //ID of the newly created invoice = billNew.ID
	        pmt.ItemID = billNew.ID
	        pmt.Item   = recvdPmt.Item
	        pmt.Amount = recvdPmt.Amount
	        pmt.CustomerID = recvdPmt.CustomerID
	        pmt.CreatedBy = recvdPmt.CreatedBy
	        database.Omit("Customer").Create(&pmt)
	        //finished.
	        
	        specificActionDetails := fmt.Sprintf("Customer: %d Ordered: %d",pmt.CustomerID, billNew.ID)
	        newActionRecord(pmt.CreatedBy, "ACN0002/23", "Payment received, customer is registered, new customer info not needed, customer ordered for an item, item=bill", "Payment", specificActionDetails)     
	        respondToClient(w, 201, pmt, "Payment recorded successfully")
	        //done.
        default:
            respondToClient(w, 400, nil, "Bad Request (the service request is not well specified)")
    }
}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
    //TODO: implement
}

func paymentExists (identifier uint) (bool, db.Payment, error) {
    //the identifier must be ID
    var payment db.Payment
    response := database.Where("id = ?", identifier).First(&payment)                   
    numberOfRowsFound := response.RowsAffected
    exists := numberOfRowsFound > 0
    fmt.Println(numberOfRowsFound, "payments exist =", exists)
    return exists, payment, response.Error
}

func ReadPayment(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    identf := params["id"]
    if identifier, err := strconv.Atoi(identf); err == nil {
        ok, payment, err := paymentExists(uint(identifier))
        if err != nil {
            respondToClient(w, 400, nil, err.Error())
        }
        
        if !ok {
            respondToClient(w, 404, nil, "Specified payment record not found")
        }
        
        respondToClient(w, 200, payment, "")
    }else{
        respondToClient(w, 403, nil, "Invalid payment identifier")
    }
}

func ReadAllPayments(w http.ResponseWriter, r *http.Request) {
    var payments []db.Payment
    response := database.Find(&payments)
    numberOfRowsFound := response.RowsAffected
    msg := fmt.Sprintf("Found %d payments", numberOfRowsFound)
    respondToClient(w, 200, payments, msg)
}

func ReadCountPayments(w http.ResponseWriter, r *http.Request) {
    var payments []db.Payment
    response := database.Find(&payments)
    numberOfRowsFound := response.RowsAffected
    msg := fmt.Sprintf("Counted %d payments", numberOfRowsFound)
    respondToClient(w, 200, numberOfRowsFound, msg)
}

func ReadPaymentsTotalAmount(w http.ResponseWriter, r *http.Request) {
    var payments []db.Payment
    _ = database.Find(&payments)
    var totalAmount uint
    for _, pmt := range payments {
        totalAmount = totalAmount + pmt.Amount
    }
    respondToClient(w, 200, totalAmount, "")
}





