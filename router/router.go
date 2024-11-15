package router

import (
	"learning/middleware"
	"github.com/gorilla/mux"
)

func Router () *mux.Router{

	router := mux.NewRouter()

	router.HandleFunc("/api/stocks/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stocks", middleware.GetAllStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newStock", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/deletestock/{id}", middleware.DeleteStock).Methods("DELETE","OPTIONS")
}