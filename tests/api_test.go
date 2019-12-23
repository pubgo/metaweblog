package tests

import (
	"fmt"
	"github.com/kolo/xmlrpc"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/metaweblog/metaweblog/models"
	"testing"
)

var client *xmlrpc.Client
var username string
var password string

func init() {
	xerror.Panic(xenv.LoadFile("../.env"))
	username = xenv.GetEnv("username")
	password = xenv.GetEnv("password")
	//client = xmlrpc.NewClient("http://rpc.cnblogs.com/metaweblog/bergus")
}

func TestNewEdit(t *testing.T) {
	//fmt.Println(client.Call("metaWeblog.newPost",
	//	"222586",
	//	username,
	//	password,
	//	models.Post{
	//		Description: "测试 # sss" +
	//			"## njsnsjsnj",
	//		Title:      "测试",
	//		//Categories: []string{cnst.Categoryid__5.Title, cnst.Categoryid__2.Title},
	//		//WpSlug:     "143504-0-35-111111",
	//		//MtKeywords: "he,dd,ss,dnjs,dss",
	//	},
	//	true,
	//))

	client, _ := xmlrpc.NewClient("http://rpc.cnblogs.com/metaweblog/bergus", nil)
	var data interface{}
	xerror.Panic(client.Call("metaWeblog.newPost",
		[]interface{}{
			"222586",
			username,
			password,
			models.Post{
				Description: "测试 # sss" +
					"## njsnsjsnj",
				Title: "测试",
				//Categories: []string{cnst.Categoryid__5.Title, cnst.Categoryid__2.Title},
				//WpSlug:     "143504-0-35-111111",
				//MtKeywords: "he,dd,ss,dnjs,dss",
			},
			true,
		}, &data))

	fmt.Println(data)
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
	//fmt.Printf("%#v", xerror.PanicErr(client.Call(
	//	"blogger.getUsersBlogs",
	//	"",
	//	username,
	//	password,
	//)))
}