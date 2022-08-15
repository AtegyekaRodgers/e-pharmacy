package usecases

const (
    _ = 1<<iota
    IsSystemAdmin      uint32 = 1<<iota             //00000000000000000000000000000010        2
    IsStaffAdmin       uint32 = 1<<iota             //00000000000000000000000000000100        4
    IsStaff            uint32 = 1<<iota             //00000000000000000000000000001000        8
    IsCustomer         uint32 = 1<<iota             //00000000000000000000000000010000       16
    IsVisitor          uint32 = 1<<iota             //00000000000000000000000000100000       32
    CanRegisterUser           uint32 = 1<<iota      //00000000000000000000000001000000       64
    CanRegisterCustomer       uint32 = 1<<iota      //00000000000000000000000100000000      256
    CanViewRevenueBanner      uint32 = 1<<iota      //00000000000000000000001000000000      512
    CanViewCustomersBanner    uint32 = 1<<iota      //00000000000000000000010000000000     1024
    CanViewGuestsBanner       uint32 = 1<<iota      //00000000000000000000100000000000     2048
    CanViewOrdersBanner       uint32 = 1<<iota      //00000000000000000001000000000000     4096
    CanViewExpensesBanner     uint32 = 1<<iota      //00000000000000000010000000000000     8192
    CanViewPaymentsBanner     uint32 = 1<<iota      //00000000000000000100000000000000    16384
    CanGraphCustomersThisWeek uint32 = 1<<iota      //00000000000000001000000000000000    32768
    CanGraphRevenueThisWeek   uint32 = 1<<iota      //00000000000000010000000000000000    65536
    CanGraphAnnualRevenue     uint32 = 1<<iota      //00000000000000100000000000000000   131072
    CanViewStockTransactions  uint32 = 1<<iota      //00000000000001000000000000000000   262144
)


/*
	RegisteringUser
	RegisterNewGuest 
	
	
	AssignUserRights
	ReadUser         
	ReadAllUsers
	ReadCountCustomers
	ReadAllCustomers
	ReadUser
	ReadCountGuests
	ReadAllGuests
	ReadUser
	CreateProduct
	UpdateProduct
	ReadProduct
	ReadAllProducts
	CreateService
	UpdateProduct
	ReadService
	ReadAllServices
	CreatePackage
	UpdatePackage
	ReadPackage
	ReadAllPackages
	CreateVisit
	UpdateVisit
	ReadVisit
	ReadAllVisits
	CreateOrderOrBooking
	UpdateOrderOrBooking
	ReadOrderOrBooking
	ReadAllOrdersOrBookings
	ReadOrdersByCustomer
	ReadOrdersByInvoice
	ReadOrdersByBill
	ReadCountOrdersOrBookings
	StatsCustomersThisWeek
	StatsRevenueThisWeek
	StatsRevenueSrcThisWeek
	StatsRevenueSrcAnnual
	StatsAnnualRevenue
	CreatePayment
	UpdatePayment
	ReadPayment
	ReadCountPayments
	ReadPaymentsTotalAmount
	ReadAllPayments
	CreateInvoice
	ReadInvoice
	ReadAllInvoices
	CreateBill
	ReadBill
	ReadAllBills
	CreateStockTransaction
	UpdateStockTransaction
	ReadStockTransaction
	ReadAllStockTransactions
	ReadStockTxnsForBarner
	CreateExpense
	UpdateExpense
	ReadExpense
	ReadAllExpenses
	ReadExpensesTotalAmount
	ReadUserAction
	ReadAllUserActions
	CreateMessage
	UpdateMessage
	ReadMessage
	ReadAllMessages
	CreateNotification
	UpdateNotification
	ReadNotification
	ReadAllNotifications
*/





