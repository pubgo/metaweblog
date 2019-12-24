package tests

import (
	"fmt"
	"github.com/pubgo/metaweblog/metaweblog"

	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
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
			Description: `<!DOCTYPE html>
<html>
<head>
<title>开发趋势 - 幕布</title>
<meta charset="utf-8"/>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<meta name="renderer" content="webkit"/>
<meta name="author" content="mubu.com"/>
</head>
<body style="margin: 50px 20px;color: #333;font-family: SourceSansPro,-apple-system,BlinkMacSystemFont,'PingFang SC',Helvetica,Arial,'Microsoft YaHei',微软雅黑,黑体,Heiti,sans-serif,SimSun,宋体,serif">
<div class="export-wrapper"><div style="font-size: 22px; padding: 0 15px 0;"><div style="padding-bottom: 24px">开发趋势</div><div style="background: #e5e6e8; height: 1px; margin-bottom: 20px;"></div></div><ul style="list-style: disc outside;"><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">前端</span><ul class="children" style="list-style: disc outside; padding-bottom: 4px;"><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">前端移动端趋于合并，dart统一了移动端(ios, 安卓), 统一了web开发(结合vue,angular), 同时能够转化为js以及开发API server，所以，只需要dart就什么都可以了</span></li><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;"></span></li></ul></li><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">后端</span><ul class="children" style="list-style: disc outside; padding-bottom: 4px;"><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">云应用开发，go基本一家独大</span></li><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">爬虫nodejs是一批黑马</span></li><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">pyhton发力AI</span></li></ul></li><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">中间件</span><ul class="children" style="list-style: disc outside; padding-bottom: 4px;"><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">redis</span></li><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">rabbitmq</span></li><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">mongodb</span></li><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">mysql</span></li><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">有序，持久化，分布式事务</span></li></ul></li><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">数据库</span><ul class="children" style="list-style: disc outside; padding-bottom: 4px;"><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">数据存储归一，现有通过tidb，蟑螂数据库，都可以分布式存储查询</span></li></ul></li><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">大数据</span><ul class="children" style="list-style: disc outside; padding-bottom: 4px;"><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">现有平台hadoop，spark，flink</span></li><li style="line-height: 24px;"><span class="content mubu-node" style="line-height: 24px; min-height: 24px; font-size: 16px; padding: 2px 0px; display: inline-block; vertical-align: top;">批处理，流处理</span></li></ul></li></ul></div>

</body>
</html>`,
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

func TestLink2PostId(t *testing.T) {
	fmt.Printf("%#v\n", xerror.PanicErr(client.API().Link2PostId("https://www.cnblogs.com/bergus/p/143504-0-35-111111.html")))
}
