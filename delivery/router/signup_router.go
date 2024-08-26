package router

import (
	"time"

	"github.com/Afomiat/E-Commerce-API---Mini-PRD/config"
	"github.com/Afomiat/E-Commerce-API---Mini-PRD/delivery/controller"
	"github.com/Afomiat/E-Commerce-API---Mini-PRD/domain"
	"github.com/Afomiat/E-Commerce-API---Mini-PRD/repository"
	"github.com/Afomiat/E-Commerce-API---Mini-PRD/usecase"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewSignUpRouter(env *config.Env, timeout time.Duration, db *mongo.Database, PublicRout *gin.RouterGroup){
	userR := repository.NewSignupRepo(db, domain.UserCollection)
	tokenR := repository.NewTokenRepository(db, domain.TokenCollection)
	otpR := repository.NewOtpRepository(db, domain.OtpCollection)

	signUsecase := usecase.NewSignupUsecase(userR, tokenR, otpR, timeout, env)
	signController := controller.NewSignupController(signUsecase, env)

	PublicRout.POST("/signup", signController.Signup)

	
}