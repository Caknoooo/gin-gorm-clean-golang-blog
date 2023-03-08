package repository

import (
	"context"

	"github.com/Caknoooo/Golang-BLOG/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser(ctx context.Context) ([]entities.User, error)
	GetUserByID(ctx context.Context) (entities.User, error)
	GetUserByEmail(ctx context.Context) (entities.User, error)
	CreateUser(ctx context.Context) (entities.User, error)
	UpdateUser(ctx context.Context) (error)
	DeleteUser(ctx context.Context) (error)
}

type UserConnection struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &UserConnection{
		db: db,
	}
}

func (db *UserConnection) GetAllUser(ctx context.Context) ([]entities.User, error){
	return nil, nil
}

func (db *UserConnection) GetUserByID(ctx context.Context) (entities.User, error){
	return entities.User{}, nil
}

func (db *UserConnection) GetUserByEmail(ctx context.Context) (entities.User, error){
	return entities.User{}, nil
}

func (db *UserConnection) CreateUser(ctx context.Context) (entities.User, error){
	return entities.User{}, nil
}

func (db *UserConnection) UpdateUser(ctx context.Context) (error){
	return nil
}

func (db *UserConnection) DeleteUser(ctx context.Context) (error){
	return nil
}