package delivery

import (
	"Clean-Architecture/entities"
	error1 "Clean-Architecture/errorPackage"
	_ "Clean-Architecture/service"
	service "Clean-Architecture/service"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	body, _ := json.Marshal(resp)
	_, _ = w.Write(body)
}
func (a CustomerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	resp, err := a.service.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
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
		_, _ = w.Write([]byte(error1.InvalidId))
		return
	}
	resp, err := a.service.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
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
		_, _ = w.Write([]byte(error1.MissingBodyJson))
		return
	}
	err := json.Unmarshal(body, &customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(error1.JsonParsingError))
		return
	}
	resp, err := a.service.Create(customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
	body, _ = json.Marshal(resp)
	_, _ = w.Write(body)
}
func (a CustomerHandler) Edit(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	var customer entities.Customer
	body, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(error1.JsonParsingError))
		return
	}
	resp, err := a.service.Edit(id, customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
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
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	resp, err := a.service.DeleteById(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return

		return
	} else {
		w.WriteHeader(http.StatusNoContent)
		body, _ := json.Marshal(resp)
		_, _ = w.Write(body)
	}
}
