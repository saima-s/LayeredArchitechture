package store

import "C"
import (
	customer "Clean-Architecture/entities"
	error1 "Clean-Architecture/errorPackage"
	"database/sql"
	"log"
)

type CustomerStorer struct {
	db *sql.DB
}

func (c CustomerStorer) CloseDB() {
	c.db.Close()
}
func New() CustomerStorer {
	var db, err = sql.Open("mysql", "root:saima@123Sult@/CustomerDB")
	if err != nil {
		log.Fatal(error1.DbError)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(error1.DbError)
	}
	return CustomerStorer{db: db}
}
func (c CustomerStorer) GetByName(name string) (customer.Customer, error) {
	query := "SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId;"
	var data []interface{}
	query = "SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId where Customers.Name = ?;"
	data = append(data, name)
	rows, err := c.db.Query(query, data...)
	if err != nil {
		return customer.Customer{}, error1.NoDataError
	}
	defer rows.Close()
	var cust customer.Customer
	for rows.Next() {
		if err := rows.Scan(&cust.ID, &cust.Name, &cust.DOB, &cust.Address.ID, &cust.Address.City, &cust.Address.State, &cust.Address.StreetName, &cust.Address.CustId); err != nil {
			return customer.Customer{}, err
		}
	}
	return cust, nil
}
func (c CustomerStorer) GetAll() ([]customer.Customer, error) {
	query := "SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId;"
	rows, err := c.db.Query(query)
	if err != nil {
		return []customer.Customer{}, error1.NoDataError
	}
	defer rows.Close()
	var result []customer.Customer
	for rows.Next() {
		var cust customer.Customer
		if err := rows.Scan(&cust.ID, &cust.Name, &cust.DOB, &cust.Address.ID, &cust.Address.City, &cust.Address.State, &cust.Address.StreetName, &cust.Address.CustId); err != nil {
			return []customer.Customer{}, err
		}
		result = append(result, cust)
	}
	return result, nil
}
func (c CustomerStorer) GetById(id int) (customer.Customer, error) {
	var ids []interface{}
	ids = append(ids, id)
	query := `SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId where Customers.ID = ?; `
	rows, err := c.db.Query(query, ids...)
	if err != nil {
		return customer.Customer{}, error1.NoDataError
	}
	defer rows.Close()
	var cust customer.Customer
	for rows.Next() {
		if err := rows.Scan(&cust.ID, &cust.Name, &cust.DOB, &cust.Address.ID, &cust.Address.City, &cust.Address.State, &cust.Address.StreetName, &cust.Address.CustId); err != nil {
			return customer.Customer{}, err
		}
	}
	return cust, nil
}
func (c CustomerStorer) Create(customer1 customer.Customer) (customer.Customer, error) {
	var customers []interface{}
	customers = append(customers, customer1.Name)
	customers = append(customers, customer1.DOB)
	query := `INSERT INTO Customers(name, DOB) VALUES(?,?);`
	rows, err := c.db.Exec(query, customers...)
	if err != nil {
		return customer.Customer{}, err
	}
	id, _ := rows.LastInsertId()
	var addr []interface{}
	addr = append(addr, customer1.Address.City)
	addr = append(addr, customer1.Address.State)
	addr = append(addr, customer1.Address.StreetName)
	addr = append(addr, id)
	query1 := `INSERT INTO Address(City,State,StreetName,CustId) VALUES(?,?,?,?)`
	row, err1 := c.db.Exec(query1, addr...)
	if err1 != nil {
		return customer.Customer{}, err1
	}
	idAddr, _ := row.LastInsertId()
	customer1.ID = int(id)
	customer1.Address.ID = int(idAddr)
	customer1.Address.CustId = int(id)
	return customer1, nil
}
func (c CustomerStorer) Edit(id int, customer2 customer.Customer) (customer.Customer, error) {
	query1 := `SELECT * from Customers where ID =?`
	var id1 []interface{}
	id1 = append(id1, id)
	_, err := c.db.Query(query1, id1...)
	if err != nil {
		return customer.Customer{}, err
	}
	if customer2.Name != "" {
		_, err := c.db.Exec("update Customers set Name=? where ID=?", customer2.Name, id)
		if err != nil {
			return customer.Customer{}, err
		}
	}
	var custId []interface{}
	custId = append(custId, id)
	q := `SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId  where Address.CustID =?`
	r, _ := c.db.Query(q, custId...)
	var cu customer.Customer
	for r.Next() {
		e := r.Scan(&cu.ID, &cu.Name, &cu.DOB, &cu.Address.ID, &cu.Address.City, &cu.Address.State, &cu.Address.StreetName, &cu.Address.CustId)
		if e != nil {
			return customer.Customer{}, err
		}
	}
	var data []interface{}
	query := "update Address set "
	if customer2.Address.City != "" {
		query += "City = ? ,"
		data = append(data, customer2.Address.City)
	}
	if customer2.Address.State != "" {
		query += "State = ? ,"
		data = append(data, customer2.Address.State)
	}
	if customer2.Address.StreetName != "" {
		query += "StreetName = ? ,"
		data = append(data, customer2.Address.StreetName)
	}
	query = query[:len(query)-1]
	query += "where CustId = ? and ID = ?"
	data = append(data, id)
	data = append(data, cu.Address.ID)
	_, err = c.db.Exec(query, data...)
	if err != nil {
		return customer.Customer{}, err
	}
	var custIds []interface{}
	custIds = append(custIds, id)
	query2 := `SELECT * FROM Customers INNER JOIN Address on Customers.ID = Address.CustId  where Address.CustID =?`
	ro, _ := c.db.Query(query2, custIds...)
	defer ro.Close()
	var cu1 customer.Customer
	for ro.Next() {
		e := ro.Scan(&cu1.ID, &cu1.Name, &cu1.DOB, &cu1.Address.ID, &cu1.Address.City, &cu1.Address.State, &cu1.Address.StreetName, &cu1.Address.CustId)
		if e != nil {
			return customer.Customer{}, e
		}
	}
	return cu1, nil
}
func (c CustomerStorer) DeleteById(id int) (customer.Customer, error) {
	var ids []interface{}
	ids = append(ids, id)
	query := `SELECT * FROM Customers INNER JOIN Address ON Customers.ID = Address.CustId where Customers.ID = ?; `
	rows, err := c.db.Query(query, ids...)
	if err != nil {
		return customer.Customer{}, err
	}
	var cust customer.Customer
	for rows.Next() {
		if err := rows.Scan(&cust.ID, &cust.Name, &cust.DOB, &cust.Address.ID, &cust.Address.City, &cust.Address.State, &cust.Address.StreetName, &cust.Address.CustId); err != nil {
			return customer.Customer{}, err
		}
	}
	query = `DELETE  FROM Customers where ID =?; `
	_, err1 := c.db.Exec(query, ids...)
	if err1 != nil {
		return customer.Customer{}, err1
	}
	defer rows.Close()
	return cust, nil
}
