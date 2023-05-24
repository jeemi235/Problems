package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/tealeg/xlsx"
)

// type Person struct {
// 	email      string `json:"email"`
// 	first_name string `json:"first_name"`
// 	gender     string `json:"gender"`
// 	id         string `json:"id"`
// 	ip_address string `json:"ip_address"`
// 	last_name  string `json:"last_name"`
// }

func main() {
	//router := http.NewServeMux()

	http.HandleFunc("/convert", convert)
	http.HandleFunc("/alldata", alldata)
	http.HandleFunc("/add", add)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/update", update)
	http.HandleFunc("/single", single)

	http.ListenAndServe(":1234", nil)
}

func convert(w http.ResponseWriter, r *http.Request) {

	// Open the XLSX file
	xlFile, err := xlsx.OpenFile("USER_MOCK_DATA.xlsx")
	if err != nil {
		panic(err)
	}
	sheet := xlFile.Sheets[0]

	// Convert the sheet contents into a slice of maps
	rows := make([]map[string]string, 0, len(sheet.Rows)-1)
	headers := make([]string, 0)
	for i, row := range sheet.Rows {
		if i == 0 {
			// Save the headers for later use
			for _, cell := range row.Cells {
				headers = append(headers, cell.String())
			}
		} else {
			// Create a new map for the row data
			rowData := make(map[string]string)
			for j, cell := range row.Cells {
				// Use the headers as keys for the map
				if j < len(headers) {
					rowData[headers[j]] = cell.String()
				}
				//fmt.Println(rowData)

			}
			rows = append(rows, rowData)
		}
	}
	// Encode the data into a JSON file
	jsonData, err := json.MarshalIndent(rows, "", "  ")
	if err != nil {
		panic(err)
	}

	// Write the JSON data to a file
	outputFile, err := os.Create("output.json")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	_, err = outputFile.Write(jsonData)
	if err != nil {
		panic(err)
	}

	fmt.Println("Conversion completed successfully!")
}

func alldata(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("output.json")
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(content))
	w.Write([]byte(string(content)))
}

func add(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"email":      "xyz@gmail.com",
		"first_name": "jonny",
		"gender":     "male",
		"id":         "s4ss5558s4d1x",
		"ip_address": "27.58.47",
		"lastname":   "jordan",
	}
	bytes, _ := json.MarshalIndent(data, "", " ")
	fmt.Println(string(bytes))
	outputFile, _ := os.Open("output.json")
	defer outputFile.Close()
	err := ioutil.WriteFile("output.json", bytes, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(bytes))
}

func delete(w http.ResponseWriter, r *http.Request) {

}

func update(w http.ResponseWriter, r *http.Request) {

}

func single(w http.ResponseWriter, r *http.Request) {
	//queryParams := r.URL.Query().Get("id")
}
