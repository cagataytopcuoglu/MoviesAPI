package user

import (
	"MovieAPI/pkg/mongoHelper"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetUser(entity *Login) (*User, error)
}

type repository struct {
	collection *mongo.Collection
}

func (r *repository) GetUser(entity *Login) (user *User, err error) {
	query := mongoHelper.BuildQuery(map[string]string{"UserName": entity.UserName, "Password": entity.Password})
	r.collection.FindOne(context.TODO(), query).Decode(&user)
	return user, nil
}

func NewRepository(db *mongo.Database) Repository {

	col := db.Collection("users")
	return &repository{
		collection: col,
	}
}
