package databases

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

var SqlDB *sql.DB

func init() {
    var err error
    SqlDB, err = sql.Open("mysql", "root:R0otAwc10ud@tcp(172.16.2.100:8090)/ailab?charset=utf8")
    if err != nil {
        log.Fatal(err.Error())
    }

    err = SqlDB.Ping()
    if err != nil {
        log.Fatal(err.Error())
    }
}
