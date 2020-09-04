package main

import (
	"encoding/json"
	"fmt"
)

//Write a program which prompts the user to first enter a name, and then enter an address.
//Your program should create a map and add the name and address to the map using the keys “name” and “address”,
//respectively.
//Your program should use Marshal() to create a JSON object from the map,
//and then your program should print the JSON object

func main() {
	var name, address string
	fmt.Println("Enter a name: ")
	fmt.Scan(&name)
	fmt.Println("Enter an address: ")
	fmt.Scan(&address)

	//inserts the address to a map
	addressbook := make(map[string]string, 1)
	addressbook["name"] = name
	addressbook["address"] = address

	//create the JSON object
	jobj, err := json.Marshal(addressbook)
	if err != nil {
		return
	}

	fmt.Println("The JSON object is: ", string(jobj))

	//check the unmarshalled version
	/*add:=make(map[string]string)
	json.Unmarshal(jobj,&add)
	fmt.Println("After unmarshalling: the address of", name, " is ", add[name])*/
}
