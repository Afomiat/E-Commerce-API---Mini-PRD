package repository

import (
	"context"

	"github.com/Afomiat/E-Commerce-API---Mini-PRD/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SignupReo struct {
	collection *mongo.Collection
	
}

func NewSignupRepo(db *mongo.Database, coll string) *SignupReo{
	return &SignupReo{
		collection: db.Collection(coll),
		
	}
}

func (sr *SignupReo) GetUserByUserName(ctx context.Context, username string)(*domain.SignupForm, error){
	
	var user domain.SignupForm

	filter := bson.D{{Key: "username", Value: username}}
	err := sr.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No user found
		}
		return nil, err // Other errors
	}
	return &user, nil
}

func (sr *SignupReo) GetUserByEmail(ctx context.Context, Email string)(*domain.SignupForm, error){
	
	var user domain.SignupForm

	filter := bson.D{{Key: "username", Value: Email}}
	err := sr.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No user found
		}
		return nil, err // Other errors
	}
	return &user, nil
}

func (sr *SignupReo) CreateUser(ctx context.Context, user *domain.SignupForm) error{
	_, err := sr.collection.InsertOne(ctx, user)

	return err
}
