package main

import (
	"ArmadaCMS/Controllers"
	utilities "ArmadaCMS/Utilities"
	"ArmadaCMS/db"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "ArmadaCMS/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Your API
// @version 1.0
// @description CMS API
// @BasePath /api/cms
func main() {
	utilities.CheckEnvVariables("JWT_SECRET_ARMADA_CMS", "ENC_KEY_ARMADA_CMS", "DB_HOST", "DB_PORT", "API_PORT", "DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_SSLMODE")
	db.ConnectDB()
	// db.DB.AutoMigrate(&Structure.User{}, &Structure.Blogpost{}, &Structure.BlogpostTag{}, &Structure.Tokens{}, &Structure.RefreshTokenDB{})

	// defer db.DB.

	// db.CreateTables()

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

	swaggerPath := "docs/swagger.json"
	mux.HandleFunc("/api/cms/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Serving swagger.json")

		w.Header().Set("Content-Type", "application/json")
		http.ServeFile(w, r, swaggerPath)
	})
	mux.PathPrefix("/api/cms/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/api/cms/swagger.json"), // Must match your route
		httpSwagger.PersistAuthorization(true),
	))

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
