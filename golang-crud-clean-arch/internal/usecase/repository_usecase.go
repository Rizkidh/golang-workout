package usecase

import (
	"context"
	"golang-crud-clean-arch/m/internal/entity"
	"golang-crud-clean-arch/m/internal/repository"
)

type RepositoryUsecase struct {
	repoRepo *repository.RepoRepository // ✅ BENAR
}

func NewRepositoryUsecase(repoRepo *repository.RepoRepository) *RepositoryUsecase { // ✅ BENAR
	return &RepositoryUsecase{repoRepo}
}

func (u *RepositoryUsecase) CreateRepository(ctx context.Context, repo *entity.Repository) error {
	return u.repoRepo.Create(ctx, repo)
}

func (u *RepositoryUsecase) GetAllRepositories(ctx context.Context) ([]entity.Repository, error) {
	return u.repoRepo.GetAll(ctx)
}

func (u *RepositoryUsecase) GetRepository(ctx context.Context, id int) (*entity.Repository, error) {
	return u.repoRepo.GetByID(ctx, id)
}

func (u *RepositoryUsecase) UpdateRepository(ctx context.Context, repo *entity.Repository) error {
	return u.repoRepo.Update(ctx, repo)
}

func (u *RepositoryUsecase) DeleteRepository(ctx context.Context, id int) error {
	return u.repoRepo.Delete(ctx, id)
}
