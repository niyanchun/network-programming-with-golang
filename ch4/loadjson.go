package main

import (
	"os"
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  Name
	Email []Email
}

type Name struct {
	Family   string
	Personal string
}

type Email struct {
	Kind    string
	Address string
}

func (p Person) String() string {
	s := p.Name.Personal + " " + p.Name.Family

	for _, v := range p.Email {
		s += "\n" + v.Kind + ": " + v.Address
	}

	return s
}

func main() {
	var person Person
	loadJson("person.json", &person)

	fmt.Println("Person", person.String())
}

func loadJson(fileName string, key interface{}) {
	inFile, err := os.Open(fileName)
	checkErr(err)
	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(key)
	checkErr(err)
	inFile.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}