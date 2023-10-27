package models

import "time"

type User struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	NIK       uint       `json:"nik" gorm:"unique"`
	Nama      string     `json:"nama"`
	Loker 	  []Loker	 `json:"loker" gorm:"constraint:OnDelete:CASCADE"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}