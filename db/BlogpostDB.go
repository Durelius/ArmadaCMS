package db

import (
	"ArmadaCMS/Structure"
)

func InsertBlogpostDB(blogpost Structure.NewBlogpost, userId int) error {
	// 1. Create the Blogpost
	post := Structure.Blogpost{
		Text:   blogpost.Text,
		Title:  blogpost.Title,
		Author: blogpost.Author,
		UserID: userId,
	}

	// Insert the blog post into the database
	if err := DB.Create(&post).Error; err != nil {
		return err
	}

	// 2. Create BlogpostTags and associate them with the Blogpost
	var tagEntries []Structure.BlogpostTag
	for _, tag := range blogpost.Tags {
		// Create a BlogpostTag for each tag
		tagEntries = append(tagEntries, Structure.BlogpostTag{
			BlogpostID: post.ID, // Associate with the created blog post
			Tag:        tag,
		})
	}

	// Insert the tags into the database
	if len(tagEntries) > 0 {
		if err := DB.Create(&tagEntries).Error; err != nil {
			return err
		}
	}

	// Success
	return nil
}

func GetAllBlogpostsDB() ([]Structure.Blogpost, error) {
	var blogposts []Structure.Blogpost

	if err := DB.Preload("Tags").Find(&blogposts).Error; err != nil {
		return nil, err
	}

	return blogposts, nil
}
