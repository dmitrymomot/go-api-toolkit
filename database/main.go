package database

import (
	"github.com/jinzhu/gorm"
	// mysql init
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connect returns database connect
func Connect(c Configer) (*gorm.DB, error) {
	db, err := gorm.Open(c.ConnectionStringWithDriver())
	if err != nil {
		return nil, err
	}
	db.LogMode(c.Logging())
	db.DB().SetConnMaxLifetime(c.ConnectionMaxLifetime())
	db.DB().SetMaxOpenConns(c.MaxOpenConnections())
	db.DB().SetMaxIdleConns(c.MaxIdleConnections())
	if err = db.DB().Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

// Migrate models
func Migrate(db *gorm.DB, models ...interface{}) {
	for _, model := range models {
		if db.HasTable(model) == false {
			db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(model)
		}
		db.AutoMigrate(model)
	}
}
