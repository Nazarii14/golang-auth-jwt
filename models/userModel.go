package models

import (
	_ "go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           int       `bson:"_id"`
	Firstname    *string   `json:"firstname" validate:"required" minLength:"2" maxLength:"100"`
	Lastname     *string   `json:"lastname" validate:"required" minLength:"2" maxLength:"100"`
	Username     *string   `json:"username" validate:"required" minLength:"2" maxLength:"100"`
	Email        *string   `json:"email" validate:"email, required" minLength:"2" maxLength:"100"`
	Phone        *string   `json:"phone"`
	Password     *string   `json:"password"`
	Token        *string   `json:"token"`
	UserType     *string   `json:"userType" validate:"required, eq=ADMIN|USER, default=USER"`
	RefreshToken *string   `json:"refreshToken"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
