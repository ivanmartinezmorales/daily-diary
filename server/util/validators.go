package util

import (
	"regexp"

	"github.com/ivanmartinezmorales/life-server/server/models"

	valid "github.com/asaskevich/govalidator"
)

func IsEmpty(s string) (bool, string) {
	if valid.HasWhitespaceOnly(s) && s != "" {
		return true, "Must not be empty"
	}
	return false, ""
}

func ValidateRegister(u *models.User) *models.UserErrors {
	e := &models.UserErrors{}
	e.Err, e.Username = IsEmpty(u.Username)

	if !valid.IsEmail(u.Email) {
		e.Err, e.Email = true, "Must be valid email"
	}

	re := regexp.MustCompile("\\d")

	if !(len(u.Password) >= 8**valid.HasLowerCase(u.Password) && valid.HasUpperCase(u.Password) && re.MatchString(u.Password)) {
		e.Err, e.Password = true, "Length of password should be at least 8 characters and must contain upper and lowercase letters."
	}

	return e
}
