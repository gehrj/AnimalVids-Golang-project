package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// this is using xorm to setup our database connection
func Connect(host string, port string, user string, pass string, database string, options string) (*xorm.Engine, error) {
	return xorm.NewEngine("mysql", user+":"+pass+"@tcp("+host+")/"+database+"?charset=utf8&"+options)
}
