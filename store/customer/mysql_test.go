package store

import (
	"Clean-Architecture/entities"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"testing"
)
//func initializeMySQL(t *testing.T) *sql.DB {
//	conf := driver.MySQLConfig{
//		Host:     os.Getenv("SQL_HOST"),
//		User:     os.Getenv("SQL_USER"),
//		Password: os.Getenv("SQL_PASSWORD"),
//		Port:     os.Getenv("SQL_PORT"),
//		Db:       os.Getenv("SQL_DB"),
//	}
//
//	var err error
//	db, err := driver.ConnectToMySQL(conf)
//	if err != nil {
//		t.Errorf("could not connect to sql, err:%v", err)
//	}
//
//	return db
//}

func TestDatastore(t *testing.T) {
	a := New()
	testCustomer_GetbyId(t, a)
	testCustomer_Create(t, a)
	testCustomer_GetbyName(t,a)
	testCustomer_Edit(t, a)
	testCustomer_DeleteById(t,a)
}
func testCustomer_Create(t *testing.T, db CustomerStorer) {
	testcases := []struct {
		req      entities.Customer
		response entities.Customer
	}{
		{entities.Customer{Name: "CustomerA1", DOB: "10/10/2000", Address: entities.Address{City:"Hyderabad", State: "Telangana", StreetName: "HSR"}}, entities.Customer{ID:49, Name: "CustomerA1", DOB: "10/10/2000", Address: entities.Address{ID:21, City:"Hyderabad", State: "Telangana", StreetName: "HSR",CustId: 49}}},
		//{entities.Customer{Name: "CustomerA2", DOB: "10/10/2000", Address: entities.Address{City:"Hyderabad", State: "Telangana", StreetName: "HSR"}}, entities.Customer{Name: "CustomerA2", DOB: "10/10/2010", Address: entities.Address{City:"Bangalore1", State: "Bangalore", StreetName: "HSR"}}},
		//{entities.Customer{Name: "CustomerA3", DOB: "10/10/2000", Address: entities.Address{City:"Hyderabad", State: "Telangana", StreetName: "HSR"}}, entities.Customer{Name: "CustomerA3", DOB: "10/10/2010", Address: entities.Address{City:"Bangalore2", State: "Bangalore", StreetName: "HSR"}}},

	}
	for i, v := range testcases {
		resp, _ := db.Create(v.req)

		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.response)
		}
	}
}
func testCustomer_GetbyName(t *testing.T, db CustomerStorer) {
	testcases := []struct {
		name   string
		resp entities.Customer
	}{
		{"Bruce Wayne", entities.Customer{ID: 30, Name: "Bruce Wayne", DOB: "12/12/1999", Address: entities.Address{2, "Bangalore", "Karnataka", "Wayne Mansion", 30}}},
		{"CustomerX1", entities.Customer{ID: 33, Name: "CustomerX1", DOB: "10/10/1997", Address: entities.Address{5, "Bangalore101", "Karnataka4", "HSR Layout",33}}},
		{"CustomerY1", entities.Customer{}},
	}
	for i, v := range testcases {
		resp, _ := db.GetByName(v.name)

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.resp)
		}
	}
}
func testCustomer_GetbyId(t *testing.T, db CustomerStorer) {
	testcases := []struct {
		id   int
		resp entities.Customer
	}{
		{30, entities.Customer{ID: 30, Name: "Bruce Wayne", DOB: "12/12/1999", Address: entities.Address{2, "Bangalore", "Karnataka", "Wayne Mansion", 30}}},
		{70, entities.Customer{}},
	}
	for i, v := range testcases {
		resp, _ := db.GetById(v.id)

		if !reflect.DeepEqual(resp, v.resp) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.resp)
		}
	}
}
func testCustomer_Edit(t *testing.T, db CustomerStorer) {
	testcases := []struct {
		id int
		req      entities.Customer
		response entities.Customer
	}{
		{49,entities.Customer{Name: "CustomerA0", Address: entities.Address{City:"Hyderabad0", State: "Telangana0", StreetName: "HSR0"}}, entities.Customer{ID:49,Name: "CustomerA0", DOB: "10/10/2000", Address: entities.Address{ID: 21,City:"Hyderabad0", State: "Telangana0", StreetName: "HSR0",CustId: 49}}},
		//{39, entities.Customer{Name: "CustomerA2", Address: entities.Address{City:"Hyderabad2", State: "Telangana2", StreetName: "HSR2"}}, entities.Customer{ID:39,Name: "CustomerA2", DOB: "10/10/2000", Address: entities.Address{ID:11,City:"Hyderabad2", State: "Telangana2", StreetName: "HSR2",CustId: 39}}},
		//{40, entities.Customer{Name: "CustomerA3", Address: entities.Address{City:"Hyderabad3", State: "Telangana3", StreetName: "HSR3"}}, entities.Customer{ID:40, Name: "CustomerA3", DOB: "10/10/2000", Address: entities.Address{ID:12, City:"Hyderabad3", State: "Telangana3", StreetName: "HSR3",CustId: 40}}},

	}
	for i, v := range testcases {
		resp, _ := db.Edit(v.id,v.req)

		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.response)
		}
	}
}
func testCustomer_DeleteById(t *testing.T, db CustomerStorer) {
	testcases := []struct {
		id int
		response entities.Customer
	}{
		{50, entities.Customer{ID:50,Name: "CustomerA1", DOB: "10/10/2000", Address: entities.Address{ID: 22,City:"Hyderabad", State: "Telangana", StreetName: "HSR",CustId: 50}}},
		//{39, entities.Customer{Name: "CustomerA2", Address: entities.Address{City:"Hyderabad2", State: "Telangana2", StreetName: "HSR2"}}, entities.Customer{ID:39,Name: "CustomerA2", DOB: "10/10/2000", Address: entities.Address{ID:11,City:"Hyderabad2", State: "Telangana2", StreetName: "HSR2",CustId: 39}}},
		//{40, entities.Customer{Name: "CustomerA3", Address: entities.Address{City:"Hyderabad3", State: "Telangana3", StreetName: "HSR3"}}, entities.Customer{ID:40, Name: "CustomerA3", DOB: "10/10/2000", Address: entities.Address{ID:12, City:"Hyderabad3", State: "Telangana3", StreetName: "HSR3",CustId: 40}}},

	}
	for i, v := range testcases {
		resp, _ := db.DeleteById(v.id)
        fmt.Println(resp)
		if !reflect.DeepEqual(resp, v.response) {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i+1, resp, v.response)
		}
	}
}