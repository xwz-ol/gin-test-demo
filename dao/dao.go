package dao

import (
	"fmt"
	"gin-test-demo/conf"
	"gin-test-demo/dao/redisc"
	"gin-test-demo/model"
	"github.com/garyburd/redigo/redis"
	"github.com/go-kratos/kratos/pkg/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

// Dao .
type Dao struct {
	DB	*xorm.Engine
	RedisConnPool *redis.Pool
}

// New  database xorm client.
func New(conf *conf.Config) (dao *Dao){

	fmt.Println(conf.ORM)
	fmt.Println(conf.Redis)
	db, err := xorm.NewEngine("mysql", conf.ORM.DSN)

	if err != nil {
		log.Error("db dsn(%s) error: %v", conf.ORM.DSN, err)
		panic(err)
	}

	tmp := &xorm.Engine{}

	dao = &Dao{DB:tmp}

	if db == tmp {
		return
	} else {
		dao.DB = db
	}

	dao.initORM()

	db.DB().SetMaxIdleConns(conf.ORM.Idle)
	db.DB().SetMaxOpenConns(conf.ORM.Active)
	db.DB().SetConnMaxLifetime(time.Duration(conf.ORM.IdleTimeout) / time.Second)
	redisc.InitRedis()
	return
}

// initORM .
func (d *Dao) initORM (){
	if err :=d.DB.Sync2(&model.User{});err != nil {
		log.Error("DAO.initORM ERROR ：%v", err)
		panic(err)
	}
	d.DB.ShowSQL(true)
}