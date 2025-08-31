package database

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type OrderV1 struct {
	OrderUID string         `gorm:"primaryKey"`
	Data     datatypes.JSON `gorm:"type:jsonb"`
}

func (OrderV1) TableName() string {
	return "orders"
}

type OrderV2 struct {
	OrderUID  string         `gorm:"primaryKey"`
	Data      datatypes.JSON `gorm:"type:jsonb"`
	Timestamp time.Time
}

func (OrderV2) TableName() string {
	return "orders"
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&OrderV1{}); err != nil {
		return err
	}

	if !db.Migrator().HasColumn(&OrderV2{}, "timestamp") {
		if err := db.Migrator().AddColumn(&OrderV2{}, "Timestamp"); err != nil {
			return err
		}
	}

	return nil
}
