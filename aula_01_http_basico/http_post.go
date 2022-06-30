package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {

	post := Post{
		UserId: 1,
		Title:  "Title",
		Body:   "Body",
	}

	jsonBytes, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(resp.Status)
	println(string(bytes))

}
