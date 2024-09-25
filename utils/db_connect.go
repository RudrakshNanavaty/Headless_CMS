package utils

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect to DB

func ConnectDBByCredentials(host, user, password, port, name, sslMode string) (*gorm.DB, error) {
	if sslMode == "" {
		sslMode = "disable"
	}
	if user == "" || password == "" || port == "" {
		return nil, fmt.Errorf("DB credentials not provided")
	}

	var dsn string
	if name != "" {
		if host != "" {
			dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, name, port, sslMode)
		} else {
			dsn = fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", user, password, name, port)
		}
	} else {
		if host != "" {
			dsn = fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=%s", host, user, password, port, sslMode)
		} else {
			dsn = fmt.Sprintf("user=%s password=%s port=%s sslmode=disable", user, password, port)
		}
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database")
	}
	return db, nil
}

func ConnectDBByURI(connectionURI string) (*gorm.DB, error) {
	if connectionURI != "" {
		db, err := gorm.Open(postgres.Open(connectionURI), &gorm.Config{})
		if err != nil {
			return nil, fmt.Errorf("failed to connect to database")
		}
		return db, nil
	}
	return nil, fmt.Errorf("DB connection URI not provided")
}
