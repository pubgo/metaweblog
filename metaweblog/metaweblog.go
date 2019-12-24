package metaweblog

import (
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/g/xmlrpc"
	"github.com/pubgo/metaweblog/metaweblog/abc"
)

func WithAuth(username, password string) func(*metaweblog) {
	return func(m *metaweblog) {
		m.username = username
		m.password = password
	}
}

func WithUrl(url string) func(*metaweblog) {
	return func(m *metaweblog) {
		m.url = url
	}
}

func WithRetry(retryCount, retryWaitTime, retryMaxWaitTime int) func(*metaweblog) {
	return func(m *metaweblog) {
		m.retryCount = retryCount
		m.retryWaitTime = retryWaitTime
		m.retryMaxWaitTime = retryMaxWaitTime
	}
}

func WithTimeout(timeout int) func(*metaweblog) {
	return func(m *metaweblog) {
		m.timeout = timeout
	}
}

type metaweblog struct {
	retryCount       int  // default 3
	retryWaitTime    int  // default 5 second
	retryMaxWaitTime int  // default 20 second
	timeout          int  // default 60 second
	debug            bool // default true
	url              string
	username         string
	password         string
	client           *xmlrpc.Client
}

func (t *metaweblog) API() abc.IMetaweblogAPI {
	return &metaweblogImpl{c: t.client, username: t.username, password: t.password}
}

func (t *metaweblog) _init() {
	xerror.PanicT(t.url == "", "rpc api url is null")
	xerror.PanicT(t.username == "" || t.password == "", "username or password is null")

	if t.retryCount < 1 {
		t.retryCount = 3
	}

	if t.retryWaitTime < 1 {
		t.retryWaitTime = 5
	}

	if t.retryMaxWaitTime < 1 {
		t.retryMaxWaitTime = 20
	}

	if t.timeout < 1 {
		t.timeout = 60
	}

	t.debug = xenv.IsDebug()
	t.client = xmlrpc.NewClient(t.url)
}

func New(fn ...func(*metaweblog)) abc.IMetaweblog {
	_mb := &metaweblog{}
	for _, f := range fn {
		f(_mb)
	}
	_mb._init()
	return _mb
}
