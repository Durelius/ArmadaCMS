package Controllers

import (
	"ArmadaCMS/Flow"
	"ArmadaCMS/Structure"
	"ArmadaCMS/db"
	"encoding/json"
	"log"
	"net/http"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user Structure.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	createdUser, err := db.InsertUser(user)
	if err != nil {
		http.Error(w, "AU1", http.StatusBadRequest)
		return
	}
	response, err := Flow.VerificationFlow(*createdUser)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var user Structure.User
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	user.Username = data["username"].(string)
	user.Password = data["password"].(string)

	w.Header().Set("Content-Type", "application/json")

	response, err := Flow.VerificationFlow(user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(response)

}
