package database

import (
	"database/sql"
	"fmt"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

var MyWinnerDB *gorp.DbMap

func InitDB() {
	driver := "mysql"
	spec := "root:password@tcp(127.0.0.1:3306)/winner94"
	LOCAL_TYC2, err := sql.Open(driver, spec)
	if err != nil {
		fmt.Println("initDB: ", err.Error())
	}
	MyWinnerDB = &gorp.DbMap{Db: LOCAL_TYC2, Dialect: gorp.OracleDialect{}}
}
