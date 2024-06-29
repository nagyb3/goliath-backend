package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", helloWorldHandler)

	fmt.Println("Listening on port 8080")

	err := http.ListenAndServe(":8080", nil)
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
