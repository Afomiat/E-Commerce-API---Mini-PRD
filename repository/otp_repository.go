package repository

import (
	"context"

	"github.com/Afomiat/E-Commerce-API---Mini-PRD/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OtpRepository struct {
	collection *mongo.Collection
}

func NewOtpRepository(db *mongo.Database, coll string) *OtpRepository{
	return &OtpRepository{
		collection: db.Collection(coll),
	}

}

func (op *OtpRepository) GetOTPByEmail(ctx context.Context, email string) (*domain.OTP, error){
	
	filter := bson.M{"email": email}
	var otp domain.OTP
	err := op.collection.FindOne(ctx, filter).Decode(&otp)
	return &otp, err
}

func (op *OtpRepository) DeleteOTP(ctx context.Context, email string) error{
	filter := bson.M{"email": email}
	_, err := op.collection.DeleteOne(ctx,filter)
	return err
}

func (op *OtpRepository) SaveOTP(ctx context.Context, otp *domain.OTP) error{
	
	_, err := op.collection.InsertOne(ctx,otp)
	return err
}
