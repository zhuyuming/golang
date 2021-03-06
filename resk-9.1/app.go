package resk

import (
	"imooc.com/resk/apis/gorpc"
	_ "imooc.com/resk/apis/gorpc"
	_ "imooc.com/resk/apis/web"
	_ "imooc.com/resk/core/accounts"
	_ "imooc.com/resk/core/envelopes"
	"imooc.com/resk/infra"
	"imooc.com/resk/infra/base"
	"imooc.com/resk/jobs"
	_ "imooc.com/resk/public/ui"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.GoRPCStarter{})
	infra.Register(&gorpc.GoRpcApiStarter{})
	infra.Register(&jobs.RefundExpiredJobStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
