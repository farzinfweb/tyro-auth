package repo

import (
	"authn/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepo interface {
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
}

type mongoUserRepo struct {
	db *mongo.Database
}

func NewMongoUserRepo(db *mongo.Database) IUserRepo {
	return &mongoUserRepo{
		db,
	}
}

func (m *mongoUserRepo) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user *domain.User

	err := m.db.Collection("users").FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
