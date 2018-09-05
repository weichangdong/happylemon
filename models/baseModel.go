package models

import (
	"fmt"
	"happylemon/conf"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
)

var MyDb gorose.Connection

type BaseModels struct {
}

// MysqlPool mysql pool
func InitMysqlPool(conf conf.Conf) {
	host := conf.Mysql.Host
	username := conf.Mysql.UserName
	password := conf.Mysql.Password
	port := conf.Mysql.Port
	database := conf.Mysql.Database
	maxOpenConns := conf.Mysql.MaxOpenConns
	maxIdleConns := conf.Mysql.MaxIdleConns

	var dbConfig = map[string]interface{}{
		"Default":         "mysql_j20",  // 默认数据库配置
		"SetMaxOpenConns": maxOpenConns, // (连接池)最大打开的连接数，默认值为0表示不限制
		"SetMaxIdleConns": maxIdleConns, // (连接池)闲置的连接数, 默认1
		"Connections": map[string]map[string]string{
			"mysql_j20": map[string]string{ // 定义名为 mysql_dev 的数据库配置
				"host":     host,      // 数据库地址
				"username": username,  // 数据库用户名
				"password": password,  // 数据库密码
				"port":     port,      // 端口
				"database": database,  // 链接的数据库名字
				"charset":  "utf8mb4", // 字符集
				"protocol": "tcp",     // 链接协议
				"prefix":   "",        // 表前缀
				"driver":   "mysql",   // 数据库驱动(mysql,sqlite,postgres,oracle,mssql)
			}}}
	var err error
	MyDb, err = gorose.Open(dbConfig)
	if err != nil {
		fmt.Println(err)
		panic(err)

	}
}
