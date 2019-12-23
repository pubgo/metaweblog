package metaweblog

import (
	"github.com/mattn/go-xmlrpc"
	"github.com/pubgo/metaweblog/metaweblog/abc"
	"github.com/pubgo/metaweblog/metaweblog/models"
)

var (
	_                    abc.IMetaweblog = (*metaweblogImpl)(nil)
	editPost                             = "metaWeblog.editPost"
	getCategories                        = "metaWeblog.getCategories"
	getPost                              = "metaWeblog.getPost"
	getRecentPosts                       = "metaWeblog.getRecentPosts"
	newMediaObject                       = "metaWeblog.newMediaObject"
	newPost                              = "metaWeblog.newPost"
	bloggerDeletePost                    = "blogger.deletePost"
	bloggerGetUsersBlogs                 = "blogger.getUsersBlogs"
	wpNewCategory                        = "wp.newCategory"
)

type metaweblogImpl struct {
	c *xmlrpc.Client
}

func (t *metaweblogImpl) BloggerDeletePost(appKey, postid, username, password string, publish bool) (bool, error) {
	panic("implement me")
}

func (t *metaweblogImpl) BloggerGetUsersBlogs(appKey, username, password string) ([]models.BlogInfo, error) {
	panic("implement me")
}

func (t *metaweblogImpl) WpNewCategory(blogId, username, password string, category models.WpCategory) (int, error) {
	panic("implement me")
}

func (t *metaweblogImpl) EditPost(postId, username, password string, post models.Post, publish bool) (interface{}, error) {
	return t.c.Call(editPost, postId, username, password, post, publish)
}

func (t *metaweblogImpl) GetCategories(blogid, username, password string) ([]models.CategoryInfo, error) {
	panic("implement me")
}

func (t *metaweblogImpl) GetPost(postId, username, password string) (models.Post, error) {
	panic("implement me")
}

func (t *metaweblogImpl) GetRecentPosts(blogid, username, password string, numberOfPosts int) ([]models.Post, error) {
	panic("implement me")
}

func (t *metaweblogImpl) NewMediaObject(blogid, username, password string, file models.MediaObject) (models.UrlData, error) {
	panic("implement me")
}

func (t *metaweblogImpl) NewPost(blogid, username, password string, post models.Post, publish bool) (string, error) {
	panic("implement me")
}
