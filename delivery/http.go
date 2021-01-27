package delivery

import (
	"Clean-Architecture/entities"
	_ "Clean-Architecture/service"
	service "Clean-Architecture/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type CustomerHandler struct {
	service service.Customers
}
func New(serv service.Customers) CustomerHandler {
	return CustomerHandler{service: serv}
}
func (a CustomerHandler) GetByName(w http.ResponseWriter, r *http.Request) {
	        name := r.URL.Query()["name"]
			resp, err := a.service.GetByName(name[0])
			if err != nil {
				_, _ = w.Write([]byte("could not get customer information"))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			body, _ := json.Marshal(resp)
			_, _ = w.Write(body)
}
func (a CustomerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
		resp, err := a.service.GetAll()
		if err != nil {
			_, _ = w.Write([]byte("could not get customer information"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		body, _ := json.Marshal(resp)
		_, _ = w.Write(body)
	}
func (a CustomerHandler) GetById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("invalid id"))

		return
	}
	resp, err := a.service.GetById(id)
	if err != nil {
		_, _ = w.Write([]byte("could not retrieve customer"))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
	body, _ := json.Marshal(resp)
	_, _ = w.Write(body)
}
func (a CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var customer entities.Customer
	body, _ := ioutil.ReadAll(r.Body)
	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}else {

		err := json.Unmarshal(body, &customer)
		if err != nil {
			fmt.Println(err)
			_, _ = w.Write([]byte("invalid body"))
			w.WriteHeader(http.StatusBadRequest)

			return
		}
		resp, err := a.service.Create(customer)
		if err != nil {
			_, _ = w.Write([]byte("could not create customer"))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		body, _ = json.Marshal(resp)
		_, _ = w.Write(body)
	}
}
func (a CustomerHandler) Edit(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var customer entities.Customer
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &customer)
	if err != nil {
		fmt.Println(err)
		_, _ = w.Write([]byte("invalid body"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := a.service.Edit(id, customer)
	if err != nil {
		_, _ = w.Write([]byte("age is less than 18"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	body, _ = json.Marshal(resp)
	_, _ = w.Write(body)
}
func (a CustomerHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	resp, err := a.service.DeleteById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("No data to delete customer"))

		return
	}else {
		w.WriteHeader(http.StatusNoContent)
		body, _ := json.Marshal(resp)
		_, _ = w.Write(body)
	}
}

