package controller

import (
	"net/http"

	"github.com/Afomiat/E-Commerce-API---Mini-PRD/config"
	"github.com/Afomiat/E-Commerce-API---Mini-PRD/domain"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	SignupUsecase    domain.SignupUsecase
	env                 *config.Env

}

func NewSignupController(SignupUsecase domain.SignupUsecase, env *config.Env ) *SignupController{
	return &SignupController{
		SignupUsecase: SignupUsecase,
		env:           env,
	}
}

func (sc *SignupController) Signup(ctx *gin.Context) {
    var user domain.SignupForm

    if err := ctx.ShouldBindJSON(&user); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Check if the username or email already exists
    returnUser, _ := sc.SignupUsecase.GetUserByUserName(ctx, user.Username)
    if returnUser != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username Already Exists!"})
        return
    }

    returnUser, _ = sc.SignupUsecase.GetUserByEmail(ctx, user.Email)
    if returnUser != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email Already Exists!"})
        return
    }

    // Send OTP
    err := sc.SignupUsecase.SendOtp(ctx, &user, sc.env.SMTPUsername, sc.env.SMTPPassword)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "OTP sending failed"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"message": "OTP Sent"})
}
