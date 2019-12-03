package main

// func performWebscrape() {
// Make request
// response, err := http.Get("https://www.devdungeon.com/archive")
// response, err := http.Get("https://stocktwits.com/symbol/AAPL")

// if err != nil {
// 	log.Fatal(err)
// }
// defer response.Body.Close()

// // Create output file
// outFile, err := os.Create("output.html")
// if err != nil {
// 	log.Fatal(err)
// }
// defer outFile.Close()

// // Copy data from HTTP response to file
// _, err = io.Copy(outFile, response.Body)
// if err != nil {
// 	log.Fatal(err)
// }

// outFile, err := os.Create("output.html")
// if err != nil {
// 	log.Fatal(err)
// }
// defer outFile.Close()

// // Copy data from HTTP response to file
// _, err = io.Copy(outFile, response.Body)
// if err != nil {
// 	log.Fatal(err)
// }

// b, err := ioutil.ReadFile("input.txt")
// if err != nil {
// 	panic(err)
// }

// b := "documentation awesome"
// // write the whole body at once
// err := ioutil.WriteFile("output.txt", b, 0644)
// if err != nil {
// 	panic(err)
// }

// f, err := os.Create("output1.txt")
// check(err)

// defer f.Close()

// n3, err := f.WriteString("writes\n")
// fmt.Printf("wrote %d bytes\n", n3)

// Copy data from the response to standard output
// n, err := io.Copy(os.Stdout, response.Body)
// if err != nil {
// 	log.Fatal(err)
// }

// log.Println("Number of bytes copied to STDOUT:", n)

// }

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }
