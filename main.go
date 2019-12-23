package main

import (
	"fmt"
	"github.com/mattn/go-xmlrpc"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/metaweblog/cmds"
)

func main() {
	res, e := xmlrpc.Call(
		"http://rpc.cnblogs.com/metaweblog/bergus",
		"metaWeblog.getRecentPosts",
		"",
		0,
	)
	xerror.Panic(e)

	for k, v := range res.(xmlrpc.Array) {
		for _k, _v := range v.(xmlrpc.Struct) {
			fmt.Println(k, _k, _v)
		}
	}
	cmds.Execute()
}

// http://rpc.cnblogs.com/metaweblog/bergus#FileData
// metaWeblog.editPost
//metaWeblog.getCategories
//metaWeblog.getPost
//metaWeblog.getRecentPosts
//metaWeblog.newMediaObject
//metaWeblog.newPost
