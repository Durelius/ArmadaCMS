package Controllers

import (
	"ArmadaCMS/db"
	"encoding/json"
	"net/http"
)

func ExampleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var success struct {
		IsSuccess    bool   `json:"testNumber"`
		ReturnString string `json:"testString"`
	}
	success.IsSuccess = true
	success.ReturnString = db.TestQuery()
	json.NewEncoder(w).Encode(success)
}
func ExamplePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var body struct {
		TestNumber int    `json:"testNumber"`
		TestString string `json:"testString"`
	}
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var success struct {
		IsSuccess    bool
		ReturnString string
	}
	success.IsSuccess = true
	success.ReturnString = "It worked!"
	json.NewEncoder(w).Encode(success)
}
