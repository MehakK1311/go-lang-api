package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"learning/models"
	"log"
	"os"
	"strconv"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type response struct {
	ID int64 `json:"id,omitempty"`
	Message string `json:"message,omitempty"`

}


func createConnection() *sql.DB{
	err:= godotenv.Load(".env")

	if err != nil{
		log.Fatal("Error loading .env file")
	}

	db, err:= sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err!=nil{
		panic(err)
	}

	err = db.Ping()

	if err!=nil{
		panic(err)
	}

	fmt.Println("Successfully connected to postgres")

	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request){

	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err!=nil{
		log.Fatal("Unable to decode the request body. %v", err)

	}
	

	insertID:=insertStock()

	res:=response{
		ID: insertID,
		Message: "stock created sucessfully",
	}

	json.NewEncoder(w).Encode(res)
}

func GetStock(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	id, err:=strconv.Atoi(params["id"])

	if err!=nil{
		log.Fatalf("unable to convert the string into int %v", err)
	}

	stock, err := getStock(int64(id))

	if err!= nil{
		log.Fatalf("unable to get stock %v", err)
	}

	json.NewEncoder(w).Encode(stock)

}

func GetAllStock(w http.ResponseWriter, r *http.Request){
	stocks, err:= getAllStock()

	if err!=nil{
		log.Fatalf("Unabe to get all the stocks %v", err)
	}

	json.NewEncoder(w).Encode(stocks)
}

func UpdateStock(w http.ResponseWriter, r *http.Request){
	params:=mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err!=nil{
		log.Fatalf("Unabe to convert the string into int. %v", err)
	}

	var stock models.Stock

	err = json.NewDecoder(r.Body).Decode(&stock)

	if err!=nil{
		log.Fatalf("Unabe to decode the request body. %v", err)
	}	

	updateRows:= updateStock(int64(id), stock)

	msg:= fmt.Sprintf("Stock updated successfully. Total rows/records affected %v", updateRows)

	res:= response{
		ID: int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeleteStock(w http.ResponseWriter, r *http.Request){
	params:=mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err!=nil{
		log.Fatalf("Unabe to convert the string into int. %v", err)
	}

	deleteRows:= deleteStock(int64(id))

	msg:= fmt.Sprintf("Stock deleted successfully. Total rows/records affected %v", deleteRows)

	res:= response{
		ID: int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
