package usecases


import (
    "fmt"
    "time"
    "net/http"
    //"encoding/json"
    //"github.com/mahani-software-engineering/bms-server/db"
)

type RevenueStat struct {
    Date       time.Time  `json:"date"`
    Department string     `json:"department"`
    Number     string     `json:"number"`
    Revenue    uint       `json:"revenue"`
}

func StatsCustomersThisWeek(w http.ResponseWriter, r *http.Request) {
    var orders []RevenueStat
    var serviceOrders []RevenueStat
    //find the date 7 days ago
    dateOf7daysAgo := time.Now().AddDate(0, 0, -7)
    fmt.Println("dateOf7daysAgo = ", dateOf7daysAgo)
    _ = database.Table("order_or_bookings as o").Joins("JOIN products as p on o.item_id=p.id").Joins("JOIN departments as d on d.id=p.department").Where("o.item = ? AND o.created_at >= ?", "product", dateOf7daysAgo).Group("p.department, day(o.created_at)").Select("o.created_at as date, d.name as department, count(o.created_at) as number, sum(p.price) as revenue").Find(&orders)            
    _ = database.Table("order_or_bookings as o").Joins("JOIN services as s on s.id=o.item_id").Joins("JOIN departments as d on d.id=s.department").Where("o.item = ? AND o.created_at >= ?", "service", dateOf7daysAgo).Group("s.department, day(o.created_at)").Select("o.created_at as date, d.name as department, count(o.created_at) as number, sum(s.price) as revenue").Find(&serviceOrders)     
    
    //concatenate results 
    orders = append(orders, serviceOrders...)
    
    respondToClient(w, 200, orders, "")
}

func StatsRevenueThisWeek(w http.ResponseWriter, r *http.Request) {
    var paidOrders []RevenueStat
    var paidServiceOrders []RevenueStat
    var paidPackageOrders []RevenueStat
    //find the date 7 days ago
    dateOf7daysAgo := time.Now().AddDate(0, 0, -7)
    fmt.Println("dateOf7daysAgo = ", dateOf7daysAgo)
    response1 := database.Table("order_or_bookings as o").Joins("JOIN products as p on o.item_id=p.id").Joins("JOIN departments as d on d.id=p.department").Where("o.item = ? AND o.created_at >= ? AND o.status = ? AND o.paid = true", "product", dateOf7daysAgo,"served").Group("p.department, day(o.created_at)").Select("o.created_at as date, d.name as department, count(o.created_at) as number, sum(p.price) as revenue").Find(&paidOrders)            
    response2 := database.Table("order_or_bookings as o").Joins("JOIN services as s on s.id=o.item_id").Joins("JOIN departments as d on d.id=s.department").Where("o.item = ? AND o.created_at >= ? AND o.status = ? AND o.paid = true", "service", dateOf7daysAgo,"served").Group("s.department, day(o.created_at)").Select("o.created_at as date, d.name as department, count(o.created_at) as number, sum(s.price) as revenue").Find(&paidServiceOrders)
    response3 := database.Table("order_or_bookings as o").Joins("JOIN packages as pk on pk.id=o.item_id").Joins("JOIN departments as d on d.id=pk.department").Where("o.item = ? AND o.created_at >= ? AND o.status = ? AND o.paid = true", "package", dateOf7daysAgo,"served").Group("pk.department, day(o.created_at)").Select("o.created_at as date, d.name as department, count(o.created_at) as number, sum(pk.price) as revenue").Find(&paidPackageOrders)
    
    fmt.Println("response1.RowsAffected=",response1.RowsAffected,"response2.RowsAffected=",response2.RowsAffected, "response3.RowsAffected=",response3.RowsAffected)
    //concatenate results 
    paidOrders = append(paidOrders, paidServiceOrders...) 
    paidOrders = append(paidOrders, paidPackageOrders...)
    
    respondToClient(w, 200, paidOrders, "")
}

func StatsRevenueSrcThisWeek(w http.ResponseWriter, r *http.Request) {
    var paidOrders []RevenueStat
    var paidServiceOrders []RevenueStat
    var paidPackageOrders []RevenueStat
    //find the date 7 days ago
    dateOf7daysAgo := time.Now().AddDate(0, 0, -7)
    fmt.Println("dateOf7daysAgo = ", dateOf7daysAgo)
    _ = database.Table("order_or_bookings as o").Joins("JOIN products as p on o.item_id=p.id").Joins("JOIN departments as d on d.id=p.department").Where("o.item = ? AND o.created_at >= ? AND o.status = ? AND o.paid = true", "product", dateOf7daysAgo,"served").Group("p.department, day(o.created_at)").Select("o.created_at as date, d.name as department, count(o.created_at) as number, sum(p.price) as revenue").Find(&paidOrders)            
    _ = database.Table("order_or_bookings as o").Joins("JOIN services as s on s.id=o.item_id").Joins("JOIN departments as d on d.id=s.department").Where("o.item = ? AND o.created_at >= ? AND o.status = ? AND o.paid = true", "service", dateOf7daysAgo,"served").Group("s.department, day(o.created_at)").Select("o.created_at as date, d.name as department, count(o.created_at) as number, sum(s.price) as revenue").Find(&paidServiceOrders)
    _ = database.Table("order_or_bookings as o").Joins("JOIN packages as pk on pk.id=o.item_id").Joins("JOIN departments as d on d.id=pk.department").Where("o.item = ? AND o.created_at >= ? AND o.status = ? AND o.paid = true", "package", dateOf7daysAgo,"served").Group("pk.department, day(o.created_at)").Select("o.created_at as date, d.name as department, count(o.created_at) as number, sum(pk.price) as revenue").Find(&paidPackageOrders)
    
    //concatenate results 
    paidOrders = append(paidOrders, paidServiceOrders...) 
    paidOrders = append(paidOrders, paidPackageOrders...)
    
    respondToClient(w, 200, paidOrders, "")
}

func StatsAnnualRevenue(w http.ResponseWriter, r *http.Request) {
    //var payments []db.Payment
    payments := []struct{Date time.Time `json:"date"`; Amount uint `json:"amount"`}{}
    dateOf1yearAgo := time.Now().AddDate(-1, 0, 0)
    _ = database.Table("payments").Group("month(created_at)").Where("created_at >= ?",dateOf1yearAgo).Select("created_at as date, sum(amount) as amount").Find(&payments)
    respondToClient(w, 200, payments, "")
}

func StatsRevenueSrcAnnual(w http.ResponseWriter, r *http.Request) {
    var paidOrders []RevenueStat
    var paidServiceOrders []RevenueStat
    var paidPackageOrders []RevenueStat
    //find the date 7 days ago
    dateOf1yearAgo := time.Now().AddDate(-1, 0, 0)
    fmt.Println("dateOf1yearAgo = ", dateOf1yearAgo)
    _ = database.Table("order_or_bookings as o").Joins("JOIN products as p on o.item_id=p.id").Joins("JOIN departments as d on d.id=p.department").Where("o.item = ? AND o.created_at >= ? AND o.status = ? AND o.paid = true", "product", dateOf1yearAgo,"served").Group("p.department, month(o.created_at)").Select("o.created_at as date, d.name as department, count(o.created_at) as number, sum(p.price) as revenue").Find(&paidOrders)            
    _ = database.Table("order_or_bookings as o").Joins("JOIN services as s on s.id=o.item_id").Joins("JOIN departments as d on d.id=s.department").Where("o.item = ? AND o.created_at >= ? AND o.status = ? AND o.paid = true", "service", dateOf1yearAgo,"served").Group("s.department, month(o.created_at)").Select("o.created_at as date, d.name as department, count(o.created_at) as number, sum(s.price) as revenue").Find(&paidServiceOrders)
    _ = database.Table("order_or_bookings as o").Joins("JOIN packages as pk on pk.id=o.item_id").Joins("JOIN departments as d on d.id=pk.department").Where("o.item = ? AND o.created_at >= ? AND o.status = ? AND o.paid = true", "package", dateOf1yearAgo,"served").Group("pk.department, month(o.created_at)").Select("o.created_at as date, d.name as department, count(o.created_at) as number, sum(pk.price) as revenue").Find(&paidPackageOrders)
    
    //concatenate results 
    paidOrders = append(paidOrders, paidServiceOrders...) 
    paidOrders = append(paidOrders, paidPackageOrders...)
    
    respondToClient(w, 200, paidOrders, "")
}












