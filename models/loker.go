package models

import "time"

type Loker struct {
	ID         uint        `json:"id" gorm:"primaryKey"`
	UserId	   uint        `json:"user_id"`
	NoLoker    uint        `json:"no_loker"`
	Saldo      uint        `json:"saldo"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}
