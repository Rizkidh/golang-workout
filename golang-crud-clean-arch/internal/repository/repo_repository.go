package repository

import (
	"context"
	"errors"
	"golang-crud-clean-arch/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepoRepository struct {
	db *mongo.Client
}

func NewRepoRepository(db *mongo.Client) *RepoRepository {
	return &RepoRepository{db}
}

func (r *RepoRepository) Create(ctx context.Context, repo *entity.Repository) error {
	collection := r.db.Database("gotrial").Collection("repo")

	// Pastikan repo.ID dan repo.UserID adalah ObjectID
	repo.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(ctx, repo)
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
	if err := cursor.All(ctx, &repos); err != nil {
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
	collection := r.db.Database("gotrial").Collection("repo")

	update := bson.M{
		"$set": bson.M{
			"name":       repo.Name,
			"url":        repo.URL,
			"ai_enabled": repo.AIEnabled,
			"user_id":    repo.UserID, // Pastikan UserID tetap sebagai ObjectID
		},
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": repo.ID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("repository not found")
	}
	return nil
}

func (r *RepoRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	collection := r.db.Database("gotrial").Collection("repo")

	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("repository not found")
	}
	return nil
}
