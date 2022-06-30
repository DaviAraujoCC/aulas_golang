package main

import (
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

	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		panic(err)
	}

	println(response.StatusCode)

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var posts []Post

	err = json.Unmarshal(bytes, &posts)
	if err != nil {
		panic(err)
	}

	for _, post := range posts {
		println(post.Title)
	}

}
