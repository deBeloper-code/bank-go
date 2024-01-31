// internal/adapter/database/database.go
package database

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"

	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
	dsn  = "host=localhost user=root password=secret dbname=bank_db port=5432 sslmode=disable"
)

func GetDBConnection() *gorm.DB {
	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to DB: %v\n", err)
			os.Exit(1)
		}
		// Pool connection
		sqlDB, err := db.DB()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to pool connection to DB: %v\n", err)
			os.Exit(1)
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
	})

	return db
}
