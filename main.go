package main

import (
	"Clean-Architecture/delivery"
	"Clean-Architecture/service"
	"Clean-Architecture/store/customer"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main()  {
	r := mux.NewRouter()
	datastore := store.New()
    defer datastore.CloseDB()
	service := service.New(datastore)
	handler := delivery.New(service)
	r.HandleFunc("/customer",handler.GetByName).Methods(http.MethodGet).Queries("name","{name}")
	r.HandleFunc("/customer",handler.GetAll).Methods(http.MethodGet)
	r.HandleFunc("/customer/{id}", handler.GetById).Methods(http.MethodGet)
	r.HandleFunc("/customer", handler.Create).Methods(http.MethodPost)
	r.HandleFunc("/customer/{id}", handler.Edit).Methods(http.MethodPut)
	r.HandleFunc("/customer/{id}", handler.DeleteById).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", r))
}
