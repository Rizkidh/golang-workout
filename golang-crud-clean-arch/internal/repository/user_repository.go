package repository

import (
	"context"
	"golang-crud-clean-arch/m/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Client
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	collection := r.db.Database("gotrial").Collection("users")
	_, err := collection.UpdateOne(ctx, bson.M{"id": user.ID}, bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"updated_at": time.Now(),
		},
	})
	return err
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	collection := r.db.Database("gotrial").Collection("users")
	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	return err
}

func NewUserRepository(db *mongo.Client) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	collection := r.db.Database("gotrial").Collection("users")
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*entity.User, error) {
	collection := r.db.Database("gotrial").Collection("users")
	var user entity.User
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&user)
	return &user, err
}
