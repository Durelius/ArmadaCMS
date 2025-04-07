package Controllers

import (
	"ArmadaCMS/Flow"
	"ArmadaCMS/Structure"
	"ArmadaCMS/Utilities"
	"ArmadaCMS/db"
	"encoding/json"
	"net/http"
	"strings"
)

func RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	refreshAuthorizationHeader := r.Header.Get("X-RefreshAuthorization")
	accessAuthorizationHeader := r.Header.Get("Authorization")

	if refreshAuthorizationHeader == "" {
		http.Error(w, "Authorization header (r1)", http.StatusBadRequest)
		return
	}
	refreshParts := strings.Fields(refreshAuthorizationHeader)
	accessParts := strings.Fields(accessAuthorizationHeader)
	if len(refreshParts) != 2 || refreshParts[0] != "Bearer" {
		http.Error(w, "Authorization header (r2)", http.StatusBadRequest)
		return
	}
	userId := Utilities.GetUserIdFromAccessToken(accessParts[1])
	if userId == nil {
		http.Error(w, "Invalid token (r455)", http.StatusBadRequest)
		return
	}
	rToken := Flow.VerifyRefreshToken(refreshParts[1], int(*userId))
	if rToken == nil {
		http.Error(w, "Invalid token (r456)", http.StatusBadRequest)
		return
	}

	newAccessToken, _ := Utilities.GenerateAccessToken(*userId)

	tokens := Structure.Tokens{RefreshToken: *rToken.RefreshToken, AccessToken: newAccessToken}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokens)

}

func TokenLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	authorizationHeader := r.Header.Get("Authorization")

	if authorizationHeader == "" {
		http.Error(w, "Authorization header (1)", http.StatusBadRequest)
		return
	}
	parts := strings.Fields(authorizationHeader)
	if len(parts) != 2 || parts[0] != "Bearer" {
		http.Error(w, "Authorization header (2)", http.StatusBadRequest)
		return
	}

	mapClaims, err := Utilities.VerifyAccessToken(parts[1])
	if err != nil {
		if err.Error() == "token has invalid claims: token is expired" {
			http.Error(w, "Invalid token (321)", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Invalid token (456)", http.StatusBadRequest)
		return
	}
	userID, ok := (*mapClaims)["user_id"].(float64)
	if !ok {
		http.Error(w, "Invalid token (111)", http.StatusBadRequest)
		return
	}
	user, err := db.PopulateUserFromId(int(userID))
	if err != nil {
		http.Error(w, "Invalid token (000)", http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)

}
