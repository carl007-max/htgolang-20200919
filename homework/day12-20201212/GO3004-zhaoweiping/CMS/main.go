package main

import (
	"CMS/models"
	_ "CMS/routers"
	"CMS/utils"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// driverName := "mysql"
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local&charset=utf8mb4",
	// 	beego.AppConfig.String("db::mysqluser"),
	// 	beego.AppConfig.String("db::mysqlpasswd"),
	// 	beego.AppConfig.String("db::mysqlip"),
	// 	beego.AppConfig.String("db::mysqlport"),
	// 	beego.AppConfig.String("db::mysqldb"),
	// )
	// // dsn := "root:root@tcp(127.0.0.1:3306)/user?parseTime=true&loc=Local&charset=utf8mb4"

	// if err := models.InitDb(driverName, dsn); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// // defer models.CloseDb()
	utils.InitLogger()
	utils.InitDB()
}

func main() {
	// services.InitLogger()
	defer models.CloseDb()
	beego.Run()
}
