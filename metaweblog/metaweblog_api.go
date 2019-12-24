package metaweblog

import (
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/g/xmlrpc"
	"github.com/pubgo/metaweblog/metaweblog/abc"
	"github.com/pubgo/metaweblog/metaweblog/models"
)

var (
	_                    abc.IMetaweblogAPI = (*metaweblogImpl)(nil)
	editPost                                = "metaWeblog.editPost"
	getCategories                           = "metaWeblog.getCategories"
	getPost                                 = "metaWeblog.getPost"
	getRecentPosts                          = "metaWeblog.getRecentPosts"
	newMediaObject                          = "metaWeblog.newMediaObject"
	newPost                                 = "metaWeblog.newPost"
	bloggerDeletePost                       = "blogger.deletePost"
	bloggerGetUsersBlogs                    = "blogger.getUsersBlogs"
	wpNewCategory                           = "wp.newCategory"
)

type metaweblogImpl struct {
	c        *xmlrpc.Client
	username string
	password string
}

func (t *metaweblogImpl) BloggerDeletePost(appKey, postId string, publish bool) (b bool, err error) {
	defer xerror.RespErr(&err)
	b = xerror.PanicErr(t.c.Call(bloggerDeletePost, appKey, postId, t.username, t.password, publish)).(bool)
	return
}

func (t *metaweblogImpl) BloggerGetUsersBlogs(appKey string) (blogInfos []models.BlogInfo, err error) {
	defer xerror.RespErr(&err)

	for _, v := range xerror.PanicErr(t.c.Call(bloggerGetUsersBlogs, appKey, t.username, t.password)).(xmlrpc.Array) {
		_v := v.(xmlrpc.Map)
		blogInfos = append(blogInfos, models.BlogInfo{
			BlogId:   _v["blogid"].(string),
			URL:      _v["url"].(string),
			BlogName: _v["blogName"].(string),
		})
	}

	return
}

func (t *metaweblogImpl) WpNewCategory(blogId string, category models.WpCategory) (int, error) {
	t.c.Call(bloggerDeletePost, blogId, t.username, t.password, category)
	panic("")
}

func (t *metaweblogImpl) EditPost(postId string, post models.Post, publish bool) (interface{}, error) {
	return t.c.Call(editPost, postId, t.username, t.password, post, publish)
	panic("")
}

func (t *metaweblogImpl) GetCategories(blogid string) ([]models.CategoryInfo, error) {
	t.c.Call(bloggerDeletePost, blogid, t.username, t.password)
	panic("")
}

func (t *metaweblogImpl) GetPost(postId string) (models.Post, error) {
	t.c.Call(bloggerDeletePost, postId, t.username, t.password)
	panic("")
}

func (t *metaweblogImpl) GetRecentPosts(blogid string, numberOfPosts int) ([]models.Post, error) {
	t.c.Call(bloggerDeletePost, blogid, t.username, t.password, numberOfPosts)
	panic("")
}

func (t *metaweblogImpl) NewMediaObject(blogid string, file models.MediaObject) (models.UrlData, error) {
	t.c.Call(bloggerDeletePost, blogid, t.username, t.password, file)
	panic("")
}

func (t *metaweblogImpl) NewPost(blogid string, post models.Post, publish bool) (string, error) {
	t.c.Call(bloggerDeletePost, blogid, t.username, t.password, post, publish)
	panic("")
}
