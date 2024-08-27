package repository

import (
	"context"
	"fmt"

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

func (op *OtpRepository) GetOtpByEmail(ctx context.Context, email string) (*domain.OTP, error) {
    filter := bson.M{"email": email}
    var otp domain.OTP
    err := op.collection.FindOne(ctx, filter).Decode(&otp)
    
    if err == mongo.ErrNoDocuments {
        fmt.Println("No OTP found for email:", email)
        return nil, nil // No OTP found
    } else if err != nil {
        fmt.Println("Error finding OTP:", err)
        return nil, err // Some other error occurred
    }

    fmt.Println("Found OTP:", otp)
    return &otp, nil
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
