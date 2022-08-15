package usecases


import (
    "fmt"
    "strconv"
    "net/http"
    "gorm.io/gorm"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/mahani-software-engineering/bms-server/db"
)

func balanceStock(productId uint, op string, qnty uint, byStaffID uint){
    var product db.Product
    switch op {
        case "add":
            database.Model(&product).Where("id = ?", productId).UpdateColumn("quantity", gorm.Expr("quantity + ?", qnty))
        case "remove":
            database.Model(&product).Where("id = ?", productId).UpdateColumn("quantity", gorm.Expr("quantity - ?", qnty))
            //create a transaction 
            var transcn db.StockTransaction
            if product.ID > 0 {
                //this confirms the product actually exists
                transcn.ProductCategory = fmt.Sprintf("%s (%s)",product.Category, product.Name)
                transcn.OldQuantity = product.Quantity
                transcn.NewQuantity = transcn.OldQuantity + qnty
                transcn.Amount = qnty * product.Price
                transcn.CreatedBy = byStaffID
            
                database.Create(&transcn)
                
                specificActionDetails := fmt.Sprintf("Transation ID = %v, by: %d, %s %s", transcn.ID, transcn.CreatedBy, transcn.Transaction, transcn.ProductCategory)
                newActionRecord(transcn.CreatedBy, "ACN0003", "Stock transaction recorded", "Stock transaction", specificActionDetails)     
            }
        default: 
          fmt.Printf("Unknwn operation %s product\n", op)
    }
}

func CreateStockTransaction(w http.ResponseWriter, r *http.Request) {
    var transcn db.StockTransaction
    _ = json.NewDecoder(r.Body).Decode(&transcn)
    //retrieve product whose id is transcn.ProductID
    ok, product, err := productExists(transcn.ProductID)
    if err != nil {
        respondToClient(w, 404, nil, "Error! transaction not recorded.")
    }
    
    if ok {
        transcn.ProductCategory = product.Category
        transcn.OldQuantity = product.Quantity
        transcn.NewQuantity = transcn.OldQuantity + transcn.Quantity
    
        database.Create(&transcn)
        balanceStock(transcn.ProductID, transcn.Transaction, transcn.Quantity, transcn.CreatedBy)
        
        specificActionDetails := fmt.Sprintf("Transation ID = %v, by: %d, %s %s", transcn.ID, transcn.CreatedBy, transcn.Transaction, transcn.ProductCategory)
        newActionRecord(transcn.CreatedBy, "ACN0003", "Stock transaction recorded", "Stock transaction", specificActionDetails)     
        respondToClient(w, 201, transcn, "Stock transaction recorded successfully")
    }else{
        respondToClient(w, 404, nil, "Error! The specified product is not registered in the system.")
    }
}

func UpdateStockTransaction(w http.ResponseWriter, r *http.Request) {
    //TODO: implement
}

func stockTransactionExists (identifier uint) (bool, db.StockTransaction, error) {
    //the identifier can be ID, phone, email, username
    var transcn db.StockTransaction
    response := database.Where("id = ?", identifier).First(&transcn)                   
    numberOfRowsFound := response.RowsAffected
    exists := numberOfRowsFound > 0
    return exists, transcn, response.Error
}

func ReadStockTransaction(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type","application/json")
    params := mux.Vars(r)
    identf := params["id"]
    if identifier, err := strconv.Atoi(identf); err == nil {
        ok, transcn, err := stockTransactionExists(uint(identifier))
        if err != nil {
            respondToClient(w, 400, nil, err.Error())
        }
        
        if !ok {
            respondToClient(w, 404, nil, "Specified transaction record not found")
        }
        
        respondToClient(w, 200, transcn, "")
    }else{
        respondToClient(w, 403, nil, "Invalid transaction identifier")
    }
}

func ReadAllStockTransactions(w http.ResponseWriter, r *http.Request) {
    var transcns []db.StockTransaction
    response := database.Find(&transcns)
    numberOfRowsFound := response.RowsAffected
    msg := fmt.Sprintf("Found %d records", numberOfRowsFound)
    respondToClient(w, 200, transcns, msg)
}

func ReadStockTxnsForBarner(w http.ResponseWriter, r *http.Request){
    type StockLedgerRecord struct {
      Yyyy             int     `json:"yyyy"`
      Mm               int     `json:"mm"`
      Dd               int     `json:"dd"`
      Commodity        string  `json:"commodity"`     //txn.ProductCategory
      Units            string  `json:"units"`
      Openingstock     uint    `json:"openingstock"`  //txn.OldQuantity
      Newstock         uint    `json:"newstock"`      //txn.Quantity(add), 
      Sales            uint    `json:"sales"`         //txn.Quantity(remove),	
      Rate             uint    `json:"rate"`        	
      ClosingStock     uint    `json:"closingstock"`  //txn.NewQuantity,	
      Revenue          uint    `json:"revenue"`       //txn.Amount
    }
    
    var ledgerRecords []StockLedgerRecord
    subTable1 := database.Model(&db.StockTransaction{}).Where("transaction = ?", "remove").Select("product_id, created_at, year(created_at) as yyyy, month(created_at) as mm, day(created_at) as dd, product_category as commodity, old_quantity as openingstock, sum(quantity) as sales, new_quantity as closingstock, sum(amount) as revenue").Group("product_id")      
    subTable2 := database.Model(&db.StockTransaction{}).Where("transaction = ?", "add").Select("product_id as pid, created_at as at, sum(quantity) as newstock").Group("product_id")
    joinedTable := database.Table("(?) as s",subTable1).Joins("JOIN (?) as p on p.pid=s.product_id",subTable2)
    database.Table("(?) as t",joinedTable).Joins("JOIN products on products.id=t.product_id").Select("t.*, products.quantity_units as units, products.price as rate").Find(&ledgerRecords)
    
    respondToClient(w, 200, ledgerRecords, "")
}   

//db.Where("name1 = @name OR name2 = @name", sql.Named("name", "jinzhu")).Find(&user)

