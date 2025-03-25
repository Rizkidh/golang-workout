package repository

import (
	"context"
	"golang-crud-clean-arch/m/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepoRepository struct {
	db *mongo.Client
}

func NewRepoRepository(db *mongo.Client) *RepoRepository {
	return &RepoRepository{db}
}

func (r *RepoRepository) Create(ctx context.Context, repo *entity.Repository) error {
	collection := r.db.Database("gotrial").Collection("repositories")
	_, err := collection.InsertOne(ctx, repo)
	return err
}

func (r *RepoRepository) GetAll(ctx context.Context) ([]entity.Repository, error) {
	collection := r.db.Database("gotrial").Collection("repositories")
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

func (r *RepoRepository) GetByID(ctx context.Context, id int) (*entity.Repository, error) {
	collection := r.db.Database("gotrial").Collection("repositories")
	var repo entity.Repository
	err := collection.FindOne(ctx, bson.M{"id": id}).Decode(&repo)
	return &repo, err
}

func (r *RepoRepository) Update(ctx context.Context, repo *entity.Repository) error {
	collection := r.db.Database("gotrial").Collection("repositories")
	_, err := collection.UpdateOne(ctx, bson.M{"id": repo.ID}, bson.M{
		"$set": bson.M{
			"name":       repo.Name,
			"url":        repo.URL,
			"ai_enabled": repo.AIEnabled,
		},
	})
	return err
}

func (r *RepoRepository) Delete(ctx context.Context, id int) error {
	collection := r.db.Database("gotrial").Collection("repositories")
	_, err := collection.DeleteOne(ctx, bson.M{"id": id})
	return err
}
