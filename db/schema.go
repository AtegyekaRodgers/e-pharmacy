package db

import (
  "gorm.io/gorm"
)

type BusinessGroup struct {
    gorm.Model
    Name               string `json:"name" gorm:"size:200"`
    CreatedBy          uint   `json:"createdby"`
}

type Business struct {
  gorm.Model
  BusinessGroupID uint   `json:"businessgroupid"`
  Name           string  `json:"name" gorm:"size:100"`
  AboutSummary   string  `json:"aboutsummary" gorm:"size:200"`
  Logo           string  `json:"logo" gorm:"size:200"`
  Phone          string  `json:"phone" gorm:"size:15"`
  Email          string  `json:"email" gorm:"size:60"`
  Industry        string `json:"industry" gorm:"size:60"`
  PhysicalAddress string `json:"physicaladdress" gorm:"size:120"`
  Website        string  `json:"website" gorm:"size:100"`
  Facebook       string  `json:"facebook" gorm:"size:100"`
  Twitter        string  `json:"twitter" gorm:"size:100"`
  Instagram      string  `json:"instagram" gorm:"size:100"`
  Linkedin       string  `json:"linkedin" gorm:"size:100"`
  Telegram       string  `json:"telegram" gorm:"size:100"`
  Whatsapp       string  `json:"whatsapp" gorm:"size:100"`
  Slack          string  `json:"slack" gorm:"size:100"`
  Discord        string  `json:"discord" gorm:"size:100"`
  CreatedBy      uint    `json:"createdby"`
  Partners       []Business `json:"partners"`
}

type BranchGroup struct {
    gorm.Model
    Name               string `json:"name" gorm:"size:200"`
    BusinessID         uint   `json:"businessid"`
    CreatedBy          uint   `json:"createdby"`
}

type Branch struct {
  gorm.Model
  BusinessID     uint     `json:"businessid"`
  Name           string   `json:"name" gorm:"size:100"`
  AboutSummary   string   `json:"aboutsummary" gorm:"size:200"`
  Logo           string   `json:"logo" gorm:"size:200"`
  Phone          string   `json:"phone" gorm:"size:15"`
  Email          string   `json:"email" gorm:"size:60"`
  PhysicalAddress string  `json:"physicaladdress" gorm:"size:120"`
  Website        string   `json:"website" gorm:"size:100"`
  Facebook       string   `json:"facebook" gorm:"size:100"`
  Twitter        string   `json:"twitter" gorm:"size:100"`
  Instagram      string   `json:"instagram" gorm:"size:100"`
  Linkedin       string   `json:"linkedin" gorm:"size:100"`
  Telegram       string   `json:"telegram" gorm:"size:100"`
  Whatsapp       string   `json:"whatsapp" gorm:"size:100"`
  Slack          string   `json:"slack" gorm:"size:100"`
  Discord        string   `json:"discord" gorm:"size:100"`
  CreatedBy      uint     `json:"createdby"`
  Partners       []Branch `json:"partners"`
}

type BranchJoinGroup {
    BranchID         uint   `json:"branchid"`
    BranchGroupID    uint   `json:"branchgroupid"`
}

type Job struct {
    gorm.Model
    BusinessID         uint    `json:"businessid"`
    BranchID           uint    `json:"branchid"`
    Title              string  `json:"title" gorm:"size:200"`
    Type               string  `json:"type" gorm:"size:30;default:full-time"`
    Wage               float32 `json:"wage"`
    WageDuration       string  `json:"wageduration" gorm:"size:15;default:per month"`
    Description        string  `json:"description" gorm:"size:2000"`
    RequiredExperience string  `json:"requiredexperience" gorm:"size:60"`
    Hiring             bool    `json:"hiring"      gorm:"default:false"`
    HiringSlots        bool    `json:"hiringslots" gorm:"default:0"`
    ApplicationLink    string  `json:"applicationlink" gorm:"size:100"`
    CreatedBy          uint    `json:"createdby"`
    RequiredSkills     []Skill `json:"requiredskills"`
    AdvantageSkills    []Skill `json:"advantageskills"`
    Eligibility        []JobCritereon `json:"eligibility"`
    Responsibilities   []JobRole `json:"responsibilities"`
}

type Skill struct {
    gorm.Model
    Name               string `json:"name" gorm:"size:30;unique"`
}

type JobCritereon struct {
    gorm.Model
    Name               string `json:"name" gorm:"size:100;unique"`
}

type JobRole struct {
    gorm.Model
    Name               string `json:"name" gorm:"size:100;unique"`
}

type Department struct {
    gorm.Model
    Name               string `json:"name" gorm:"size:200"`
    BusinessID         uint   `json:"businessid"`
    BranchID           uint   `json:"branchid"`
    CreatedBy          uint   `json:"createdby"`
}

type Service struct {
  gorm.Model
  Name               string  `json:"name" gorm:"size:200"`
  Category           string  `json:"category" gorm:"size:40"`
  Availability       bool    `json:"availability" gorm:"default:true"`
  Price              uint    `json:"price"`
  Department         uint    `json:"department"`
  ServiceProviderID  uint    `json:"provider"`
  CreatedBy          uint    `json:"createdby"`
  BusinessID         uint    `json:"businessid"`
  BranchID           uint    `json:"branchid"`
}

type Product struct {
  gorm.Model
  Name             string  `json:"name" gorm:"size:150"`
  Category         string  `json:"category" gorm:"size:40"`      //soda
  Brand            string  `json:"brand" gorm:"size:20"`         //fanta, crest, etc
  Type             string  `json:"type" gorm:"size:20"`          //plastic, bottled, etc
  Price            uint    `json:"price"`                        //4000
  Department       uint    `json:"department"`
  Quantity         uint    `json:"quantity"`                     //450
  QuantityUnits    string  `json:"quantityUnits" gorm:"size:20"` //crates
  SupplierID       uint    `json:"supplierid"`
  CreatedBy        uint    `json:"createdby"`
  BusinessID       uint    `json:"businessid"`
  BranchID         uint    `json:"branchid"`
}

type Package struct {
  gorm.Model
  Name          string            `json:"name" gorm:"size:100"`
  Category      string            `json:"category" gorm:"size:20;default:guests"`
  Availability  bool              `json:"availability" gorm:"default:true"`
  Price         uint              `json:"price"`
  Department    uint              `json:"department"`
  Products      []PackageProduct  `json:"products"`
  Services      []PackageService  `json:"services"`
  CreatedBy     uint              `json:"createdby"`
  BusinessID    uint              `json:"businessid"`
  BranchID      uint              `json:"branchid"`
}

type Decorator struct {
  gorm.Model
  Name             string  `json:"name" gorm:"size:150"`
  Price            uint    `json:"price"`
  CreatedBy        uint    `json:"createdby"`
  BusinessID       uint    `json:"businessid"`
  BranchID         uint    `json:"branchid"`
}

type PackageProduct struct {
  gorm.Model
  ProductID   uint    `json:"productid"`
  PackageID   uint    `json:"packageid"`
  Quantity    uint    `json:"quantity"`
  QuantityServed uint `json:"quantityserved"`
  QuantityPaid uint   `json:"quantitypaid"`
  BusinessID  uint    `json:"businessid"`
  BranchID    uint    `json:"branchid"`
}

type PackageService struct {
  gorm.Model
  ServiceID   uint  `json:"serviceid"`
  PackageID   uint  `json:"packageid"`
  Quantity    uint  `json:"quantity"`
  QuantityServed uint `json:"quantityserved"`
  QuantityPaid uint   `json:"quantitypaid"`
  BusinessID  uint  `json:"businessid"`
  BranchID    uint  `json:"branchid"`
}

type Visit struct {
  gorm.Model
  Customer    User     `json:"customer" gorm:"foreignKey:CustomerID;references:id"`
  CustomerID  uint     `json:"customerid"`
  CreatedBy   uint     `json:"createdby"`
  BusinessID  uint     `json:"businessid"`
  BranchID    uint     `json:"branchid"`
}

type OrderOrBooking struct {
  gorm.Model
  Item       string   `json:"item" gorm:"size:8"`           //product, service, package
  ItemID     uint     `json:"itemid"`
  Quantity   uint     `json:"quantity"`
  Status     string   `json:"status" gorm:"size:20"`        //pending, invoiced, served, billed, cancelled
  Paid       bool     `json:"paid" gorm:"default:false"`    //true/false
  Customer   User     `json:"customer" gorm:"foreignKey:CustomerID;references:id"`
  Visit      Visit    `json:"visit"`
  CustomerID uint     `json:"customerid"`
  VisitID    uint     `json:"visitid"`
  InvoiceID  uint     `json:"invoiceid"`
  BillID     uint     `json:"billid"`
  CreatedBy  uint     `json:"createdby"`                    //id of the user who created a record
  BusinessID uint     `json:"businessid"`
  BranchID   uint     `json:"branchid"`
}

type Invoice struct {
  gorm.Model
  Amount     uint             `json:"amount"`
  Status     string           `json:"status" gorm:"size:20;default:pending"` //pending, billed, paid, cancelled
  Customer   User             `json:"customer" gorm:"foreignKey:CustomerID;references:id"`
  CustomerID uint             `json:"customerid"`
  BillID     uint             `json:"billid"`
  CreatedBy  uint             `json:"createdby"`                             //id of the user who created a record
  Orders     []OrderOrBooking `json:"orders"`
  BusinessID uint             `json:"businessid"`
  BranchID   uint             `json:"branchid"`
}

type Bill struct {
  gorm.Model
  Amount     uint             `json:"amount"`
  Status     string           `json:"status" gorm:"size:20;default:pending"` //pending, paid, cancelled
  Customer   User             `json:"customer" gorm:"foreignKey:CustomerID;references:id"`
  CustomerID uint             `json:"customerid"`
  Invoices   string           `json:"invoices" gorm:"size:200;default:none"` //list of comma-separated, stringified invoice IDs 
  CreatedBy  uint             `json:"createdby"`                             //id of the user who created a record
  Orders     []OrderOrBooking `json:"orders"`  
  BusinessID uint             `json:"businessid"`
  BranchID   uint             `json:"branchid"` 
}

type Payment struct {
  gorm.Model
  Amount     uint     `json:"amount"`
  Item       string   `json:"item" gorm:"size:8"`              //order, booking, invoice, bill
  ItemID     uint     `json:"itemid"`
  Instalment bool     `json:"instalment" gorm:"default:false"` //true/false
  Customer   User     `json:"customer" gorm:"foreignKey:CustomerID;references:id"`
  CustomerID uint     `json:"customerid"`
  CreatedBy  uint     `json:"createdby"`                       //id of the user who created a record
  BusinessID uint     `json:"businessid"`
  BranchID   uint     `json:"branchid"`
}

type StockTransaction struct {
  gorm.Model
  Transaction      string   `json:"transaction" gorm:"size:8;default:add"`  //add, remove,
  ProductCategory   string  `json:"productCategory" gorm:"size:150"`        //
  ProductID         uint    `json:"productid"`
  OldQuantity       uint    `json:"oldQuantity"`
  Quantity          uint    `json:"quantity"`
  NewQuantity       uint    `json:"newQuantity"`
  Returned          bool    `json:"returned" gorm:"default:false"`     //true/false
  Amount            uint    `json:"amount"`
  CreatedBy         uint    `json:"createdby"`
  BusinessID        uint    `json:"businessid"`
  BranchID          uint    `json:"branchid"`
}

type Expense struct {
  gorm.Model
  Amount        uint    `json:"amount"`
  SpentOn       string  `json:"spenton"`     //stock, petty-cash, regular, ...
  Description   string  `json:"description"`
  CreatedBy     uint    `json:"createdby"`
  BusinessID    uint    `json:"businessid"`
  BranchID      uint    `json:"branchid"`
}

type User struct {
  gorm.Model
  ProfileID          uint           `json:"profileid"`
  Username           string         `json:"username" gorm:"size:20"`
  Password           string         `json:"pxwd" gorm:"size:150"`
  Firstname          string         `json:"firstname" gorm:"size:20"`
  Lastname           string         `json:"lastname" gorm:"size:20"`
  Phone              string         `json:"phone" gorm:"size:20"`
  Email              string         `json:"email" gorm:"size:100"`
  Address            string         `json:"address" gorm:"size:60"`
  UserType           string         `json:"userType" gorm:"default:customer;size:20"` //staff, customer, guest, partner, service_provider, supplier...
  StaffTitle         string         `json:"stafftitle" gorm:"size:20"`
  IdentityCardNumber string         `json:"idcardnumber" gorm:"size:20"`
  IdentityCardType   string         `json:"idtype" gorm:"default:EMPLOYEE_ID" gorm:"size:20"`
  Nationality        string         `json:"nationality" gorm:"size:50"`
  ProfilePicture     string         `json:"profilepicture" gorm:"size:150;default:xxxxx"`
  AccessRights       uint           `json:"accessRights"`
  CreatedBy          uint           `json:"createdby"`
  BusinessID         uint           `json:"businessid"`
  BranchID           uint           `json:"branchid"`
  Invoices           []Invoice      `json:"invoices" gorm:"foreignKey:CustomerID;references:id;->"`
  Bills              []Bill         `json:"bills" gorm:"foreignKey:CustomerID;references:id;->"`
  Messages           []Message      `json:"messages" gorm:"foreignKey:UserID;references:id;->"`
  Notifications      []Notification `json:"notifications" gorm:"foreignKey:UserID;references:id;->"`
  UserActions        []UserAction   `json:"userActions" gorm:"foreignKey:UserID;references:id;->"`
}

type UserAction struct {
  gorm.Model
  ActionNumber   string  `json:"actionNumber" gorm:"size:16"`
  UserID         uint    `json:"userid"`
  ProfileID      uint    `json:"profileid"`
  Description    string  `json:"description" gorm:"size:100"`
  OnEntity       string  `json:"onEntity" gorm:"size:20"`
  SpecificEntity string  `json:"specificEntity" gorm:"size:100"`
  BusinessID     uint    `json:"businessid"`
  BranchID       uint    `json:"branchid"`
}

type Message struct {
  gorm.Model
  Text       string    `json:"text"`
  Read       bool      `json:"read" gorm:"default:false"`
  UserID     uint      `json:"userid"`
  ProfileID  uint      `json:"profileid"`
  CreatedBy  uint      `json:"createdby"`
  BusinessID uint      `json:"businessid"`
  BranchID   uint      `json:"branchid"`
}

type Notification struct {
  gorm.Model
  Message    string
  Read       bool
  UserID     uint
  ProfileID  uint     `json:"profileid"`
  CreatedBy  uint     `json:"createdby"`
  BusinessID uint     `json:"businessid"`
  BranchID   uint     `json:"branchid"`
}


type Profile struct {
  gorm.Model
  Firstname          string         `json:"firstname" gorm:"size:20"`
  Lastname           string         `json:"lastname" gorm:"size:20"`
  Phone              string         `json:"phone" gorm:"size:20"`
  Email              string         `json:"email" gorm:"size:100"`
  Address            string         `json:"address" gorm:"size:60"`
  UserType           string         `json:"userType" gorm:"default:user;size:10"` //developer, user
  IdentityCardNumber string         `json:"idcardnumber" gorm:"size:20"`
  IdentityCardType   string         `json:"idtype" gorm:"default:EMPLOYEE_ID" gorm:"size:20"`
  Nationality        string         `json:"nationality" gorm:"size:50"`
  ProfilePicture     string         `json:"profilepicture" gorm:"size:150;default:xxxxx"`
  CreatedBy          uint           `json:"createdby"`
  Businesses         []Business     `json:"businesses"`
  Branches           []Branch       `json:"branches"`
  Jobs               []Job          `json:"jobs"`
  Messages           []Message      `json:"messages"`
  Friends            []Profile      `json:"friends"`
}

//=============================wagepad================================
type Employee struct {
  gorm.Model
  BusinessID         uint           `json:"businessid"`
  BranchID           uint           `json:"branchid"`
  JobID              uint           `json:"jobid"`
  ProfileID          uint           `json:"profileid"`
  UserID             uint           `json:"userid"`
  Title              string         `json:"title" gorm:"size:20"`
  EmploymentType     string         `json:"employmenttype" gorm:"size:20;default:full-time"`
  GrossPay           float64        `json:"grosspay"`
  PayDuration        string         `json:"payduration" gorm:"size:20;default:per-month"`
  Currency           string         `json:"currency" gorm:"size:4;default:USD"`
  BankAccount        string         `json:"bankaccount" gorm:"size:100"`
  BankName           string         `json:"bankname" gorm:"size:100"`
  BankBranch         string         `json:"bankbranch" gorm:"size:100"`
  Budgetline         string         `json:"budgetline" gorm:"size:100"`
  StartDate          string         `json:"startdate" gorm:"size:30"`
  DutyStation        string         `json:"dutystation" gorm:"size:30"`
  WeeklyHours        uint           `json:"weeklyhours" gorm:"default:40"`
  Status             string         `json:"status" gorm:"default:current"`        //current, left, suspended, expelled
  ExtraTimeAllowance float64        `json:"extratimeallowance"`
  NoKin              string         `json:"address" gorm:"size:60"`
  NoKinPhone         string         `json:"Nokinphone" gorm:"default:user;size:20"`
  NoKinEmail         string         `json:"nokinemail" gorm:"size:60"`
  NoKinRelationship  string         `json:"nokinrelationship" gorm:"size:30"`
}

type Project struct {
  gorm.Model
  BusinessID         uint           `json:"businessid"`
  BranchID           uint           `json:"branchid"`
  Name               string         `json:"name" gorm:"size:100"`
  StartDate          string         `json:"startdate" gorm:"size:30"`
  EndDate            string         `json:"enddate" gorm:"size:30"`
  OperationArea      string         `json:"operationarea" gorm:"size:60"`
  ProgresStatus      string         `json:"progresstatus" gorm:"size:10;default:planned"` //planned, started, 1/4, 1/2, 3/4, completed.
  CurrentStatus      string         `json:"currentstatus" gorm:"size:10;default:paused"`  //waiting, running, paused, complete 
  Sponsors           string         `json:"description" gorm:"size:100;default:inhouse"`   //comma separated list of sponsors
  Budget             float64        `json:"budget"`
  Currency           string         `json:"currency" gorm:"size:4;default:USD"`
  Description        string         `json:"description" gorm:"size:200"`
}

type Event struct {
  gorm.Model
  BusinessID         uint           `json:"businessid"`
  BranchID           uint           `json:"branchid"`
  Name               string         `json:"name" gorm:"size:100"`
  StartDate          string         `json:"startdate" gorm:"size:30"`
  EndDate            string         `json:"enddate" gorm:"size:30"`
  Venue              string         `json:"venue" gorm:"size:60"`
  ProgresStatus      string         `json:"progresstatus" gorm:"size:10;default:starting"` //planned, started, 1/4, 1/2, 3/4, completed.
  CurrentStatus      string         `json:"currentstatus" gorm:"size:10;default:paused"`   //waiting, running, paused, complete 
  Sponsors           string         `json:"description" gorm:"size:100;default:inhouse"`   //comma separated list of sponsors
  Budget             float64        `json:"budget"`
  Currency           string         `json:"currency" gorm:"size:4;default:USD"`
  Description        string         `json:"description" gorm:"size:200"`
}

type Management struct {
  gorm.Model
  BusinessID         uint           `json:"businessid"`
  BranchID           uint           `json:"branchid"`
  EmployeeID         uint           `json:"employeeid"`
  ManagerOf          string         `json:"managerof" gorm:"size:20"`         //project, event, department, branch, ...
  ManagedEntityID    uint           `json:"managedentityid"`
  ManagementRole     string         `json:"managementrole" gorm:"size:30"`    //supervisor, project-coordinator, ...
}

type Task struct {
  gorm.Model
  BusinessID         uint           `json:"businessid"`
  BranchID           uint           `json:"branchid"`
  Supervisor         uint           `json:"supervisor"`                           //employeeid
  Name               string         `json:"name" gorm:"size:100"`
  Frequency          string         `json:"frequency" gorm:"size:30;default:one-off"`  //one-off, daily, weekly, monthly
  Description        string         `json:"description" gorm:"size:200"`
  Deadline           string         `json:"deadline" gorm:"size:30"`                  //date string
  Allowance          float64        `json:"allowance" gorm:"default:0.00"`
  Commision          float64        `json:"commision" gorm:"default:0.00"`
  Currency           string         `json:"currency" gorm:"size:4;default:USD"`
}

type Activity struct {
  gorm.Model
  BusinessID         uint           `json:"businessid"`
  BranchID           uint           `json:"branchid"`
  EmployeeID         uint           `json:"employeeid"`
  TimeWindowFrom     string         `json:"timewindowfrom"`
  EmployeeStatus     string         `json:"employeestatus"`
  StatusFrom         string         `json:"statusfrom"`
  ProjectID          uint           `json:"projectid"`
  EventID            uint           `json:"eventid"`
  TaskID             uint           `json:"taskid"`
  TaskFrom           string         `json:"taskfrom"`
  DeploymentStation  string         `json:"deploymentstation"`
  HoursOnDutty       uint           `json:"hoursondutty"`
  HoursOnLeave       uint           `json:"hoursonleave"`
  HoursOnOffWithPermission uint     `json:"hoursoffwithpermission"`
}

type Wage struct {
  gorm.Model
  BusinessID         uint           `json:"businessid"`
  BranchID           uint           `json:"branchid"`
  EmployeeID         uint           `json:"employeeid"`
  TimeWindowFrom     string         `json:"timewindowfrom" gorm:"size:30"`
  TimeWindowTo       string         `json:"timewindowto" gorm:"size:30"`
  WageType           string         `json:"wagetype" gorm:"size:15;default:salary"` //salary, commission, allowance, apprisal
  Description        string         `json:"description" gorm:"size:100"`
  Amount             float64        `json:"amount"`
  PayDurationUnits   float64        `json:"paydurationunits" gorm:"default:1.00"`
  PayRate            float64        `json:"payrate"`
  Currency           string         `json:"currency" gorm:"size:4;default:USD"`
  Status             string         `json:"status" gorm:"size:10;default:due"`     //due, paid, omitted
  ActivityLevel      float64        `json:"activitylevel"`                         //eg, 0.80 : the percentage of time on duty out of time supposed to be on duty
}

type ChangeHistory struct {
  gorm.Model
  BusinessID         uint           `json:"businessid"`
  BranchID           uint           `json:"branchid"`
  EmployeeID         uint           `json:"employeeid"`
  ChangeOf           string         `json:"changeof" gorm:"size:30"`  //employeestatus, activitystatus, task, deploymentstation
  ChangeTo           string         `json:"changeto" gorm:"size:30"`
  Description        string         `json:"description" gorm:"size:100"`
}


/*
Ref: https://gorm.io/docs/models.html
NOTE ignored fields won’t be created when using GORM Migrator to create table

type User struct {
  Name string `gorm:"<-:create"` // allow read and create
  Name string `gorm:"<-:update"` // allow read and update
  Name string `gorm:"<-"`        // allow read and write (create and update)
  Name string `gorm:"<-:false"`  // allow read, disable write permission
  Name string `gorm:"->"`        // readonly (disable write permission unless it configured )
  Name string `gorm:"->;<-:create"` // allow read and create
  Name string `gorm:"->:false;<-:create"` // createonly (disabled read from db)
  Name string `gorm:"-"`  // ignore this field when write and read with struct
}
*/ 
  
  
  





