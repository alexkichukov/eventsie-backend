package models

import (
	"github.com/kamva/mgm/v3"
)

const (
	UserRole  string = "user"
	AdminRole string = "admin"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	FirstName        string   `json:"firstName" bson:"firstName" validate:"required,alpha"`
	LastName         string   `json:"lastName" bson:"lastName" validate:"required,alpha"`
	Password         string   `json:"password" bson:"password" validate:"required,min=8"`
	Email            string   `json:"email" bson:"email" validate:"required,email"`
	Role             string   `json:"role" bson:"role" validate:"required,oneof=user admin"`
	FavouriteEvents  []string `json:"favouriteEvents" bson:"favouriteEvents"`
	AttendingEvents  []string `json:"attendingEvents" bson:"attendingEvents"`
}
