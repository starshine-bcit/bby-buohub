package service

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/starshine-bcit/bby-buohub/auth/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToMariaDB(cfg *util.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"unique;index;not null"`
	Password  string    `gorm:"not null"`
	Salt      string    `gorm:"not null"`
	Created   time.Time `gorm:"not null"`
	LastLogin sql.NullTime
}

func Migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	if err := db.AutoMigrate(&User{}); err != nil {
		util.ErrorLogger.Fatalf("Could not automigrate db. Err: %v\n", err.Error())
	}
}

func ValidatePassword(db *gorm.DB, username, password string) bool {
	user := &User{}
	tx := db.Where("username = ?", username).First(user)
	if tx.Error != nil {
		util.InfoLogger.Printf("Could not retrieve user from db. Err: %v\n", tx.Error.Error())
		return false
	}
	ok := CompareHashFromDB(user.Password, user.Salt, password)
	if !ok {
		return false
	}
	user.LastLogin.Time = time.Now()
	user.LastLogin.Valid = true
	result := db.Save(user)
	if result.Error != nil {
		util.ErrorLogger.Printf("Error updating user in database. err: %v\n", result.Error.Error())
	}
	return true
}

func CreateUser(db *gorm.DB, username, password string) bool {
	hash, salt, err := GenerateHashForDB(password)
	if err != nil {
		util.ErrorLogger.Printf("Error creating user in database. err: %v\n", err.Error())
	}
	user := &User{
		Username:  username,
		Password:  hash,
		Salt:      salt,
		Created:   time.Now(),
		LastLogin: time.Now(),
	}
	result := db.Create(user)
	if result.Error != nil {
		util.WarningLogger.Printf("Could not create user in database. err: %v\n", result.Error.Error())
		return false
	}
	return true
}

func CheckUsername(db *gorm.DB, username string) bool {
	user := &User{}
	tx := db.Where("username = ?", username).First(user)
	if tx.Error != nil {
		return false
	}
	if user.Username == "" {
		return false
	}
	return true
}
