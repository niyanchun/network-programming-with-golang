package main

import (
	"os"
	"encoding/json"
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

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind: "home", Address: "jan@newmarch.com"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	saveJson("person.json", person)
}

func saveJson(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkErr(err)

	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkErr(err)
	outFile.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
