package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID    `bson:"id"`
	Firstname     *string               `json:"firstname"`
	Lastname      *string               `json:"lastname"`
	Phone         *string               `json:"phone"`
	Email         *string               `json:"email"`
	OTP           *string               `json:"otp"`
	OTP_expires   time.Time             `json:"otp_expires"`
    Created_At    time.Time             `json:"created_at"`
	Updated_At    time.Time             `json:"updated_at"`
	User_ID       string                `json:"user_id"`

}