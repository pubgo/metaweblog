package tests

import (
	"fmt"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/metaweblog/metaweblog"
	"github.com/pubgo/metaweblog/metaweblog/abc"
	"github.com/pubgo/metaweblog/metaweblog/models"
	"testing"
)

var client abc.IMetaweblog
var username string
var password string

func init() {
	xerror.Panic(xenv.LoadFile("../.env"))
	username = xenv.GetEnv("username")
	password = xenv.GetEnv("password")
	client = metaweblog.New(
		metaweblog.WithUrl("http://rpc.cnblogs.com/metaweblog/bergus"),
		metaweblog.WithAuth(username, password),
	)
}

func TestGetPost(t *testing.T) {
	fmt.Printf("%#v", xerror.PanicErr(client.API().GetPost("12091487")))
}

func TestEditPost(t *testing.T) {
	fmt.Printf("%#v", xerror.PanicErr(client.API().EditPost(
		"12091487",
		models.Post{
			Description: "测试 # sss" +
				"## njsnsjsnj",
			Title: "测试",
			//Categories: []string{cnst.Categoryid__5.Title, cnst.Categoryid__2.Title},
			//WpSlug:     "143504-0-35-111111",
			MtKeywords: "he,dd,ss,dnjs,dss",
		},
		true,
	)))
}

func TestNewEdit(t *testing.T) {
	defer xerror.Assert()

	fmt.Printf("%#v", xerror.PanicErr(client.API().NewPost(
		"222586",
		models.Post{
			Description: "测试 # sss" +
				"## njsnsjsnj",
			Title: "测试",
			//Categories: []string{cnst.Categoryid__5.Title, cnst.Categoryid__2.Title},
			WpSlug:     "143504-0-35-111111",
			MtKeywords: "he,dd,ss,dnjs,dss",
		},
		true,
	)))
}

func TestCategories(t *testing.T) {
	fmt.Printf("%#v", xerror.PanicErr(client.API().GetCategories("bergus")))
}

func TestUser(t *testing.T) {
	// "blogName":"白云辉", "blogid":"222586", "url":"https://www.cnblogs.com/bergus"
	fmt.Printf("%#v", xerror.PanicErr(client.API().BloggerGetUsersBlogs("")))
}

func TestTag(t *testing.T) {
	//https://www.cnblogs.com/bergus/tag/
	fmt.Printf("%#v", xerror.PanicErr(client.API().Tags("bergus")))
}

func TestPostLink(t *testing.T) {
	//"https://www.cnblogs.com/bergus/default.html?page=300"
	fmt.Printf("%#v", xerror.PanicErr(client.API().GetPostList("bergus")))
}
