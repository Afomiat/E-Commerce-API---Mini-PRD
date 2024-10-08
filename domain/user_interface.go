package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupUsecase interface {
	// Signup(ctx context.Context, user *SignupForm) error
	RegisterUser(ctx context.Context, user *SignupForm) (*primitive.ObjectID, error)
	VerifyOtp(ctx context.Context, otp *VerifyOtp)(*OTP, error)
	GetUserByUserName(ctx context.Context, username string) (*SignupForm, error)
	GetUserByEmail(ctx context.Context, Email string) (*SignupForm, error)
	SendOtp(cxt context.Context, user *SignupForm,  stmpName, stmpPass string) error
	GetOtpByEmail(ctx context.Context, email string) (*OTP, error)
	SaveOTP(ctx context.Context, otp *OTP) error
	SendEmail(email string, otpValue, smtpusername, smtppassword string) error 

}

type SignupRepository interface {
	// (ctx, username)
	CreateUser(ctx context.Context, user *SignupForm) error
	GetUserByUserName(ctx context.Context, username string) (*SignupForm, error)
	GetUserByEmail(ctx context.Context, Email string) (*SignupForm, error)
}

type OtpRepository interface{
	GetOtpByEmail(ctx context.Context, email string) (*OTP, error)
	DeleteOTP(ctx context.Context, email string) error
	SaveOTP(ctx context.Context, otp *OTP) error
}