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

type BrokerageQuery struct {
	RequestType string
	Name        string
}

func coolPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the coolPage!")
	fmt.Println("Endpoint Hit: coolPage")
}

func stockQuery(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var brokerageQuery BrokerageQuery
	err := decoder.Decode(&brokerageQuery)
	if err != nil {
		panic(err)
		fmt.Println("Stock error1")
	}
	fmt.Println(brokerageQuery.RequestType) //brokerageQuery.Name
	// fmt.Println("Endpoint Hit: coolPage")

	// m := map[string]string{}
	// 	json.Unmarshal([]byte(body), &m)
	// 	fmt.Println(m)
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

	http.HandleFunc("/stockQuery", stockQuery)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	// post()
	// get()
	handleRequests()
}
