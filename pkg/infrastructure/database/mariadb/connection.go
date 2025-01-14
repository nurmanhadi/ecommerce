package mariadb

import (
	"database/sql"
	"ecommerce/config"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	host := &config.Viper.Database.Mariadb.Host
	port := &config.Viper.Database.Mariadb.Port
	user := &config.Viper.Database.Mariadb.User
	password := &config.Viper.Database.Mariadb.Password
	dbname := &config.Viper.Database.Mariadb.Dbname
	max_idle_conns := &config.Viper.Database.Mariadb.MaxIdleConns
	max_open_conns := &config.Viper.Database.Mariadb.MaxOpenConns
	max_idle_time := &config.Viper.Database.Mariadb.ConnMaxIdleTime
	max_lifetime := &config.Viper.Database.Mariadb.ConnMaxLifetime
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", *user, *password, *host, *port, *dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(*max_idle_conns)
	db.SetMaxOpenConns(*max_open_conns)
	db.SetConnMaxIdleTime(time.Minute * time.Duration(*max_idle_time))
	db.SetConnMaxLifetime(time.Minute * time.Duration(*max_lifetime))

	log.Println("connected to database")
	return db
}
