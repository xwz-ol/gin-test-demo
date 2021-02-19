package dao

import (

	"fmt"
	"gin-test-demo/conf"
	"gin-test-demo/model"
	"github.com/go-kratos/kratos/pkg/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
	_ "github.com/go-sql-driver/mysql"

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