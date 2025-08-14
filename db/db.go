package db

import (
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	OrderUID string         `gorm:"primaryKey"`
	Data     datatypes.JSON `gorm:"type:jsonb"`
}

var DB *gorm.DB

func Init(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
	err = DB.AutoMigrate(&Order{})
	if err != nil {
		panic(err)
	}
}

func InsertOrderWithTx(tx *gorm.DB, order *Order) error {
	if err := tx.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrder(uid string) (*Order, error) {
	var order Order
	err := DB.First(&order, "order_uid = ?", uid).Error
	fmt.Println(err)
	return &order, err
}

func GetAllOrders() ([]Order, error) {
	var orders []Order
	err := DB.Find(&orders).Error
	return orders, err
}

func BeginTransaction() (*gorm.DB, error) {
	tx := DB.Begin()
	if tx.Error != nil {
		return nil, errors.New("failed to begin transaction")
	}
	return tx, nil
}

func CommitTransaction(tx *gorm.DB) {
	tx.Commit()
}

func RollbackTransaction(tx *gorm.DB) {
	tx.Rollback()
}

func (o *Order) UnmarshalData(dest interface{}) error {
	return json.Unmarshal(o.Data, dest)
}
