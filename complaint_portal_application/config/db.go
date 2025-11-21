package config

import (
	"database/sql"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

var DB *gorm.DB

func SetupDB() (*gorm.DB, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if sslmode == "" {
		sslmode = "disable"
	}

	postgresDSN := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=postgres sslmode=%s",
		host, port, user, password, sslmode,
	)

	sqlDB, err := sql.Open("postgres", postgresDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %v", err)
	}

	_, err = sqlDB.Exec("CREATE DATABASE " + dbname)
	if err != nil && !isDatabaseExistsError(err, dbname) {
		return nil, fmt.Errorf("failed to create database: %v", err)
	}

	sqlDB.Close()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed connecting to database: %v", err)
	}

	return db, nil
}

func isDatabaseExistsError(err error, dbname string) bool {
	return err != nil && err.Error() == fmt.Sprintf(`pq: database "%s" already exists`, dbname)
}
