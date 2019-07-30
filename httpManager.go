package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type response1 struct {
}

// type response2 struct {
// 	Page   int      `json:"page"`
// 	Fruits []string `json:"fruits"`
// }

// type User struct {
// 	ID        int
// 	Age       int
// 	FirstName string
// 	LastName  string
// 	Email     string
// }

// func get() {
// 	// resp, err := http.Get("http://localhost:20000/")
// 	// resp, err := http.Get("https://api.tradeking.com/v1/")
// 	// resp, err := http.Get("http://api.open-notify.org/astros.json")

// 	// resp, err := http.Get("https://blog.golang.org/go-maps-in-action")

// 	// req.Header.Set("User-Agent", "spacecount-tutorial")

// 	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
// 	resp, err := http.Get("https://api.tradeking.com/v1/")

// 	if err != nil {
// 		// handle error
// 		fmt.Println("Error 1")
// 	}
// 	defer resp.Body.Close()
// 	body, err1 := ioutil.ReadAll(resp.Body)
// 	if err1 != nil {
// 		// handle error
// 		fmt.Println("Error 2")
// 	}
// 	// fmt.Println(body)

// 	m := map[string]string{}
// 	json.Unmarshal([]byte(body), &m)
// 	fmt.Println(m)

// 	// str := `{"page": 1, "fruits": ["apple", "peach"]}`[]byte(resp)
// 	// res := response1{}
// 	// json.Unmarshal([]byte(body), &res)
// 	// fmt.Println(res)
// 	// fmt.Println(res.ID)
// }

func get() {
	url := "http://localhost:20000/"

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	httpClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")
	req.Header.Set("Content-Type", "application/json")
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(body)

	m := map[string]string{}
	json.Unmarshal([]byte(body), &m)
	fmt.Println(m["ID"])
}
func post(url string, json string) {
	// resp, err := http.Post("http://example.com/upload")
	// defer resp.Body.Close()
	// body, err := ioutil.ReadAll(resp.Body)

	// m := map[string]string{}
	// json.Unmarshal([]byte(body), &m)
	// fmt.Println(m)

	// http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// url := "https://postman-echo.com/post"

	// httpClient := http.Client{
	// 	Timeout: time.Second * 2, // Maximum of 2 secs
	// }

	// req, err := http.NewRequest(http.MethodPost, url, nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// req.Header.Set("User-Agent", "spacecount-tutorial")
	// req.Header.Set("Content-Type", "application/json")

	// fmt.Println("URL:>", url)

	var jsonStr = []byte(json)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
