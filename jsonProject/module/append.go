// package main

// import (
// 	"encoding/json"
// 	//"error"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// )

// type Person struct {
// 	email      string `json:"email"`
// 	first_name string `json:"first_name"`
// 	gender     string `json:"gender"`
// 	id         string `json:"id"`
// 	ip_address string `json:"ip_address"`
// 	last_name  string `json:"last_name"`
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":8081", nil)

// }
// func handler(w http.ResponseWriter, r *http.Request) {
// 	// content, err := ioutil.ReadFile("output.json")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Println(string(content))
// 	// w.Write([]byte(string(content)))
// 	file, err := ioutil.ReadFile("output.json")
// 	if err != nil {
// 		fmt.Println("Error reading file:", err)
// 		return
// 	}

// 	var people []Person
// 	err = json.Unmarshal(file, &people)
// 	if err != nil {
// 		fmt.Println("Error unmarshalling JSON:", err)
// 		return
// 	}
// 	fmt.Println(people)

// 	newPerson := Person{email: "jonny@gmail.com", first_name: "jonny", gender: "male", id: "1001", ip_address: "10.25.368", last_name: "jordan"}
// 	people = append(people, newPerson)

// 	updatedJSON, err := json.Marshal(people)
// 	if err != nil {
// 		fmt.Println("Error marshalling JSON:", err)
// 		return
// 	}

//		// Write the JSON data to the file
//		err = ioutil.WriteFile("output.json", updatedJSON, os.ModePerm)
//		if err != nil {
//			fmt.Println("Error writing file:", err)
//			return
//		}
//		w.Write([]byte(string(updatedJSON)))
//	}
package main
