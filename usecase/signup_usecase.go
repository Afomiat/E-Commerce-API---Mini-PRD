package usecase

import (
	"context"
	"errors"

	"net/smtp"
	"time"

	"github.com/Afomiat/E-Commerce-API---Mini-PRD/config"
	"github.com/Afomiat/E-Commerce-API---Mini-PRD/domain"
	"github.com/Afomiat/E-Commerce-API---Mini-PRD/internal/userutil"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignupUsecase struct {
	signupRepo     domain.SignupRepository
	contextTimeout time.Duration
	otpRepo        domain.OtpRepository
	env            *config.Env
}

// SendOtp implements domain.SignupUsecase.
// func (su *SignupUsecase) SendOtp(cxt context.Context, user *domain.SignupForm, stmpName string, stmpPass string) error {
// 	panic("unimplemented")
// }

func NewSignupUsecase(signupRepo domain.SignupRepository, otpRepo domain.OtpRepository, timeout time.Duration, env *config.Env) *SignupUsecase {
	return &SignupUsecase{
		signupRepo:     signupRepo,
		contextTimeout: timeout,
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

func (su *SignupUsecase) SendOtp(c context.Context, user *domain.SignupForm, smtpusername, smtppassword string) error {
    // Retrieve existing OTP if any
    storedOTP, err := su.GetOtpByEmail(c, user.Email)
    if err != nil && err != mongo.ErrNoDocuments {
        return err
    }

    if storedOTP != nil {
        if time.Now().Before(storedOTP.ExpiresAt) {
            return errors.New("OTP already sent")
        }

        // Delete the expired OTP
        if err := su.otpRepo.DeleteOTP(c, storedOTP.Email); err != nil {
            return err
        }
    }

    // Generate a new OTP
    otp := domain.OTP{
        Value:     userutil.GenerateOTP(),
        Username:  user.Username,
        Email:     user.Email,
        Password:  user.Password,
        CreatedAt: time.Now(),
        ExpiresAt: time.Now().Add(time.Minute * 5),
    }

    // Save the new OTP
    if err := su.otpRepo.SaveOTP(c, &otp); err != nil {
        return err
    }

    // Send the OTP via email
    if err := su.SendEmail(user.Email, otp.Value, smtpusername, smtppassword); err != nil {
        return err
    }

    return nil
}


func (su *SignupUsecase) SendEmail(email string, otpValue, smtpusername, smtppassword string) error {
    from := smtpusername
    password := smtppassword

    to := []string{email}
    smtpHost := "smtp.gmail.com"
    smtpPort := "587"
    message := []byte("Your OTP is " + otpValue)

    auth := smtp.PlainAuth("", from, password, smtpHost)
    return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
}


func (su *SignupUsecase) GetOtpByEmail(ctx context.Context, email string) (*domain.OTP, error) {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	
	return su.otpRepo.GetOtpByEmail(ctx, email)
}

func (su *SignupUsecase) SaveOTP(ctx context.Context, otp *domain.OTP) error {
	ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()
	return su.otpRepo.SaveOTP(ctx, otp)
}

func (su *SignupUsecase) VerifyOtp(ctx context.Context, otp *domain.VerifyOtp)(*domain.OTP, error){
    ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
	defer cancel()

    storedOTP, err := su.GetOtpByEmail(ctx, otp.Email)

    if err != nil{
        return nil, errors.New("OTP not found Please signup again")
    }

    if storedOTP.Value != otp.Value{
        return nil, errors.New("invslid OTP")
    }

    if time.Now().After(storedOTP.ExpiresAt) {
		return nil, errors.New("OTP expired")
	}

    err = su.otpRepo.DeleteOTP(ctx, storedOTP.Email)
    if err != nil{
        return nil, err
    }

    return storedOTP, nil
}
func (su *SignupUsecase) RegisterUser(ctx context.Context, user *domain.SignupForm)(*primitive.ObjectID, error){
    ctx, cancel := context.WithTimeout(ctx, su.contextTimeout)
    defer cancel()

    hashedPass, err := userutil.HassPassword(user.Password)

    if err != nil{
        return nil, err
    }

    addUser := domain.SignupForm{
        ID: primitive.NewObjectID(),
        Username: user.Username,
        Password: hashedPass,
        Email: user.Email,
        Role: "user",
    }
    err = su.signupRepo.CreateUser(ctx, &addUser)
    
    return &addUser.ID, err
}