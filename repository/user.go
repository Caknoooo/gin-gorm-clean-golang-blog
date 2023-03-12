package repository

import (
	"context"

	"github.com/Caknoooo/Golang-BLOG/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUser(ctx context.Context) ([]entities.User, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (entities.User, error)
	RegisterUser(ctx context.Context, user entities.User) (entities.User, error)
	UpdateUser(ctx context.Context, user entities.User) (error)
	DeleteUser(ctx context.Context, userID uuid.UUID) (error)
}

type userConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository{
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) RegisterUser(ctx context.Context, user entities.User) (entities.User, error){
	if tx := db.connection.Create(&user).Error; tx != nil {
		return entities.User{}, tx 
	}
	return user, nil
}

func (db *userConnection) GetAllUser(ctx context.Context) ([]entities.User, error){
	var listUser []entities.User
	if tx := db.connection.Find(&listUser).Error; tx != nil {
		return nil, tx
	}
	return listUser, nil
}

func (db *userConnection) GetUserByID(ctx context.Context, userID uuid.UUID) (entities.User, error){
	var user entities.User
	if tx := db.connection.Where("id = ?", userID).Take(&user).Error; tx != nil {
		return user, tx
	}
	return user, nil
}

func (db *userConnection) GetUserByEmail(ctx context.Context, email string) (entities.User, error){
	var user entities.User
	if tx := db.connection.Where("email = ?", email).Take(&user).Error; tx != nil {
		return entities.User{}, tx
	}
	return user, nil
}

func (db *userConnection) UpdateUser(ctx context.Context, user entities.User) (error){
	if tx := db.connection.Updates(&user).Error; tx != nil {
		return tx
	}
	return nil
}

func (db *userConnection) DeleteUser(ctx context.Context, userID uuid.UUID) (error){
	if tx := db.connection.Delete(&entities.User{}, &userID).Error; tx != nil {
		return tx
	}
	return nil
}