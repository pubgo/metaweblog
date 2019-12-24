package tests

import (
	"fmt"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/metaweblog/metaweblog"
	"github.com/pubgo/metaweblog/metaweblog/abc"
	"runtime/debug"
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

func TestNewEdit(t *testing.T) {
	defer xerror.Resp(func(err xerror.IErr) {
		err.P()
		debug.PrintStack()
	})

	//fmt.Println(xerror.PanicErr(client.Call("metaWeblog.newPost",
	//	"222586",
	//	username,
	//	password,
	//	models.Post{
	//		Description: "测试 # sss" +
	//			"## njsnsjsnj",
	//		Title: "测试",
	//		//Categories: []string{cnst.Categoryid__5.Title, cnst.Categoryid__2.Title},
	//		//WpSlug:     "143504-0-35-111111",
	//		MtKeywords: "he,dd,ss,dnjs,dss",
	//	},
	//	true,
	//)))
}

func TestCategories(t *testing.T) {
	//fmt.Printf("%#v", xerror.PanicErr(client.Call(
	//	"metaWeblog.getCategories",
	//	"bergus",
	//	username,
	//	password,
	//)))
}

func TestUser(t *testing.T) {
	// "blogName":"白云辉", "blogid":"222586", "url":"https://www.cnblogs.com/bergus"
	fmt.Printf("%#v", xerror.PanicErr(client.API().BloggerGetUsersBlogs("")))
}
