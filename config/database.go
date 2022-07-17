package config

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func InitDatabase() *sqlx.DB {
	// Database
	driver := viper.GetString("db.driver")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.datbase"),
	)

	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		panic(err)
	}

	// Connection pool
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
