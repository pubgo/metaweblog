package abc

import "github.com/pubgo/metaweblog/metaweblog/models"

// doc http://rpc.cnblogs.com/metaweblog/bergus

type IMetaweblog interface {
	EditPost(postId, username, password string, post models.Post, publish bool) (interface{}, error)
	GetCategories(blogid, username, password string) ([]models.CategoryInfo, error)
	GetPost(postId, username, password string) (models.Post, error)
	GetRecentPosts(blogid, username, password string, numberOfPosts int) ([]models.Post, error)
	NewMediaObject(blogid, username, password string, file models.MediaObject) (models.UrlData, error)
	NewPost(blogid, username, password string, post models.Post, publish bool) (string, error)
	BloggerDeletePost(appKey, postid, username, password string, publish bool) (bool, error)
	BloggerGetUsersBlogs(appKey, username, password string) ([]models.BlogInfo, error)
	WpNewCategory(blogId, username, password string, category models.WpCategory) (int, error)
}
