package repository

import (
	"context"
	"errors"
	"fmt"

	"golang-crud-clean-arch/internal/entity"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepoRepository struct {
	db       *mongo.Client
	redis    *redis.Client
	validate *validator.Validate
}

func NewRepoRepository(db *mongo.Client, redis *redis.Client) *RepoRepository {
	return &RepoRepository{
		db:       db,
		redis:    redis,
		validate: validator.New(),
	}
}

func (r *RepoRepository) Create(ctx context.Context, repo *entity.Repository) error {
	if err := repo.Validate(); err != nil {
		return errors.New("validasi gagal: " + err.Error())
	}
	collection := r.db.Database("gotrial").Collection("repo")
	repo.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(ctx, repo)
	if err == nil {
		r.redis.Del(ctx, "repositories:all")
	}
	return err
}

func (r *RepoRepository) GetAllRepositories(ctx context.Context) ([]entity.Repository, error) {
	collection := r.db.Database("gotrial").Collection("repo")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var repos []entity.Repository
	if err = cursor.All(ctx, &repos); err != nil {
		return nil, err
	}
	return repos, nil
}

func (r *RepoRepository) GetByID(ctx context.Context, id primitive.ObjectID) (*entity.Repository, error) {
	collection := r.db.Database("gotrial").Collection("repo")
	var repo entity.Repository
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&repo)
	if err != nil {
		return nil, err
	}
	return &repo, nil
}

func (r *RepoRepository) Update(ctx context.Context, repo *entity.Repository) error {
	if err := repo.Validate(); err != nil {
		return errors.New("validasi gagal: " + err.Error())
	}
	collection := r.db.Database("gotrial").Collection("repo")
	update := bson.M{"$set": bson.M{"name": repo.Name, "url": repo.URL, "ai_enabled": repo.AIEnabled, "user_id": repo.UserID}}
	result, err := collection.UpdateOne(ctx, bson.M{"_id": repo.ID}, update)
	if err == nil && result.MatchedCount > 0 {
		r.redis.Del(ctx, "repositories:all", fmt.Sprintf("repositories:%s", repo.ID.Hex()))
	}
	return err
}

func (r *RepoRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	collection := r.db.Database("gotrial").Collection("repo")
	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err == nil && result.DeletedCount > 0 {
		r.redis.Del(ctx, "repositories:all", fmt.Sprintf("repositories:%s", id.Hex()))
	}
	return err
}
