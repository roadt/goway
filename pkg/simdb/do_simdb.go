package main

import db "github.com/sonyarouje/simdb"
import "fmt"

type Customer struct {
	CId string `json:"cid"`
  Name string `json:"name"`
	Address string `json:"address"`
	Contact Contact
}

type Contact struct {
	Phone string `json:"phone"`
	Email string  `json:"email"`
}

//ID any struct that needs to persist should implement this function defined 
//in Entity interface.
func (c Customer) ID() (jsonField string, value interface{}) {
	value=c.CId
	jsonField="cid"
	return
}
 
func main() {
	driver, err := db.New("data")
	if err != nil {
		panic(err)
	}

	customer :=Customer {
		CId: "CUST1",
		Name: "Frank",
    Address: "Nanning Street",
    Contact: Contact {
			Phone: "45533355",
			Email: "someone@gmail.com",
		},
	}

  //creates a new Customer file inside the directory passed as the parameter to New()
  //if the Customer file already exist then insert operation will add the customer data to the array
	err = driver.Insert(customer)
	if err != nil {
		panic(err)
	}

  
  //GET ALL Customer
  //opens the customer json file and filter all the customers with name sarouje.
  //AsEntity takes a pointer to Customer array and fills the result to it.
  //we can loop through the customers array and retireve the data.
	var customers  []Customer
	err = driver.Open(Customer{}).Where("name", "=", "Frank").Get().AsEntity(&customers)
	if err != nil {
		panic(err)
	}
	fmt.Println(customers)

  //GET ONE Customer
  //First() will return the first record from the results 
  //AsEntity takes a pointer to Customer variable (not an array pointer)
  var customerFirst Customer
  err=driver.Open(Customer{}).Where("cid","=","CUST1").First().AsEntity(&customerFirst)
  if err!=nil {
    panic(err)
  }
	fmt.Println(customerFirst)

  //Update function uses the ID() to get the Id field/value to find the record and update the data.
  customerFirst.Name="Sony Arouje"
  err=driver.Update(customerFirst)
  if err!=nil {
    panic(err)
  }
	fmt.Println(customerFirst)

  
  //Delete
  toDel:=Customer{
     CId:"CUST1",
  }
  err=driver.Delete(toDel)
}


