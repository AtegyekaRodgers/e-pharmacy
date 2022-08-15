package db

import (
  "gorm.io/gorm"
  
)

var database *gorm.DB

//=========================================================================

func (sv *Service) Save() {
  database.Create(sv)
}

func (sv *Service) Update(newService Service) {
  //update only non-zero values in newService
  database.Model(sv).Updates(newService)
}

//=========================================================================

func (pr *Product) Save() {
  database.Create(pr)
}

func (pr *Product) Update(newProduct Product) {
  //update only non-zero values in newProduct
  database.Model(pr).Updates(newProduct)
}

//=========================================================================

func (pp *PackageProduct) Save() {
  database.Create(pp)
}

func (pp *PackageProduct) Update(newPackageProduct PackageProduct) {
  //update only non-zero values in newPackageProduct
  database.Model(pp).Updates(newPackageProduct)
}

//=========================================================================

func (ps *PackageService) Save() {
  database.Create(ps)
}

func (ps *PackageService) Update(newPackageService PackageService) {
  //update only non-zero values in newPackageService
  database.Model(ps).Updates(newPackageService)
}

//=========================================================================

func (pa *Package) Save() {
  database.Create(pa)
}

//=========================================================================

func (vs *Visit) Save() {
  database.Create(vs)
}

//=========================================================================

func (od *OrderOrBooking) Save() {
  database.Create(od)
}

//=========================================================================

func (pm *Payment) Save() {
  database.Create(pm)
}

//=========================================================================

func (iv *Invoice) Save() {
  database.Create(iv)
}

//=========================================================================

func (bl *Bill) Save() {
  database.Create(bl)
}

//=========================================================================

func (st *StockTransaction) Save() {
  database.Create(st)
}

//=========================================================================

func (ep *Expense) Save() {
  database.Create(ep)
}

//=========================================================================

func (us *User) Save() {
  database.Create(us)
}

//=========================================================================

func (ua *UserAction) Save() {
  database.Create(ua)
}

//=========================================================================

func (ms *Message) Save() {
  database.Create(ms)
}

//=========================================================================

func (nt *Notification) Save() {
  database.Create(nt)
}

//=========================================================================



