package abc

import "github.com/pubgo/metaweblog/metaweblog/models"

// doc http://rpc.cnblogs.com/metaweblog/bergus

type IMetaweblogAPI interface {
	EditPost(postId string, post models.Post, publish bool) (bool, error)
	GetCategories(blogId string) ([]models.CategoryInfo, error)
	GetPost(postId string) (models.GetPostResp, error)
	GetRecentPosts(blogId string, numberOfPosts int) ([]models.Post, error)
	NewMediaObject(blogId string, file models.MediaObject) (models.UrlData, error)
	NewPost(blogId string, post models.Post, publish bool) (string, error)
	BloggerDeletePost(appKey, postId string, publish bool) (bool, error)
	BloggerGetUsersBlogs(appKey string) ([]models.BlogInfo, error)
	WpNewCategory(blogId string, category models.WpCategory) (int, error)
	Tags(blogId string) (map[string]string, error)
	GetPostList(blogId string) ([]models.PostLink, error)
}

type IMetaweblog interface {
	API() IMetaweblogAPI
}
