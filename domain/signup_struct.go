package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role string

const (
	UserCollection      = "users"
	AdminRole      Role = "ADMIN"
	UserRole       Role = "USER"
)

type SignupForm struct {
	ID       primitive.ObjectID `json:"_ID,omitempty" bson:"_ID,omitempty"`
	Username string             `json:"username" bson:"username" validate:"required,min=3,max=50"`
	Password string             `json:"password" bson:"password" validate:"required,min=6,max=50"`
	Email    string             `json:"email" bson:"email" validate:"required"`
	Role     string             `json:"role" bson:"role"`
}
