package usecase

import (
	"context"
	"errors"
	"net/smtp"
	"time"

	"github.com/Afomiat/E-Commerce-API---Mini-PRD/config"
	"github.com/Afomiat/E-Commerce-API---Mini-PRD/domain"
	"github.com/Afomiat/E-Commerce-API---Mini-PRD/internal/userutil"
)

type SignupUsecase struct {
	signupRepo     domain.SignupRepository
	contextTimeout time.Duration
	tokenRepo      domain.TokenRepository
	otpRepo        domain.OtpRepository
	env            *config.Env
}



func NewSignupUsecase(signupRepo domain.SignupRepository, tokenRepo domain.TokenRepository, otpRepo domain.OtpRepository, timeout time.Duration, env *config.Env) *SignupUsecase {
	return &SignupUsecase{
		signupRepo:     signupRepo,
		contextTimeout: timeout,
		tokenRepo:      tokenRepo,
		otpRepo:        otpRepo,
		env:            env,
	}
}

func (su *SignupUsecase) GetUserByUserName(ctx context.Context, username string) (*domain.SignupForm, error) {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()

	user, err := su.signupRepo.GetUserByUserName(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (su *SignupUsecase) GetUserByEmail(ctx context.Context, Email string) (*domain.SignupForm, error) {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()

	user, err := su.signupRepo.GetUserByEmail(ctx, Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (su *SignupUsecase) SendOtp(ctx context.Context, user *domain.SignupForm, stmpName string, stmpPass string) error {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()

	if !userutil.ValidateEmail(user.Email) {
		return errors.New("invalid email")
	}
	if !userutil.ValidatePassword(user.Password) {
		return errors.New("password must be at least 8 characters long")
	}

	storedOtp, _ := su.GetOtpByEmail(ctx, user.Email)
	if storedOtp != nil {
		return errors.New("Otp already sent")
	}

	if storedOtp != nil {
		if time.Now().Before(storedOtp.ExpiresAt) {
			return errors.New("OTP already sent")
		}

		errChan := make(chan error, 1)
		go func() {
			errChan <- su.otpRepo.DeleteOTP(ctx, storedOtp.Email)
		}()

		otp := domain.OTP{
			Value:     userutil.GenerateOTP(),
			Username:  user.Username,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: time.Now(),
			ExpiresAt: time.Now().Add(time.Minute * 5),
		}

		if err := su.SaveOTP(ctx, &otp); err != nil {
			return err
		}

		if err := <-errChan; err != nil {
			return err
		}

		if err := su.SendEmail(user.Email, otp.Value, stmpName, stmpPass); err != nil {
			return err
		}
	}
	return nil
}

func (su *SignupUsecase) GetOtpByEmail(ctx context.Context, email string) (*domain.OTP, error) {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	return su.otpRepo.GetOTPByEmail(ctx, email)
}

func (su *SignupUsecase) SaveOTP(ctx context.Context, otp *domain.OTP) error {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	return su.otpRepo.SaveOTP(ctx, otp)
}

func (su *SignupUsecase) SendEmail(email string, otpValue, smtpusername string, smtppassword string) error {
	errChan := make(chan error, 1)
	go func() {
		from := smtpusername
		password := smtppassword

		to := []string{email}

		smtpHost := su.env.SMTPHost
		smtpPort := su.env.SMTPPort

		message := []byte("Your OTP is " + otpValue)

		auth := smtp.PlainAuth("", from, password, smtpHost)

		err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
		if err != nil {

			// use a channel to signal the error
			errChan <- err
			return
		}
	}()
	return nil
}
