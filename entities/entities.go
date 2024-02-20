// @/entities/user.go
package entities

import "gorm.io/gorm"

type user struct {
	gorm.Model
	username  string `json:"username"`
	userID    string `json:"id"`
	email     string `json:"email"`
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
}
