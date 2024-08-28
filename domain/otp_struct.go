package domain

import "time"

const (
	OtpCollection = "otp_new"
)

type OTP struct {
	Value     string    `bson:"value"`
	Username  string    `bson:"username"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
	ExpiresAt time.Time `bson:"expires_at"`
}

type VerifyOtp struct {
	Value string `bson:"value"`
	Email string `bson:"email"`
}
