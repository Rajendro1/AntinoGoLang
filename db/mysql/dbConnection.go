package mysqldb

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Rajendro1/AntinoGoLang/config"
)

var (
	DB  *sql.DB
	err error
)

func ConnectDB() {
	var MYSQL_URL_WITH_DATABASE string = config.DbConfig().DB_USERNAME + ":" + config.DbConfig().DB_PASSWORD + "@tcp(" + config.DbConfig().DB_HOST + ":" + config.DbConfig().DB_PORT + ")/" + config.DB_NAME + "?multiStatements=true"
	DB, err = sql.Open("mysql", MYSQL_URL_WITH_DATABASE)
	if err != nil {
		log.Fatal(err)
	}
	dbMigration()
	DB.SetMaxOpenConns(10)

}
func dbMigration() {
	fmt.Println(DbAndTableCreation)
	if _, err := DB.Exec(DbAndTableCreation); err != nil {
		log.Println("dbMigration err: ", err.Error())
	}
}
