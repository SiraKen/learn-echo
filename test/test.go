package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Users struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func main() {
	
	bytes, err := ioutil.ReadFile("./hello.json")

	if err != nil {
		log.Fatal(err)
	}

	var users []Users

	if err := json.Unmarshal(bytes, &users); err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		fmt.Println(user.Id, user.Name)
	}

	fmt.Println("------------------")

	makeJson()
	
}

func makeJson() {
	users := []Users{
		{Id: 1, Name: "hoge"},
		{Id: 2, Name: "fuga"},
	}

	bytes, err := json.Marshal(users)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
}
