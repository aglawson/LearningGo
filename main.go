package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Age  int    `json:"age"`
}

type Response struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

var res = Response{Success: true, Data: "success"}

var users = []User{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Doe"},
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		queryParams := r.URL.Query()
		name := queryParams.Get("name")

		i := 0
		for i < len(users) {
			if name == users[i].Name {
				res.Success = true
				res.Data = name + " " + strconv.Itoa(users[i].ID)
				json.NewEncoder(w).Encode(res)
				return
			}
			i += 1
		}
		res.Success = false
		res.Data = "Not Found"
		json.NewEncoder(w).Encode(res)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		res.Success = false
		res.Data = err.Error()

		panic(res)
	}
}
