package router

import (
	"github.com/gorilla/mux"
	"github.com/rinerte/sharma-go-postgres/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.GetAllStock).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/stock", middleware.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/stock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")
	return router
}
