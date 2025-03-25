package repository

import (
	"context"
	"errors"
	"golang-crud-clean-arch/internal/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	db *mongo.Client
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (any, any) {
	panic("unimplemented")
}

func NewUserRepository(db *mongo.Client) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	collection := r.db.Database("gotrial").Collection("users")

	// Mencari ID terakhir untuk auto-increment
	var lastUser entity.User
	opts := options.FindOne().SetSort(bson.M{"id": -1})
	err := collection.FindOne(ctx, bson.M{}, opts).Decode(&lastUser)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	// Set ID baru
	if err == mongo.ErrNoDocuments {
		user.ID = 1
	} else {
		user.ID = lastUser.ID + 1
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*entity.User, error) {
	collection := r.db.Database("gotrial").Collection("users")
	var user entity.User
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&user)
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

	result, err := collection.UpdateOne(ctx, bson.M{"id": user.ID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	collection := r.db.Database("gotrial").Collection("users")

	result, err := collection.DeleteOne(ctx, bson.M{"id": id})
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
