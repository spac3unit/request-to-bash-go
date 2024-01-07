package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type Params struct {
	CPU string `json:"cpu"`
	RAM string `json:"ram"`
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/create/server", enableCORS(http.HandlerFunc(echoHandler)))

	fmt.Println("Server is listening on :8081 (HTTP)")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error starting HTTP server:", err)
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	var params Params
	params.CPU = r.FormValue("cpu")
	params.RAM = r.FormValue("ram")

	// Convert the struct to JSON
	jsonData, err := json.Marshal(params)
	if err != nil {
		http.Error(w, "Error creating JSON", http.StatusInternalServerError)
		return
	}

	// Execute the echo command in Bash
	cmd := exec.Command("bash", "-c", fmt.Sprintf("echo '%s'", jsonData))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error running echo command:", err)
	}

	// Write the JSON response directly to the client
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
