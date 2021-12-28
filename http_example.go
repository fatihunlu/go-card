package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Name   string `json:"name"`
	Height string `json:"height"`
}

func getPersonInfo() {
	resp, _ := http.Get("https://swapi.dev/api/people/1")

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body) // response body is []byte

	var result Person
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
		fmt.Println("err: ", err)
	}

	fmt.Println(result.Name)
}
