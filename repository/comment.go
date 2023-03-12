package repository

import (
	"context"

	"github.com/Caknoooo/Golang-BLOG/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment entities.Comment) (entities.Comment, error)
	GetAllComment(ctx context.Context) ([]entities.Comment, error)
	GetCommentsByBlogID(ctx context.Context, blogID uuid.UUID) ([]entities.Blog, error)
	UpdateComment(ctx context.Context, comment entities.Comment) error
	DeleteComment(ctx context.Context, userID uuid.UUID) error
}

type commentRepository struct {
	connection     *gorm.DB
	userRepository UserRepository
}

func NewCommentRepository(db *gorm.DB, ur UserRepository) CommentRepository {
	return &commentRepository{
		connection:     db,
		userRepository: ur,
	}
}

func (db *commentRepository) CreateComment(ctx context.Context, comment entities.Comment) (entities.Comment, error) {
	if tx := db.connection.Create(&comment).Error; tx != nil {
		return entities.Comment{}, nil
	}
	// user, err := db.userRepository.GetUserByID(ctx, comment.UserID)
	// if err != nil {
	// 	return entities.Comment{}, err
	// }
	// comment = user
	return comment, nil
}

func (db *commentRepository) GetAllComment(ctx context.Context) ([]entities.Comment, error) {
	var comments []entities.Comment
	if tx := db.connection.Find(&comments).Error; tx != nil {
		return nil, tx
	}
	return comments, nil
}

func (db *commentRepository) GetCommentsByBlogID(ctx context.Context, blogID uuid.UUID) ([]entities.Blog, error) {
	var comments []entities.Blog
	if tx := db.connection.Preload("Comments").Where("id = ?", blogID).Find(&comments).Error; tx != nil {
		return nil, tx
	}
	return comments, nil
}

func (db *commentRepository) UpdateComment(ctx context.Context, comment entities.Comment) error {
	if tx := db.connection.Save(&comment).Error; tx != nil {
		return tx
	}
	return nil
}

func (db *commentRepository) DeleteComment(ctx context.Context, userID uuid.UUID) error {
	if tx := db.connection.Delete(&entities.Comment{}, &userID).Error; tx != nil {
		return tx
	}
	return nil
}
