package usecases


import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

type InputOption struct {
    ID     uint   `json:"id"`
    Value  string `json:"value"`
}

func ReadInputOptions(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    entity := params["entity"]
    var options  []InputOption
    var msg string
    switch entity {
       case "product":
          _ = database.Table("products").Select(`id, CONCAT(name, "-",quantity_units, " - (product)") as value`).Find(&options)
       case "service":
          _ = database.Table("services").Select(`id, CONCAT(name," - (service)") as value`).Find(&options)
       case "package":
          _ = database.Table("packages").Select(`id, CONCAT(name, "-",price, " - (package)") as value`).Find(&options)
       case "order":
          _ = database.Table("order_or_bookings").Select(`id, CONCAT("order/booking-",id) as value`).Find(&options)
       case "invoice":
          _ = database.Table("invoices").Select(`id, CONCAT("invoice-",id) as value`).Find(&options)
       case "bill":
          _ = database.Table("bills").Select(`id, CONCAT("bill-",id) as value`).Find(&options)
       case "pdtsupplier":
          _ = database.Table("users").Where("user_type = ? OR user_type = ?", "supplier", "staff").Select(`id, CONCAT(firstname, " ",lastname, " - (", user_type,")") as value`).Find(&options) 
       case "svcprovider":
          _ = database.Table("users").Where("user_type = ? OR user_type = ?", "service-provider", "staff").Select(`id, CONCAT(firstname, " ",lastname, " - (", user_type,")") as value`).Find(&options) 
       case "user":
          _ = database.Table("users").Select(`id, CONCAT(firstname, " ",lastname, " - (", user_type,")") as value`).Find(&options) 
       case "orderitems":
          var productOptions  []InputOption
          var serviceOptions  []InputOption
          var packageOptions  []InputOption
          _ = database.Table("products").Select(`id, CONCAT(name, "-",quantity_units, " - (product)") as value`).Find(&productOptions)
          _ = database.Table("services").Select(`id, CONCAT(name," - (service)") as value`).Find(&serviceOptions)
          _ = database.Table("packages").Select(`id, CONCAT(name, "-",price, " - (package)") as value`).Find(&packageOptions)
          options = append(options, productOptions...)
          options = append(options, serviceOptions...)
          options = append(options, packageOptions...)
       default:
          msg := fmt.Sprintf("Invalid options path '%s'",entity)
          respondToClient(w, 403, nil, msg)
    }
    
    respondToClient(w, 200, options, msg)
}









