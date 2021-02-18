package dao

import (
<<<<<<< HEAD
	"gin-demo/conf"
	"gin-demo/model"
	"github.com/go-kratos/kratos/pkg/log"
	 "github.com/go-xorm/xorm"
=======
	"fmt"
	"gin-test-demo/conf"
	"gin-test-demo/model"
	"github.com/go-kratos/kratos/pkg/log"
	"github.com/go-xorm/xorm"
	"time"
	_ "github.com/go-sql-driver/mysql"
>>>>>>> 1f5125ac6e2dac1b553f1b3f48563f86940096d6
)

// Dao .
type Dao struct {
	DB	*xorm.Engine
}

// New  database xorm client.
func New(conf *conf.Config) (dao *Dao){
<<<<<<< HEAD
	db, err := xorm.NewEngine("mysql", conf.ORM.DSN)
=======
	fmt.Printf("%v", conf.ORM)
	db, err:= xorm.NewEngine("mysql", conf.ORM.DSN)
>>>>>>> 1f5125ac6e2dac1b553f1b3f48563f86940096d6
	if err != nil {
		log.Error("db dsn(%s) error: %v", conf.ORM.DSN, err)
		panic(err)
	}
<<<<<<< HEAD
	db.DB().SetMaxIdleConns(conf.ORM.Idle)
	db.DB().SetMaxOpenConns(conf.ORM.Active)
=======

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


>>>>>>> 1f5125ac6e2dac1b553f1b3f48563f86940096d6
	return
}

// initORM .
<<<<<<< HEAD
func (d *Dao)initORM (){
=======
func (d *Dao) initORM (){
>>>>>>> 1f5125ac6e2dac1b553f1b3f48563f86940096d6
	if err :=d.DB.Sync2(&model.User{});err != nil {
		log.Error("DAO.initORM ERROR ：%v", err)
		panic(err)
	}
	d.DB.ShowSQL(true)
}