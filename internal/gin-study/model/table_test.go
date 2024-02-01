package model

import "time"

type TableTest struct {
	ID                int64     `gorm:"column:id" json:"id"`
	OrderID           int64     `gorm:"column:order_id" json:"order_id"`
	MileageNumber     int64     `gorm:"column:mileage_number" json:"mileage_number"`
	OperationType     string    `gorm:"column:operation_type" json:"operation_type"`
	OperationUser     string    `gorm:"column:operation_user" json:"operation_user"`
	OperationNumber   int64     `gorm:"column:operation_number" json:"operation_number"`
	SendStatus        string    `gorm:"column:send_status" json:"send_status"`
	OperationDatetime time.Time `gorm:"column:operation_datetime" json:"operation_datetime"`
	Status            string    `gorm:"column:status" json:"status"`
	Created           time.Time `gorm:"column:created" json:"created"`
	Updated           time.Time `gorm:"column:updated" json:"updated"`
}

// TableName the name of table in database
func (t *TableTest) TableName() string {
	return "table_test"
}
