package svc

import (
	i18n2 "github.com/iot-synergy/oms-core/api/internal/i18n"
	"github.com/iot-synergy/oms-core/api/internal/middleware"
	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/iot-synergy/synergy-common/utils/captcha"
	"github.com/iot-synergy/synergy-job/jobclient"
	"github.com/iot-synergy/synergy-message-center/mcmsclient"
	"github.com/mojocn/base64Captcha"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/iot-synergy/oms-core/api/internal/config"
	"github.com/iot-synergy/oms-core/rpc/coreclient"

	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Authority   rest.Middleware
	CoreRpc     coreclient.Core
	JobRpc      jobclient.Job
	McmsRpc     mcmsclient.Mcms
	Redis       redis.UniversalClient
	Casbin      *casbin.Enforcer
	Trans       *i18n.Translator
	Captcha     *base64Captcha.Captcha
	BanRoleData map[string]bool // ban role means the role status is not normal
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := c.RedisConf.MustNewUniversalRedis()

	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(c.DatabaseConf.Type, c.DatabaseConf.GetDSN(),
		c.RedisConf)

	var trans *i18n.Translator
	if c.I18nConf.Dir != "" {
		trans = i18n.NewTranslatorFromFile(c.I18nConf)
	} else {
		trans = i18n.NewTranslator(i18n2.LocaleFS)
	}

	svc := &ServiceContext{
		Config:  c,
		CoreRpc: coreclient.NewCore(zrpc.NewClientIfEnable(c.CoreRpc)),
		JobRpc:  jobclient.NewJob(zrpc.NewClientIfEnable(c.JobRpc)),
		McmsRpc: mcmsclient.NewMcms(zrpc.NewClientIfEnable(c.McmsRpc)),
		Captcha: captcha.MustNewOriginalRedisCaptcha(c.Captcha, rds),
		Redis:   rds,
		Casbin:  cbn,
		Trans:   trans,
	}

	err := svc.LoadBanRoleData()
	logx.Must(err)

	svc.Authority = middleware.NewAuthorityMiddleware(cbn, rds, trans, svc.BanRoleData).Handle

	return svc
}
