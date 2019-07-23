package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type JSONObject struct {
	ID        int
	Age       int
	FirstName string
	// LastName  string
	// Email     string
}

func coolPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the coolPage!")
	fmt.Println("Endpoint Hit: coolPage")
}
func test(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var jsonObject JSONObject
	err := decoder.Decode(&jsonObject)
	if err != nil {
		panic(err)
		fmt.Println("Test error1")
	}
	log.Println(jsonObject.ID, jsonObject.Age, jsonObject.FirstName)
}

// func create(w http.ResponseWriter, r *http.Request) {
// 	create(11)
// }

// func read(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: delete")
// 	deleteEntry(11)
// 	fmt.Println("Endpoint Hit end: delete")
// }

// func update(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: delete")
// 	deleteEntry(11)
// 	fmt.Println("Endpoint Hit end: delete")
// }

// func delete(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: delete")
// 	deleteEntry(11)
// 	fmt.Println("Endpoint Hit end: delete")
// }

func handleRequests() {
	// http.HandleFunc("/create", create)
	// http.HandleFunc("/read", read)
	// http.HandleFunc("/update", update)
	// http.HandleFunc("/delete", delete)

	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
