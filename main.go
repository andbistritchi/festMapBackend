package main

import (
	"context"
	"fmt"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/swag"
	_ "go.opentelemetry.io/otel/trace"
	"net/http"
	"os"
	_ "time"
)

func main() {
	// Initialize Jaeger tracer
	//initJaegerMain()

	// Serve static files (Angular app)
	angularAppMain()

	// Serve the text data
	readJsonMain()

	// Connect to SQLite database
	//connectDB()
}

func readJsonMain() {
	http.HandleFunc("/api/text", textHandler)

	port := ":8080"
	fmt.Printf("Server running on http://localhost%s\n", port)

	err := http.ListenAndServe(port, nil)

	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}

func angularAppMain() {
	fs := http.FileServer(http.Dir("./frontend1/dist"))
	http.Handle("/", fs)
}

func initJaegerMain() {
	tracerProvider, err := initJaegerTracer()
	if err != nil {
		fmt.Printf("Error initializing tracer: %v\n", err)
		return
	}
	defer func() {
		if err := tracerProvider.Shutdown(context.Background()); err != nil {
			fmt.Printf("Error shutting down tracer: %v\n", err)
		}
	}()
}
