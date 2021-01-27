package delivery

import (
	"Clean-Architecture/entities"
	"bytes"
	"errors"
	"github.com/bearbin/go-age"
	"strconv"
	"strings"
	//age "github.com/bearbin/go-age"
	"github.com/gorilla/mux"
	_ "io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"time"
	//"strconv"
	_ "strconv"
	//"strings"
	"testing"
)
type mockService struct{}
func getAge(year, month, day int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}
func TestCustomerById(t *testing.T) {
	testcases := []struct {
		id       string
		response []byte
		code     int
	}{
		{"30", []byte(`{"id":30,"name":"Bruce Wayne","dob":"12/12/1999","address":{"id":2,"city":"Bangalore","state":"karnataka","streetName":"Wayne Mansion","custId":30}}`), http.StatusOK},
		{"abc", []byte(`invalid id`), http.StatusBadRequest},
	}
	for i, v := range testcases {
		req := httptest.NewRequest(http.MethodGet, "/customer/", nil)
		req = mux.SetURLVars(req, map[string]string{"id": testcases[i].id})
		w := httptest.NewRecorder()
		a := New(mockService{})
		a.GetById(w, req)
		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(v.response)) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), string(v.response))
		}
		if w.Code != testcases[i].code {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Code, testcases[i].code)
		}
	}
}
func (c mockService) GetById(id int) (entities.Customer,error){
	if id == 0 {
		return entities.Customer{},errors.New("No data to update")
	}else{

		return entities.Customer{ID:30, Name: "Bruce Wayne", DOB: "12/12/1999",Address: entities.Address{ID: 2,City: "Bangalore",State: "karnataka", StreetName: "Wayne Mansion",CustId: 30}}, nil
	}
}

func TestCustomerByName(t *testing.T) {
	testcases := []struct {
		name       string
		response []byte
		code     int
	}{
		{"Bruce%20Wayne", []byte(`{"id":30,"name":"Bruce Wayne","dob":"12/12/1999","address":{"id":2,"city":"Bangalore","state":"karnataka","streetName":"Wayne Mansion","custId":30}}`), http.StatusOK},
		{"", []byte("could not get customer information"), http.StatusOK},
	}
	for i, v := range testcases {
		req := httptest.NewRequest(http.MethodGet, "/customer?name="+testcases[i].name, nil)
		w := httptest.NewRecorder()
		a := New(mockService{})
		a.GetByName(w,req)
		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(v.response)) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), string(v.response))
		}
		if w.Code != testcases[i].code {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Code, testcases[i].code)
		}
	}
}
func (c mockService) GetByName(name string)(entities.Customer,error) {
if name==""{
	return entities.Customer{}, errors.New("Empty name")
}
		return entities.Customer{ID: 30, Name: "Bruce Wayne", DOB: "12/12/1999",Address: entities.Address{ID: 2,City: "Bangalore",State: "karnataka", StreetName: "Wayne Mansion",CustId: 30}}, nil

}

func TestAllCustomer(t *testing.T) {
	testcases := []struct {
		response []byte
		code     int
	}{
		{ []byte(`[{"id":30,"name":"Bruce Wayne","dob":"12/12/1999","address":{"id":2,"city":"Bangalore","state":"karnataka","streetName":"Wayne Mansion","custId":30}}]`), http.StatusOK},
	}

	for i, v := range testcases {
		req := httptest.NewRequest(http.MethodGet, "/customer", nil)
		w := httptest.NewRecorder()
		a := New(mockService{})
		a.GetAll(w,req)
		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(v.response)) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), string(v.response))
		}
		if w.Code != testcases[i].code {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Code, testcases[i].code)
		}
	}
}
func (c mockService) GetAll()([]entities.Customer,error) {
	return []entities.Customer{entities.Customer{ID: 30, Name: "Bruce Wayne", DOB: "12/12/1999",Address: entities.Address{ID: 2,City: "Bangalore",State: "karnataka", StreetName: "Wayne Mansion",CustId: 30}}}, nil
}
func TestCustomerCreate(t *testing.T) {
	testcases := []struct {
		request []byte
		response []byte
		code     int
	}{
		{[]byte(`{"id":30,"name":"Bruce Wayne","dob":"12/12/1999","address":{"id":2,"city":"Bangalore","state":"karnataka","streetName":"Wayne Mansion","custId":30}}`), []byte(`{"id":30,"name":"Bruce Wayne","dob":"12/12/1999","address":{"id":2,"city":"Bangalore","state":"karnataka","streetName":"Wayne Mansion","custId":30}}`), http.StatusCreated},
		{[]byte(`{"id":30,"name":"Bruce Wayne","dob":"12/12/2010","address":{"id":2,"city":"Bangalore","state":"karnataka","streetName":"Wayne Mansion","custId":30}}`), []byte("could not create customer"), http.StatusOK},
	}

	for i, v := range testcases {
		req := httptest.NewRequest(http.MethodPost, "/customer", bytes.NewBuffer(testcases[i].request))
		w := httptest.NewRecorder()
		a := New(mockService{})
		a.Create(w,req)
		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(v.response)) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), string(v.response))
		}
		if w.Code != testcases[i].code {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Code, testcases[i].code)
		}
	}
}
func (c mockService) Create(customer1 entities.Customer)(entities.Customer,error) {
	dob := customer1.DOB
	dob1 := strings.Split(dob, "/")
	y, _ := strconv.Atoi(dob1[2])
	m, _ := strconv.Atoi(dob1[1])
	d, _ := strconv.Atoi(dob1[0])
	getAge := getAge(y, m, d)
	if age.Age(getAge) >= 18 {
		return entities.Customer{ID: 30, Name: "Bruce Wayne", DOB: "12/12/1999",Address: entities.Address{ID: 2,City: "Bangalore",State: "karnataka", StreetName: "Wayne Mansion",CustId: 30}}, nil
	}else {
		return entities.Customer{},errors.New("cannot create customer")
	}
}

func TestCustomerEdit(t *testing.T) {
	testcases := []struct {
		id string
		request []byte
		response []byte
		code     int
	}{
		{"30",[]byte(`{"id":30,"name":"Bruce Wayne","dob":"12/12/1999","address":{"id":2,"city":"Bangalore","state":"karnataka","streetName":"Wayne Mansion","custId":30}}`), []byte(`{"id":30,"name":"Bruce Wayne","dob":"12/12/1999","address":{"id":2,"city":"Bangalore","state":"karnataka","streetName":"Wayne Mansion","custId":30}}`), http.StatusOK},
		{"30",[]byte(`{"id":30,"name":"Bruce Wayne","address":{"id":2,"city":"Bangalore","state":"karnataka","streetName":"Wayne Mansion","custId":30}}`), []byte(`{"id":30,"name":"Bruce Wayne","dob":"12/12/1999","address":{"id":2,"city":"Bangalore","state":"karnataka","streetName":"Wayne Mansion","custId":30}}`), http.StatusOK},
	}
	for i, v := range testcases {
		req := httptest.NewRequest(http.MethodPut, "/customer", bytes.NewBuffer(testcases[i].request))
		req = mux.SetURLVars(req, map[string]string{"id": testcases[i].id})
		w := httptest.NewRecorder()
		a := New(mockService{})
		a.Edit(w,req)
		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(v.response)) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), string(v.response))
		}
		if w.Code != testcases[i].code {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Code, testcases[i].code)
		}
	}
}
func (c mockService) Edit(id int, customer1 entities.Customer)(entities.Customer,error){
	if id == 0{
		return entities.Customer{},errors.New("No data to update")
	}else{
		return entities.Customer{ID: 30, Name: "Bruce Wayne", DOB: "12/12/1999",Address: entities.Address{ID: 2,City: "Bangalore",State: "karnataka", StreetName: "Wayne Mansion",CustId: 30}},nil
	}
}
func TestDeleteCustomerById(t *testing.T) {
	testcases := []struct {
		id       string
		response []byte
		code     int
	}{
		{"30", []byte(`{"id":30,"name":"Bruce Wayne","dob":"12/12/1999","address":{"id":2,"city":"Bangalore","state":"karnataka","streetName":"Wayne Mansion","custId":30}}`), http.StatusNoContent},
		//{"abc", []byte(`invalid id`), http.StatusBadRequest},
	}
	for i, v := range testcases {
		req := httptest.NewRequest(http.MethodDelete, "/customer/", nil)
		req = mux.SetURLVars(req, map[string]string{"id": testcases[i].id})
		w := httptest.NewRecorder()
		a := New(mockService{})
		a.DeleteById(w, req)
		if !reflect.DeepEqual(w.Body, bytes.NewBuffer(v.response)) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Body.String(), string(v.response))
		}
		if w.Code != testcases[i].code {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, w.Code, testcases[i].code)
		}
	}
}
func(c mockService)  DeleteById(id int)(entities.Customer,error){
	if id ==0{
		return entities.Customer{},errors.New("invalid id")
	}else {
		return entities.Customer{ID: 30, Name: "Bruce Wayne", DOB: "12/12/1999",Address: entities.Address{ID: 2,City: "Bangalore",State: "karnataka", StreetName: "Wayne Mansion",CustId: 30}},nil
	}
}
