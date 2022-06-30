package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UsersResponse struct {
	Data []struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	} `json:"data"`
}

type User struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
}

func main() {

	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://reqres.in/api/users", nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("Content-type", "application/json")
	request.SetBasicAuth("TUX", "HURB123")

	cookie := &http.Cookie{}

	cookie.Name = "SESSION"
	cookie.Value = "asndaosdnasiondiasndiasnd"

	request.AddCookie(cookie)

	enviarRequest(client)

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	var dados UsersResponse

	err = json.NewDecoder(response.Body).Decode(&dados)
	if err != nil {
		panic(err)
	}

	var usersNovo []map[string]interface{}
	for _, user := range dados.Data {

		userN := make(map[string]interface{})

		userN["ID"] = user.ID
		userN["first_name"] = user.FirstName

		usersNovo = append(usersNovo, userN)
	}

	bytesUsers, err := json.Marshal(usersNovo)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytesUsers))

}

func enviarRequest(client *http.Client) {

}
