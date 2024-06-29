package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Listening on port " + port)

	http.HandleFunc("/hello-world", helloWorldHandler)

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := map[string]interface{}{
			"hello": "world",
		}

		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
