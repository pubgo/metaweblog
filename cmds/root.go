package cmds

import (
	"github.com/pubgo/g/xcmd"
	"github.com/pubgo/g/xerror"
	"github.com/pubgo/metaweblog/cmds/metaweblog"
	"github.com/pubgo/metaweblog/version"
)

const Service = "metaweblog"
const EnvPrefix = "MW"

// Execute exec
var Execute = xcmd.Init(EnvPrefix, func(cmd *xcmd.Command) {
	defer xerror.Assert()

	cmd.Use = Service
	cmd.Version = version.Version

	cmd.AddCommand(
		metaweblog.Version(),
	)

})
