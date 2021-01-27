package service

import (
	customer "Clean-Architecture/entities"
	"Clean-Architecture/store"
	"errors"
	age "github.com/bearbin/go-age"
	"reflect"
	"strconv"
	"strings"
	"time"
)
type CustomerService struct {
	store store.Customers
}

func New(customer store.Customers) CustomerService {
	return CustomerService{store: customer}
}
func getAge(year, month, day int) time.Time {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dob
}
func (c CustomerService)GetByName(name string) (customer.Customer,error){
	if name ==""{
		return customer.Customer{}, errors.New("Empty name")
	}else {

		resp, err := c.store.GetByName(name)
		if err != nil {
			return customer.Customer{}, err
		}
		return resp, nil
	}

}
func (c CustomerService)GetAll() ([]customer.Customer,error){
		resp ,err := c.store.GetAll()
		if err!=nil{
			return resp, err
		}
		return resp,nil
}
func (c CustomerService)GetById(id int) (customer.Customer,error){
    if id == 0 {
		return customer.Customer{},errors.New("No data to update")
	}else{
		resp ,err := c.store.GetById(id)
		if err!=nil{
			return resp, err
		}
		return resp,nil
	}
}
func (c CustomerService)Create(customer1 customer.Customer) (customer.Customer, error){
	dob := customer1.DOB
	dob1 := strings.Split(dob, "/")
	y, _ := strconv.Atoi(dob1[2])
	m, _ := strconv.Atoi(dob1[1])
	d, _ := strconv.Atoi(dob1[0])
	getAge := getAge(y, m, d)
	if age.Age(getAge) >= 18 {
		resp ,err:= c.store.Create(customer1)
		if err!=nil{
			return resp, err
		}
		return resp,nil
	}else{
		return customer.Customer{},errors.New("cannot create customer")
	}
}
func (c CustomerService)Edit(id int ,customer1 customer.Customer) (customer.Customer, error){
	if id == 0 {
		return customer.Customer{},errors.New("No data to update")
	}else{
		resp ,err := c.store.Edit(id,customer1)
		if err!=nil{
			return resp, err
		}
		if reflect.ValueOf(resp) == reflect.ValueOf(customer.Customer{}){
			return resp, errors.New("No data to update")
		}
		return resp,nil
	}
}
func (c CustomerService)DeleteById(id int) (customer.Customer,error){
	if id == 0 {
		return customer.Customer{},errors.New("No data to delete")
	}else{
		resp ,err := c.store.DeleteById(id)
		if err!=nil{
			return resp, err
		}
		if reflect.ValueOf(resp) == reflect.ValueOf(customer.Customer{}){
			return resp, errors.New("No data to delete")
		}
		return resp,nil
	}
}
