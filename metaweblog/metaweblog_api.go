package metaweblog

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/g/xmlrpc"
	"github.com/pubgo/metaweblog/metaweblog/abc"
	"github.com/pubgo/metaweblog/metaweblog/models"
	"net/http"
	"strings"
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

func (t *metaweblogImpl) Link2PostId(link string) (postId string, err error) {
	defer xerror.RespErr(&err)

	res := xerror.PanicErr(http.Get(link)).(*http.Response)
	defer res.Body.Close()
	xerror.PanicT(res.StatusCode/100 != 2, "status code error: %d %s", res.StatusCode, res.Status)

	// Load the HTML document
	doc := xerror.PanicErr(goquery.NewDocumentFromReader(res.Body)).(*goquery.Document)
	doc.Find("#topics .postDesc a").Each(func(i int, sec *goquery.Selection) {
		val, exists := sec.Attr("href")
		if !exists {
			return
		}

		if strings.Contains(val, "postid") {
			_val := strings.Split(val, "postid=")
			postId = _val[len(_val)-1]
		}
	})
	return
}

func (t *metaweblogImpl) Tags(blogId string) (data map[string]string, err error) {
	defer xerror.RespErr(&err)

	res := xerror.PanicErr(http.Get(fmt.Sprintf("https://www.cnblogs.com/%s/tag", blogId))).(*http.Response)
	defer res.Body.Close()
	xerror.PanicT(res.StatusCode/100 != 2, "status code error: %d %s", res.StatusCode, res.Status)

	data = make(map[string]string)
	// Load the HTML document
	doc := xerror.PanicErr(goquery.NewDocumentFromReader(res.Body)).(*goquery.Document)
	doc.Find("#MyTag1_dtTagList a").Each(func(i int, sec *goquery.Selection) {
		val, _ := sec.Attr("href")
		data[strings.TrimSpace(sec.Text())] = val
	})
	return
}

func (t *metaweblogImpl) GetPostList(blogId string) (plinks []models.PostLink, err error) {
	for i := 1; ; i++ {
		res := xerror.PanicErr(http.Get(fmt.Sprintf("https://www.cnblogs.com/%s/default.html?page=%d", blogId, i))).(*http.Response)
		defer res.Body.Close()
		xerror.PanicT(res.StatusCode/100 != 2, "status code error: %d %s", res.StatusCode, res.Status)

		// Load the HTML document
		doc := xerror.PanicErr(goquery.NewDocumentFromReader(res.Body)).(*goquery.Document)
		days := doc.Find("#mainContent .day")
		if len(days.Nodes) == 0 {
			break
		}

		days.Each(func(i int, sec *goquery.Selection) {
			var pls []models.PostLink
			sec.Find(".postTitle2").Each(func(i int, selection *goquery.Selection) {
				val, _ := selection.Attr("href")
				pls = append(pls, models.PostLink{Link: val, Title: strings.TrimSpace(selection.Text())})
			})
			sec.Find(".postDesc a").Each(func(i int, selection *goquery.Selection) {
				val, _ := selection.Attr("href")
				_val := strings.Split(strings.TrimSpace(val), "postid=")
				pls[i].Postid = _val[len(_val)-1]
			})
			plinks = append(plinks, pls...)
		})
	}
	return
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

func (t *metaweblogImpl) WpNewCategory(blogId string, category models.WpCategory) (c int, err error) {
	defer xerror.RespErr(&err)
	c = xerror.PanicErr(t.c.Call(wpNewCategory, blogId, t.username, t.password, category)).(int)
	return
}

func (t *metaweblogImpl) EditPost(postId string, post models.Post, publish bool) (b bool, err error) {
	defer xerror.RespErr(&err)
	b = xerror.PanicErr(t.c.Call(editPost, postId, t.username, t.password, post, publish)).(bool)
	return
}

func (t *metaweblogImpl) GetCategories(blogid string) (blogInfos []models.CategoryInfo, err error) {
	defer xerror.RespErr(&err)

	for _, v := range xerror.PanicErr(t.c.Call(getCategories, blogid, t.username, t.password)).(xmlrpc.Array) {
		_v := v.(xmlrpc.Map)
		blogInfos = append(blogInfos, models.CategoryInfo{
			Title:      _v["title"].(string),
			CategoryId: _v["categoryid"].(string),
		})
	}

	return
}

func (t *metaweblogImpl) GetPost(postId string) (_ models.GetPostResp, err error) {
	defer xerror.RespErr(&err)
	_p := xerror.PanicErr(t.c.Call(getPost, postId, t.username, t.password)).(xmlrpc.Map)
	p := models.GetPostResp{}
	xerror.Panic(json.Unmarshal(xerror.PanicErr(json.Marshal(_p)).([]byte), &p))
	return p, err
}

func (t *metaweblogImpl) GetRecentPosts(blogid string, numberOfPosts int) (ps []models.Post, err error) {
	defer xerror.RespErr(&err)

	for _, v := range xerror.PanicErr(t.c.Call(getRecentPosts, blogid, t.username, t.password, numberOfPosts)).(xmlrpc.Array) {
		_p := v.(xmlrpc.Map)
		ps = append(ps, models.Post{
			Title:       _p["title"].(string),
			Description: _p["description"].(string),
			WpSlug:      _p["wp_slug"].(string),
			MtKeywords:  _p["mt_keywords"].(string),
		})
	}

	return
}

func (t *metaweblogImpl) NewMediaObject(blogid string, file models.MediaObject) (m models.UrlData, err error) {
	defer xerror.RespErr(&err)

	_v := xerror.PanicErr(t.c.Call(newMediaObject, blogid, t.username, t.password, file)).(xmlrpc.Map)
	m.Url = _v["url"].(string)
	return
}

func (t *metaweblogImpl) NewPost(blogid string, post models.Post, publish bool) (s string, err error) {
	defer xerror.RespErr(&err)
	s = xerror.PanicErr(t.c.Call(newPost, blogid, t.username, t.password, post, publish)).(string)
	return
}
