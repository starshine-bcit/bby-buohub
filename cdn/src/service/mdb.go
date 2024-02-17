package service

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ryanfowler/uuid"
	"github.com/starshine-bcit/bby-buohub/cdn/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Video struct {
	ID               uint      `gorm:"primaryKey"`
	UUID             uuid.UUID `gorm:"type:uuid;unique;not null;index"`
	UploadedBy       string    `gorm:"index;not null"`
	ProcessComplete  bool      `gorm:"not null;default:false;index"`
	UploadedAt       time.Time `gorm:"not null;index"`
	Title            string    `gorm:"index;not null"`
	Description      string    `gorm:"not null"`
	OriginalFilename sql.NullString
	PosterFilename   sql.NullString
	ManifestName     sql.NullString
	AudioCodec       sql.NullString
	VideoCodec       sql.NullString
	Height           sql.NullInt32
	Width            sql.NullInt32
	RunTime          sql.NullInt32
	ProcessedAt      sql.NullTime
	UpdatedAt        time.Time
}

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

func Migrate(db *gorm.DB) {
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	if err := db.AutoMigrate(&Video{}); err != nil {
		util.ErrorLogger.Fatalf("Could not automigrate db. Err: %v\n", err.Error())
	}
}

func GetByUUID(db *gorm.DB, uuid uuid.UUID) (*Video, error) {
	video := &Video{}
	tx := db.Where("uuid = ?", uuid).First(video)
	if tx.Error != nil {
		util.InfoLogger.Printf("Could not retrieve video from db. Err: %v\n", tx.Error.Error())
		return nil, tx.Error
	}
	return video, nil
}

func UpdateVideo(db *gorm.DB, v *Video) {
	tx := db.Save(v)
	if tx.Error != nil {
		util.ErrorLogger.Printf("Could not save updated video row in table. err: %v\n", tx.Error.Error())
	}
}
