package conf

import (

	"github.com/jinzhu/gorm"

)

type Config struct {
	Version	string  	`yaml:"version"`
	Web		WebSession	`yaml:"web"`
	// Log             *log.Config                  `yaml:"log"`
	ORM             *orm.Config                  `yaml:"orm"`
	//GrpcServer      *warden.ServerConfig         `yaml:"grpcserver"`
	//GrpcClient      *client.RPCClientConfig      `yaml:"grpcclient"`

	//Redis  *RedisSection            `yaml:"redisc"`

}

type WebSession struct {
	Addr	string		`yaml:"addr"`
	MaxListen int		`yaml:"maxListen"`
	Timeout string 		`yaml:"timeout"`
}


