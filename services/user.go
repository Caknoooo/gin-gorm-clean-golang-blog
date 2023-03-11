package services

import (
	"context"

	"github.com/Caknoooo/Golang-BLOG/dto"
	"github.com/Caknoooo/Golang-BLOG/entities"
	"github.com/Caknoooo/Golang-BLOG/helpers"
	"github.com/Caknoooo/Golang-BLOG/repository"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type UserService interface {
	RegisterUser(ctx context.Context, userDTO dto.UserCreateDTO) (entities.User, error)
	GetAllUser(ctx context.Context) ([]entities.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
	CheckUser(ctx context.Context, email string) (bool, error)
	UpdateUser(ctx context.Context, userDTO dto.UserUpdateDTO) (error)
	DeleteUser(ctx context.Context, userID uuid.UUID) (error)
	Verify(ctx context.Context, email string, password string) (bool, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService{
	return &userService{
		userRepository: ur,
	}
}

func (us *userService) RegisterUser(ctx context.Context, userDTO dto.UserCreateDTO) (entities.User, error){
	user := entities.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(userDTO))
	user.Role = "user"
	if err != nil {
		return user, err
	}
	return us.userRepository.RegisterUser(ctx, user)
} 

func (us *userService) GetAllUser(ctx context.Context) ([]entities.User, error){
	return us.userRepository.GetAllUser(ctx)
}

func (us *userService) GetUserByID(ctx context.Context, userID uuid.UUID) (entities.User, error){
	return us.userRepository.GetUserByID(ctx, userID);
}

func (us *userService) GetUserByEmail(ctx context.Context, email string) (entities.User, error){
	return us.userRepository.GetUserByEmail(ctx, email)
}

func (us *userService) CheckUser(ctx context.Context, email string) (bool, error){
	res, err := us.userRepository.GetUserByEmail(ctx, email)
	if err != nil{
		return false, err
	}

	if res.Email == ""{
		return false, nil
	}
	return true, nil
}

func (us *userService) UpdateUser(ctx context.Context, userDTO dto.UserUpdateDTO) (error){
	user := entities.User{}
	if err := smapping.FillStruct(&user, smapping.MapFields(userDTO)); err != nil {
		return nil
	}
	return us.userRepository.UpdateUser(ctx, user)
}

func (us *userService) DeleteUser(ctx context.Context, userID uuid.UUID) (error){
	return us.userRepository.DeleteUser(ctx, userID)
}

func (us *userService) Verify(ctx context.Context, email string, password string) (bool, error){
	res, err := us.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	CheckPassword, err := helpers.CheckPassword(res.Password, []byte(password))
	if err != nil {
		return false, err
	}

	if res.Email == email && CheckPassword{
		return true, nil
	}
	return false, nil
}