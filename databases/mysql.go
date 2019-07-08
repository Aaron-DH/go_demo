package databases

import (
    "database/sql"
    "log"
    _ "github.com/go-sql-driver/mysql"
    "github.com/spf13/viper"
)

var SqlDB *sql.DB

func Init() {
    var err error
    engine := viper.GetString("engine")
    dburl := viper.GetString("url")
    SqlDB, err = sql.Open(engine, dburl)
    if err != nil {
        log.Fatal(err.Error())
    }

    err = SqlDB.Ping()
    if err != nil {
        log.Fatal(err.Error())
    }
}
