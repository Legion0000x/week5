package entity

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	PhoneNumber string `gorm:"uniqueIndex" valid:"required~PhoneNumber is required, stringlength(10|10)~Phone Number length is not 10 digits."`
	Password    string
	Url         string `valid:"required~Url is required, url~Url is invalid"`
}
