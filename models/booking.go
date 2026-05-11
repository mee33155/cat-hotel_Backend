package models

import (
    "time"

    "gorm.io/gorm"
)

type Booking struct {
    gorm.Model
    CustomerName string    `json:"customer_name"`
    CatName      string    `json:"cat_name"`
    RoomID       uint      `json:"room_id"`
    Room         Room      `json:"room" gorm:"foreignKey:RoomID"`
    CheckIn      time.Time `json:"check_in"`
    CheckOut     time.Time `json:"check_out"`
    Status       string    `json:"status"` // "pending", "paid"
    SessionID    string    `json:"session_id"`
}
