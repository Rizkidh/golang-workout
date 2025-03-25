package repository

import (
	"context"
	"errors"
	"golang-crud-clean-arch/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Client
}

func NewUserRepository(db *mongo.Client) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	collection := r.db.Database("gotrial").Collection("users")
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*entity.User, error) {
	collection := r.db.Database("gotrial").Collection("users")
	var user entity.User
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	collection := r.db.Database("gotrial").Collection("users")

	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"updated_at": time.Now(),
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": user.ID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	collection := r.db.Database("gotrial").Collection("users")

	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	collection := r.db.Database("gotrial").Collection("users")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []entity.User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}
