package usecase

import (
	"context"
	"golang-crud-clean-arch/internal/entity"
	"golang-crud-clean-arch/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RepositoryUsecase struct {
	repoRepo *repository.RepoRepository
}

func NewRepositoryUsecase(repoRepo *repository.RepoRepository) *RepositoryUsecase {
	return &RepositoryUsecase{repoRepo}
}

func (u *RepositoryUsecase) CreateRepository(ctx context.Context, repo *entity.Repository) error {
	repo.ID = primitive.NewObjectID()
	return u.repoRepo.Create(ctx, repo)
}

func (u *RepositoryUsecase) GetAllRepositories(ctx context.Context) ([]entity.Repository, error) {
	return u.repoRepo.GetAllRepositories(ctx)
}

func (u *RepositoryUsecase) GetRepository(ctx context.Context, id primitive.ObjectID) (*entity.Repository, error) {
	return u.repoRepo.GetByID(ctx, id)
}

func (u *RepositoryUsecase) UpdateRepository(ctx context.Context, repo *entity.Repository) error {
	return u.repoRepo.Update(ctx, repo)
}

func (u *RepositoryUsecase) DeleteRepository(ctx context.Context, id primitive.ObjectID) error {
	return u.repoRepo.Delete(ctx, id)
}
