package databases

import (
    "github.com/jinzhu/gorm"
    "github.com/lexkong/log"
    _ "github.com/go-sql-driver/mysql"
    "github.com/spf13/viper"
)

var SqlDB *gorm.DB

func Init() {
    var err error
    engine := viper.GetString("db.engine")
    dburl := viper.GetString("db.url")
    SqlDB, err = gorm.Open(engine, dburl)
    if err != nil {
        log.Fatal("Open database error", err)
    }
	SqlDB.SingularTable(true)
	//defer SqlDB.Close()

    //err = SqlDB.Ping()
    //if err != nil {
    //    log.Fatal(err.Error())
    //}
}
