package config

import (
	"github.com/iot-synergy/synergy-common/plugins/casbin"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/iot-synergy/synergy-common/config"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf config.DatabaseConf
	CasbinConf   casbin.CasbinConf
	RedisConf    config.RedisConf
}
