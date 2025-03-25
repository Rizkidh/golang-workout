package usecase

import (
	"context"
	"golang-crud-clean-arch/m/internal/entity"
	"golang-crud-clean-arch/m/internal/repository"
)

type UserUsecase struct {
	userRepo *repository.UserRepository // ✅ BENAR
}

func (u *UserUsecase) UpdateUser(ctx context.Context, user *entity.User) error {
	return u.userRepo.Update(ctx, user)
}

func (u *UserUsecase) DeleteUser(ctx context.Context, id int) error {
	return u.userRepo.Delete(ctx, id)
}

func NewUserUsecase(userRepo *repository.UserRepository) *UserUsecase { // ✅ BENAR
	return &UserUsecase{userRepo}
}

func (u *UserUsecase) CreateUser(ctx context.Context, user *entity.User) error {
	return u.userRepo.Create(ctx, user)
}

func (u *UserUsecase) GetUser(ctx context.Context, id int) (*entity.User, error) {
	return u.userRepo.GetByID(ctx, id)
}
