package main


import (
  "fmt"
  "log"
  "flag"
  "embed"
  "io/fs"
  "strings"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  uc "github.com/mahani-software-engineering/bms-server/usecases"
)

//go:embed client/web/*
var static embed.FS

func htmlWebsite(w http.ResponseWriter, r *http.Request) {
	website, _ := fs.Sub(static, "client")
    handler := http.FileServer(http.FS(website))
    handler.ServeHTTP(w, r)
}

func index(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(struct{Success string}{Success: "API home"})
}

func resourceNotFound(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(struct{Success string}{Success: "The API doesn't have what you are looking for !"})
}

func getRouter() *mux.Router {
	website, _ := fs.Sub(static, "client/web")
	router := mux.NewRouter()
	//User
	router.HandleFunc("/user/register", uc.Signup).Methods("POST")
	router.HandleFunc("/user/login", uc.UserLogin).Methods("POST")
	//Doctor
	router.HandleFunc("/doctor", uc.CreateDoctor).Methods("POST")
	router.HandleFunc("/doctor/{id}", uc.UpdateDoctor).Methods("PUT")
	router.HandleFunc("/doctor/{id}", uc.ReadDoctor).Methods("GET")
	router.HandleFunc("/doctor", uc.ReadAllDoctors).Methods("GET")
	//Pharmacist
	router.HandleFunc("/pharmacist", uc.CreatePharmacist).Methods("POST")
	router.HandleFunc("/pharmacist/{id}", uc.UpdatePharmacist).Methods("PUT")
	router.HandleFunc("/pharmacist/{id}", uc.ReadPharmacist).Methods("GET")
	router.HandleFunc("/pharmacist", uc.ReadAllPharmacists).Methods("GET")
	//Patient
	router.HandleFunc("/patient", uc.CreateDoctor).Methods("POST")
	router.HandleFunc("/patient/{id}", uc.UpdateDoctor).Methods("PUT")
	router.HandleFunc("/patient/{id}", uc.ReadDoctor).Methods("GET")
	router.HandleFunc("/patient", uc.ReadAllDoctors).Methods("GET")
	//DeliveryAgent
	router.HandleFunc("/deliveryagent", uc.CreateDeliveryAgent).Methods("POST")
	router.HandleFunc("/deliveryagent/{id}", uc.UpdateDeliveryAgent).Methods("PUT")
	router.HandleFunc("/deliveryagent/{id}", uc.ReadDeliveryAgent).Methods("GET")
	router.HandleFunc("/deliveryagent", uc.ReadAllDeliveryAgents).Methods("GET")
	//Store
	router.HandleFunc("/store", uc.CreateStore).Methods("POST")
	router.HandleFunc("/store/{id}", uc.UpdateStore).Methods("PUT")
	router.HandleFunc("/store/{id}", uc.ReadStore).Methods("GET")
	router.HandleFunc("/store", uc.ReadAllStores).Methods("GET")
	//Medicine
	router.HandleFunc("/medicine", uc.CreateMedicine).Methods("POST")
	router.HandleFunc("/medicine/{id}", uc.UpdateMedicine).Methods("PUT")
	router.HandleFunc("/medicine/{id}", uc.ReadMedicine).Methods("GET")
	router.HandleFunc("/medicine", uc.ReadAllMedicines).Methods("GET")
	//Prescription
	router.HandleFunc("/prescription", uc.CreatePrescription).Methods("POST")
	router.HandleFunc("/prescription/{id}", uc.UpdatePrescription).Methods("PUT")
	router.HandleFunc("/prescription/{id}", uc.ReadPrescription).Methods("GET")
	router.HandleFunc("/prescription", uc.ReadAllPrescriptions).Methods("GET")
	//Payment
	router.HandleFunc("/payment", uc.CreatePayment).Methods("POST")
	router.HandleFunc("/payment/{id}", uc.UpdatePayment).Methods("PUT")
	router.HandleFunc("/payment/{id}", uc.ReadPayment).Methods("GET")
	router.HandleFunc("/payment", uc.ReadAllPayments).Methods("GET")
	//Order
	router.HandleFunc("/order", uc.CreateOrder).Methods("POST")
	router.HandleFunc("/order/{id}", uc.UpdateOrder).Methods("PUT")
	router.HandleFunc("/order/{id}", uc.ReadOrder).Methods("GET")
	router.HandleFunc("/order", uc.ReadAllOrders).Methods("GET")
	//Comment
	router.HandleFunc("/comment", uc.CreateComment).Methods("POST")
	router.HandleFunc("/comment/{id}", uc.UpdateComment).Methods("PUT")
	router.HandleFunc("/comment/{id}", uc.ReadComment).Methods("GET")
	router.HandleFunc("/comment", uc.ReadAllComments).Methods("GET")
	//Delivery
	router.HandleFunc("/delivery", uc.CreateDelivery).Methods("POST")
	router.HandleFunc("/delivery/{id}", uc.UpdateDelivery).Methods("PUT")
	router.HandleFunc("/delivery/{id}", uc.ReadDelivery).Methods("GET")
	router.HandleFunc("/delivery", uc.ReadAllDeliverys).Methods("GET")
	//Message
	router.HandleFunc("/message", uc.CreateMessage).Methods("POST")
	router.HandleFunc("/message/{id}", uc.UpdateMessage).Methods("PUT")
	router.HandleFunc("/message/{id}", uc.ReadMessage).Methods("GET")
	router.HandleFunc("/message", uc.ReadAllMessages).Methods("GET")
	
	//Home
	router.HandleFunc("/", index ).Methods("POST")
	router.PathPrefix("/").Handler( http.FileServer(http.FS(website)) ).Methods("GET")
	
	//Not found
	router.NotFoundHandler = http.HandlerFunc(resourceNotFound)
	
	return router
}

func main() {
    //++++| os.Args |+++++
    wsEndPoint := ":5600" 
    addr := flag.String("addr", wsEndPoint, "E-Pharmacy API service address") 
    flag.Parse()          
    //++++++++++++++++++++
    uc.Init()
    
    fmt.Println("Server listening on port: "+(strings.Split(wsEndPoint,":")[1])) 
    log.Fatal(http.ListenAndServe(*addr, getRouter()))
}


