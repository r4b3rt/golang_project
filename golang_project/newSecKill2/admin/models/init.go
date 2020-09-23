package models

import (
	"fmt"
	"newSecKill2/common/initall"
	"github.com/astaxie/beego/orm"
	"github.com/beego/admin/src/models"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  orm.Ormer
	SecKillConf         initall.ConfigAll
)

func init (){
	// 需要在init中注册定义的model
	orm.RegisterModel(new(SecKillProduct))
	//数据库连接
	models.Connect()
	// 自动建表
	orm.RunSyncdb("default", false, true)

	DB = orm.NewOrm()

	if err := initAll(); err != nil {
		panic(fmt.Sprintln("init database failed, err:%v", err))
	}

}

func initAll() (err error) {
	//  记载秒杀系统配置信息
	if SecKillConf, err = initall.InitConfig(); err != nil {
		return
	}
	return
}