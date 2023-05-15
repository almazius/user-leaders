package usecase

import (
	"errors"
	"github.com/mod/internal"
	"github.com/mod/internal/auth"
	"github.com/mod/internal/auth/repository"
	"log"
	"time"
)

func SingIn() {

}
func SingUp() {

}

func CreateUser(user *auth.UserForRegister) internal.HackError {
	existPhone, err := repository.ExistsUser(user.Phone)
	if err != nil {
		log.Print(err)
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}
	existEmail, err := repository.ExistsUser(user.Email)
	if err != nil {
		log.Print(err)
		return internal.HackError{
			Code:      500,
			Err:       err,
			Timestamp: time.Now(),
		}
	}

	if existEmail || existPhone {
		log.Print("invaluable data")
		return internal.HackError{
			Code:      400,
			Err:       errors.New("invaluable data"),
			Message:   "the number or email is already taken",
			Timestamp: time.Now(),
		}
	}

	return repository.CreateUser(user)
}
