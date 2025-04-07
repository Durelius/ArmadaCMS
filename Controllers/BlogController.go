package Controllers

import (
	"ArmadaCMS/Flow"
	"ArmadaCMS/Structure"
	"ArmadaCMS/db"
	"encoding/json"
	"log"
	"net/http"
)

// @Summary Insert blogpost
// @Description Insert new blogpost, requires authentication
// @Tags blog
// @Produce json
// @Param   blogPost body Structure.NewBlogpost true "Blog post content"
// @Param   Authorization header string true "Bearer token"
// @Param   x-RefreshAuthorization header string true "Bearer token"
// @Success 200 {object} boolean
// @Router /insertblogpost [post]
func InsertBlogpost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("err")
		http.Error(w, "Method not allowed (0) IBP", http.StatusMethodNotAllowed)
		return
	}
	userId, doRefresh := Flow.VerifyAccessTokenWebRequest(r)
	if doRefresh {
		http.Error(w, "Invalid token (321) IBP", http.StatusUnauthorized)
		return
	}
	if userId == nil {
		http.Error(w, "Invalid token (322) IBP", http.StatusBadRequest)
		return
	}

	var newBlogpost Structure.NewBlogpost
	err := json.NewDecoder(r.Body).Decode(&newBlogpost)
	if err != nil {
		http.Error(w, "Invalid request body (323) IBP", http.StatusBadRequest)
		return
	}
	db.InsertBlogpostDB(newBlogpost)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)

}

// @Summary Get all blog posts
// @Description Get list of all blog posts, no security here at all
// @Tags blog
// @Produce json
// @Success 200 {object} []Structure.Blogpost
// @Router /getallblogposts [get]
func GetAllBlogposts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("err")
		http.Error(w, "Method not allowed (0) IBP", http.StatusMethodNotAllowed)
		return
	}

	blogposts := db.GetAllBlogpostsDB()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blogposts)

}
