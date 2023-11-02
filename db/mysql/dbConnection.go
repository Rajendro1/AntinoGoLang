package mysqldb

import (
	"database/sql"
	"log"

	"github.com/Rajendro1/AntinoGoLang/config"
)

var (
	DB  *sql.DB
	err error
)

func ConnectDB() {
	var MYSQL_URL_WITHOUT_DATABASE string = config.DbConfig().DB_USERNAME + ":" + config.DbConfig().DB_PASSWORD + "@tcp(" + config.DbConfig().DB_HOST + ":" + config.DbConfig().DB_PORT + ")/"

	var MYSQL_URL_WITH_DATABASE string = config.DbConfig().DB_USERNAME + ":" + config.DbConfig().DB_PASSWORD + "@tcp(" + config.DbConfig().DB_HOST + ":" + config.DbConfig().DB_PORT + ")/" + config.DB_NAME + "?multiStatements=true"

	CreateMySQldatabase(MYSQL_URL_WITHOUT_DATABASE)

	DB, err = sql.Open("mysql", MYSQL_URL_WITH_DATABASE)
	if err != nil {
		log.Fatal(err)
	}
	dbMigration()
	DB.SetMaxOpenConns(10)

}
func dbMigration() {
	if _, err := DB.Exec(TableCreation); err != nil {
		log.Println("dbMigration err: ", err.Error())
	}
}
func CreateMySQldatabase(url string) {
	databseCon, err := sql.Open("mysql", url)
	if err != nil {
		log.Println("Error To Connect Databae")
	}
	if _, dbExecErr := databseCon.Exec(DbCreationQuery); dbExecErr != nil {
		log.Println("**********MySqlDatabse********************")
		log.Println(dbExecErr.Error())
		log.Println("********* MySqlDatabse********************")
	}
}
