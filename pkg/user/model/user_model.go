package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string `json:"username"`
	Password     string `json:"password"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Telephone    string `json:"telephone"`
	ProfileImage string `json:"profile_image"`
	Address      string `json:"address"`
	City         string `json:"city"`
	Province     string `json:"province"`
	Country      string `json:"country"`
}
