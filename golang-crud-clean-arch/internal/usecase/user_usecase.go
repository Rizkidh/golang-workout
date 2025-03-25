package usecase

import (
	"context"
	"golang-crud-clean-arch/internal/entity"
	"golang-crud-clean-arch/internal/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase struct {
	userRepo *repository.UserRepository // ✅ BENAR
}

func (u *UserUsecase) GetAllUsers(ctx context.Context) ([]entity.User, error) {
	return u.userRepo.GetAll(ctx)
}

func (u *UserUsecase) UpdateUser(ctx context.Context, user *entity.User) error {
	return u.userRepo.Update(ctx, user)
}

func (u *UserUsecase) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	return u.userRepo.Delete(ctx, id)
}

func NewUserUsecase(userRepo *repository.UserRepository) *UserUsecase { // ✅ BENAR
	return &UserUsecase{userRepo}
}

func (u *UserUsecase) CreateUser(ctx context.Context, user *entity.User) error {
	return u.userRepo.Create(ctx, user)
}

func (u *UserUsecase) GetUser(ctx context.Context, id primitive.ObjectID) (*entity.User, error) {
	return u.userRepo.GetByID(ctx, id)
}
