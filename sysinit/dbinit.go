package sysinit

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "mbook/models"
)

func dbInit(alias string) {
	dbAlias := alias
	if alias == "w" || alias == "default" || len(alias) <= 0{
		dbAlias = "default"
		alias = "w"
	}

	//数据库信息
	dbHost := beego.AppConfig.String(fmt.Sprintf("db_%s_host", alias))
	dbPort := beego.AppConfig.String(fmt.Sprintf("db_%s_port", alias))
	dbName := beego.AppConfig.String(fmt.Sprintf("db_%s_username", alias))
	dbPassword :=  beego.AppConfig.String(fmt.Sprintf("db_%s_password", alias))
	dbDataBase :=  beego.AppConfig.String(fmt.Sprintf("db_%s_database", alias))

	_ = orm.RegisterDataBase(dbAlias, "mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbName, dbPassword, dbHost, dbPort, dbDataBase), 30)

	// 自动建立表
	isDev := beego.AppConfig.String("runmode") == "dev"
	if alias == "w" {
		_ = orm.RunSyncdb("default", false, isDev)
	}
}