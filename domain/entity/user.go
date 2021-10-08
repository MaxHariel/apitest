package entity

import "time"

type User struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id" example:"1"`
	Name      string     `gorm:"size:100;not null;unique" json:"name" example:"Max"`
	CreatedAt time.Time  `json:"created_at" example:"2021-10-06T13:47:02.127645837-03:00"`
	UpdatedAt time.Time  `json:"updated_at" example:"2021-10-06T13:47:02.127645837-03:00"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type UserDTO struct {
	Name string `json:"name" example:"Max"`
}
