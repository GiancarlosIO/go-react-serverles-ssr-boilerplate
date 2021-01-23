package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"mrn-portfolio/utils"
	"os"
)

type DB struct {

}

// Open function opens a database connection
func (db DB) Open() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=giancarlos dbname=mrnportfolio port=5432 sslmode=disable"
	}
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	utils.HandleError(err)

	return conn
}

