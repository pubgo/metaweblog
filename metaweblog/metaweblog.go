package metaweblog

import (
	"github.com/mattn/go-xmlrpc"
	"github.com/pubgo/g/xenv"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/metaweblog/metaweblog/abc"
)

type metaweblog struct {
	RetryCount       int  // default 3
	RetryWaitTime    int  // default 5 second
	RetryMaxWaitTime int  // default 20 second
	Timeout          int  // default 60 second
	Debug            bool // default true
	Url              string
	Username         string
	Password         string
	client           *xmlrpc.Client
}

func (t *metaweblog) API() abc.IMetaweblog {
	return nil
}

func (t *metaweblog) _init() {
	if t.RetryCount < 1 {
		t.RetryCount = 3
	}

	if t.RetryWaitTime < 1 {
		t.RetryWaitTime = 5
	}

	if t.RetryMaxWaitTime < 1 {
		t.RetryMaxWaitTime = 20
	}

	if t.Timeout < 1 {
		t.Timeout = 60
	}

	t.Debug = xenv.IsDebug()

	xerror.PanicT(t.Url == "", "rpc api url is null")
	t.client = xmlrpc.NewClient(t.Url)
}

func New() *metaweblog {
	_mb := &metaweblog{}
	_mb._init()
	return _mb
}
