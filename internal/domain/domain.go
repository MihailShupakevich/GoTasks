package domain

import "time"

type Task struct {
	ID        uint      `gorm:"primaryKey"`
	Text      string    `gorm:"type:text" json:"text"`
	Checkbox  bool      `gorm:"type:boolean" json:"checkbox"`
	CreatedAt time.Time `gorm:"index default current_timestamp()"`
	UpdatedAt time.Time `gorm:"index default current_timestamp()"`
}
