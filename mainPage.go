package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func textHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Open the JSON file
	file, err := os.Open("/app/homePage.json")
	if err != nil {
		// Handle error, send a 500 status and log the error
		http.Error(w, "Could not open JSON file", http.StatusInternalServerError)
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	// Read the file contents
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		// Handle error, send a 500 status and log the error
		http.Error(w, "Could not read JSON file", http.StatusInternalServerError)
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Send the file content as the response
	w.Write(byteValue)

	// Optional: log to console for debugging
	fmt.Println("Request received at /api/text, JSON file served")
}
