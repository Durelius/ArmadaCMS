package main

import (
	"ArmadaCMS/Controllers"
	utilities "ArmadaCMS/Utilities"
	"ArmadaCMS/db"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	utilities.CheckEnvVariables("JWT_SECRET_ARMADA_CMS", "ENC_KEY_ARMADA_CMS", "DB_HOST", "DB_PORT", "API_PORT", "DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_SSLMODE")
	db.ConnectDB()
	defer db.DB.Close()

	db.CreateTables()

	port := os.Getenv("API_PORT")
	wrappedMux := CreateMuxClient()
	colonPort := fmt.Sprintf(":%s", port)
	fmt.Println("Server running on http://localhost" + colonPort)
	if err := http.ListenAndServe(colonPort, wrappedMux); err != nil {
		log.Fatal("Server error:", err)
	}

}
func CreateMuxClient() http.Handler {
	mux := mux.NewRouter()
	mux = CreateControllers(mux)
	wrappedMux := HandleCORS(mux)
	return wrappedMux
}
func CreateControllers(mux *mux.Router) *mux.Router {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	mux.HandleFunc("/api/cms/newuser", Controllers.AddUser)
	mux.HandleFunc("/api/cms/tokenlogin", Controllers.TokenLogin)
	mux.HandleFunc("/api/cms/refreshAccessToken", Controllers.RefreshAccessToken)
	mux.HandleFunc("/api/cms/login", Controllers.Login)

	mux.HandleFunc("/api/cms/insertblogpost", Controllers.InsertBlogpost)
	mux.HandleFunc("/api/cms/getallblogposts", Controllers.GetAllBlogposts)

	return mux
}
func HandleCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-RefreshAuthorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
