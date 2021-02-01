package dao

import (
	"gin-demo/conf"
	"gin-demo/model"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/go-xorm/xorm"
	"time"
)

// Dao .
type Dao struct {
	DB	*xorm.Engine
}

// New  database xorm client.
func New(conf *conf.Config) (dao *Dao){
	db, err := xorm.NewEngine("mysql", conf.ORM.DSN)
	if err != nil {
		log.Error("db dsn(%s) error: %v", conf.ORM.DSN, err)
		panic(err)
	}
	db.DB().SetMaxIdleConns(conf.ORM.Idle)
	db.DB().SetMaxOpenConns(conf.ORM.Active)
	db.DB().SetConnMaxLifetime(time.Duration(conf.ORM.IdleTimeout) / time.Second)


	return
}

// initORM .
func (d *Dao)initORM (){
	if err :=d.DB.Sync2(&model.User{});err != nil {
		log.Error("DAO.initORM ERROR ：%v", err)
		panic(err)
	}
	d.DB.ShowSQL(true)
}