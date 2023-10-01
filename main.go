package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

const endpoint1 string = "https://jsonplaceholder.typicode.com/todos/1"

func get() {

	fmt.Println("Performing HTTP GET")

	response, err := http.Get(endpoint1)
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	//Converte o response body em string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	//Converte o response body para Todo Struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("API Response as struct %+v\n", todoStruct)
}

func post() {

	fmt.Println("Performing HTTP POST")

	todo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	jsonReq, err := json.Marshal(todo)
	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer response.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(response.Body)

	// Converte response body em string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Converte response body em Todo Struct
	var todoStruct Todo
	json.Unmarshal(bodyBytes, &todoStruct)
	fmt.Printf("%+v\n", todoStruct)

}

func main() {
	get()
	post()
}
