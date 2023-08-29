package domain

import "time"

type User struct {
    ID        uint   `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Username  string
    Password  string
}