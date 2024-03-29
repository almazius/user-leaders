package internal

import (
	"time"
)

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	IsAdmin    bool
	Password   string
}

type HackError struct {
	Code      int
	Err       error
	Message   string
	Timestamp time.Time
}

type IAdmin interface {
	ConfirmPlace(placeId int) (int, error)
	DeletePlace(placeId int) error
}

type IUser interface {
	CreatePlace() (int, error)
	RentPlace(placeId int) error
}
