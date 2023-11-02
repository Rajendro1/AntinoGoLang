package main

import (
	mysqldb "github.com/Rajendro1/AntinoGoLang/db/mysql"
	"github.com/Rajendro1/AntinoGoLang/route"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	
}
func main() {
	mysqldb.ConnectDB()
	route.HandleRequest()
}
