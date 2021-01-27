package service

import customer "Clean-Architecture/entities"

type Customers interface {
	GetByName(name string) (customer.Customer,error)
	GetAll()([]customer.Customer,error)
	GetById(id int) (customer.Customer,error)
	Create(c customer.Customer) (customer.Customer,error)
	Edit(id int, c customer.Customer) (customer.Customer,error)
	DeleteById(id int) (customer.Customer,error)
}