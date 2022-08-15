package db

import (
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
  "fmt"
)


func Connect() (*gorm.DB, error) {
	dbUser := "root"
	dbPswd := ""
	dbHost := "127.0.0.1"
	dbPort := "3306"
	dbName := "victoria-bms-db"
	
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPswd, dbHost, dbPort, dbName)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
      if err != nil {
        return nil, err
      }
    //======= Migrate the schema ======
    db.AutoMigrate(&Department{})
    db.AutoMigrate(&Service{})
    db.AutoMigrate(&Product{})
    db.AutoMigrate(&PackageProduct{})
    db.AutoMigrate(&PackageService{})
    db.AutoMigrate(&Package{})
    db.AutoMigrate(&Visit{})
    db.AutoMigrate(&OrderOrBooking{})
    db.AutoMigrate(&Payment{})
    db.AutoMigrate(&Invoice{})
    db.AutoMigrate(&Bill{})
    db.AutoMigrate(&StockTransaction{})
    db.AutoMigrate(&Expense{})
    db.AutoMigrate(&User{})
    db.AutoMigrate(&Message{})
    db.AutoMigrate(&UserAction{})
    db.AutoMigrate(&Notification{})
    
	return db, nil
}




