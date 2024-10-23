package db

import (
	"fmt"
	"strconv"
	"time"

	"github.com/haisabdillah/golang-auth/core/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitMySql() (*gorm.DB, error) {

	cfg := config.LoadConfig()

	dbHost := cfg.Database.Host
	dbPort := cfg.Database.Port
	dbUsername := cfg.Database.User
	dbPassword := cfg.Database.Password
	dbName := cfg.Database.DBName
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbName)
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error), // Change to Silent, Error, Warn, or Info
	})
	if err != nil {
		return nil, err

	}
	// Mendapatkan objek *sql.DB dari GORM
	sqlDB, err := DB.DB()
	if err != nil {
		return nil, err

	}

	dbMaxConn, _ := strconv.Atoi(cfg.Database.DBPoolMax)
	dbIddleConn, _ := strconv.Atoi(cfg.Database.DBPoolIddle)
	dbMaxConnTime, _ := strconv.Atoi(cfg.Database.DBPoolMaxTime)
	// Mengatur parameter connection pooling
	sqlDB.SetMaxOpenConns(dbMaxConn)                                     // Jumlah maksimum koneksi yang dapat dibuka
	sqlDB.SetMaxIdleConns(dbIddleConn)                                   // Jumlah maksimum koneksi idle (menganggur) dalam pool
	sqlDB.SetConnMaxLifetime(time.Duration(dbMaxConnTime) * time.Minute) // Batas waktu hidup maksimum untuk sebuah koneksi

	// Test the connection
	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Sucess Connect Database")
	return DB, nil

}
