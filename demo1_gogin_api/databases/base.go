package databases

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

var SqlDB *gorm.DB

func Init() {
	engine := viper.GetString("db.engine")
	if engine == "mysql" {
		SqlDB = OpenMySQLDB()
	}
}
