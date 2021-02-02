package conf

import (
	"flag"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	xtime "github.com/go-kratos/kratos/pkg/time"
	// "github.com/go-xorm/xorm"
)

// conf init
var (
	Conf = &Config{}
)

type Config struct {
	Version	string  	`yaml:"version"`
	Web		WebSession	`yaml:"web"`
	Log             	*log.Config                  `yaml:"log"`
	ORM             	ormConfig                  `yaml:"orm"`
	//GrpcServer      *warden.ServerConfig         `yaml:"grpcserver"`
	//GrpcClient      *client.RPCClientConfig      `yaml:"grpcclient"`

	//Redis  *RedisSection            `yaml:"redisc"`

}


type ormConfig struct {
	DSN         string         `yaml:"dsn"`			// data source name.
	Active      int            `yaml:"active"`		// pool
	Idle        int            `yaml:"idle"`		// pool
	IdleTimeout xtime.Duration `yaml:"idleTimeout"`	// connect max life time.
}


type WebSession struct {
	Addr	string				`yaml:"addr"`
	MaxListen int				`yaml:"maxListen"`
	Timeout  xtime.Duration 	`yaml:"timeout"`
}

// Init .
func Init() error {
	flag.Parse()
	if err :=paladin.Init(); err != nil {
		log.Error("conf.Init() error(%v)", err)
		panic(err)
	}

	if err := paladin.Get("conf.yml").UnmarshalYAML(&Conf); err != nil {
		if err != paladin.ErrNotExist {
			panic(err)
		}
	}

	return nil
}
