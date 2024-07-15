package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       int    `json:"age"`
}

func main() {
	// test: curl --json "{\"firstname\":\"Yuuki\",\"lastname\":\"NguyenNgocMinhQuocccccc\",\"age\":24}" http://localhost:8088/decode
	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var user User
		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "%s %s is %d years old.", user.FirstName, user.LastName, user.Age)
	})

	// curl -s http://localhost:8088/encode
	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		yuuki := User{
			FirstName: "Yuuki",
			LastName:  "Kaitou",
			Age:       00,
		}

		json.NewEncoder(w).Encode(yuuki)
	})

	log.Fatal(http.ListenAndServe(":8088", nil))
}
