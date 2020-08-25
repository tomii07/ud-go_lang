package main

import (
	"fmt"
	"net/http"
)

type testStruct struct {
	Value1 string
	Value2 int
}

func main() {

	// First Test
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test works")
	})

	// Getting the Query Value from: http://localhost:9001?TQ=testQuery
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Query value: " + r.URL.Query().Get("TQ"))
	})

	// Getting the Form Value from: http://localhost:9001
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Form value: " + r.FormValue("TF"))
	})

	// Getting the Header Value from: http://localhost:9001
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			
		fmt.Fprintf(w, "Header value: " + r.Header.Get("TH"))
		
		// Getting the Method
		fmt.Fprintf(w, "Header value: " + r.Method)

		// Example for checking the Method
		if r.Method == "POST" {
					 
			}
	})

	// Setting and Getting the Header Value from: http://localhost:9001 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		
		w.Header().Set("TestHeader", "hval")
		//fmt.Fprintf(w, "Test")
		
		var ts testStruct
		ts.Value1 = "Test value 1"
		ts.Value2 = 4
		
		fmt.Fprintf(w, prepResponse(ts))
	})

	// http://localhost:9001/testauth
	http.HandleFunc("/testauth", func(w http.ResponseWriter, r *http.Request) {
		if authorize(r.Header.Get("Auth")) == "" {
			w.WriteHeader(401)
			return
		}
		fmt.Fprintf(w, prepResponse("Test Passed"))
	})

	// http://localhost:9001/testmethod
	http.HandleFunc("/testmethod", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(405)
			return
		}
		fmt.Fprintf(w, prepResponse("Test Passed"))
	})

	http.ListenAndServe(":9001", nil)
}

func prepResponse(obj interface{}) string {
	resp, _ := json.Marshal(obj)
	return string(resp)
}

func authorize(token string) string {
	if token == "test" {
		return "userid"
	}
	return ""
}