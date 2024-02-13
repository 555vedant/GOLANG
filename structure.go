package main

import "fmt"

// Define a struct named 'Person'
type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   Address // Embedded struct
	Email     string  // New string field
}

// Define another struct named 'Address'
type Address struct {
	Street  string
	City    string
	ZipCode string
}

func main() {
	// Create an instance of the 'Person' struct
	// or
	// var person1 Person

	person1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
		Address: Address{
			Street:  "123 Main St",
			City:    "Anytown",
			ZipCode: "12345",
		},
		Email: "john.doe@example.com", // Initialize the new 'Email' field
	}

	// Access and print the fields of the 'Person' struct
	fmt.Println("First Name:", person1.FirstName)
	fmt.Println("Last Name:", person1.LastName)
	fmt.Println("Age:", person1.Age)

	// Access and print the fields of the embedded 'Address' struct
	fmt.Println("Street:", person1.Address.Street)
	fmt.Println("City:", person1.Address.City)
	fmt.Println("Zip Code:", person1.Address.ZipCode)

	// Access and print the new 'Email' field
	fmt.Println("Email:", person1.Email)
}
