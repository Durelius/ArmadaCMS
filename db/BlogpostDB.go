package db

import (
	"ArmadaCMS/Structure"
	"log"
)

func InsertBlogpostDB(blogpost Structure.NewBlogpost) {
	const insertBlogPostSQL string = `
	INSERT INTO blogpost(text, title, author) VALUES($1,$2,$3)
	  ;`
	res, err := DB.Exec(insertBlogPostSQL, blogpost.Text, blogpost.Title, blogpost.Author)
	if err != nil {
		log.Fatal(err)
	}
	const insertBlogPostTagsSQL string = `
	INSERT INTO blogpost_tags(blogpost_id, tag) VALUES($1, $2)
	  ;`
	blogpostId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	for _, tag := range blogpost.Tags {
		_, err = DB.Exec(insertBlogPostTagsSQL, blogpostId, tag)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func GetAllBlogpostsDB() []Structure.Blogpost {
	const selectBlogPostSQL string = `
	SELECT *
	FROM blogpost
	  ;`
	var blogposts []Structure.Blogpost
	err := DB.Select(&blogposts, selectBlogPostSQL)
	if err != nil {
		log.Fatal(err)
	}

	const selectBlogPostTagsSQL string = `
		SELECT *
		FROM blogpost_tags
		WHERE blogpost_id = $1
		  ;`
	for _, blogpost := range blogposts {
		err := DB.Select(&blogpost.Tags, selectBlogPostTagsSQL, blogpost.Id)
		if err != nil {
			log.Fatal(err)
		}
	}
	return blogposts

}
