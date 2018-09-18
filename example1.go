package main

import (
	"fmt"
	"strconv"
)

//Customer structure
type Customer struct {
	Name    string
	Address string
}

//Employee structure
type Employee struct {
	Name   string
	Salary int
}

type People interface {
	GetInfo(chan string) chan string
}

func (c Customer) GetInfo(ch chan string) chan string {
	ch <- c.Name + ", " + c.Address //Send customer name and address data to ch
	return ch
}

func (e Employee) GetInfo(ch chan string) chan string {
	ch <- e.Name + ", " + strconv.Itoa(e.Salary) //Send employee name and salary data to ch
	return ch
}

func GetPeopleInfo(p People) {
	ch := make(chan string)
	go p.GetInfo(ch)
	result := <-ch //Receive value from ch
	fmt.Println(result)
}

func main() {
	//Initialize customer object
	customer := Customer{Name: "Customer 1", Address: "Address 1"}

	//Initialize employee object
	employee := Employee{Name: "Employee 1", Salary: 3000}

	GetPeopleInfo(customer)
	GetPeopleInfo(employee)

	mp := make(map[string][]Customer)

	customers := []Customer{
		customer,
		Customer{Name: "Customer 2", Address: "Address 2"}}

	mp["Indonesia"] = customers

	fmt.Println(mp["Indonesia"])

	for country, names := range mp {
		for _, name := range names {
			fmt.Println("\nName: " + name.Name + "\nAddress: " + name.Address + "\nCountry: " + country + "\n")
		}
	}

}