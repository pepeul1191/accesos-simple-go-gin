package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Database() *gorm.DB {
	db, err := gorm.Open("sqlite3", "db/ubicaciones.db")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
