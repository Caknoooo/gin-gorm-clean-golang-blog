package repository

import (
	"context"

	"github.com/Caknoooo/Golang-BLOG/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogRepository interface {
	GetAllBlogs(ctx context.Context) ([]entities.Blog, error)
	GetBlogByID(ctx context.Context, blogID uuid.UUID) (entities.Blog, error)
	CreateBlog(ctx context.Context, blog entities.Blog) (entities.Blog, error)
	LikeBlogByBlogID(ctx context.Context, userID uuid.UUID, blogID uuid.UUID) error
	UpdateBlog(ctx context.Context, blog entities.Blog) error
	DeleteBlog(ctx context.Context, blogID uuid.UUID) error
	// DeleteLikesByBlogID(ctx context.Context, blogID uuid.UUID) error
}

type blogConnection struct {
	connection     *gorm.DB
	userRepository UserRepository
}

func NewBlogRepository(db *gorm.DB, userRepository UserRepository) BlogRepository {
	return &blogConnection{
		connection:     db,
		userRepository: userRepository,
	}
}

func (bc *blogConnection) GetAllBlogs(ctx context.Context) ([]entities.Blog, error) {
	var blogs []entities.Blog
	if tx := bc.connection.Preload("User").Preload("Likes").Preload("Comments").Find(&blogs).Error; tx != nil {
		return nil, tx
	}
	return blogs, nil
}

func (bc *blogConnection) GetBlogByID(ctx context.Context, blogID uuid.UUID) (entities.Blog, error) {
	var blog entities.Blog

	if tx := bc.connection.Preload("Likes").Preload("Comments").Where("id = ?", blogID).Take(&blog).Error; tx != nil {
		return entities.Blog{}, tx
	}

	user, err := bc.userRepository.GetUserByID(ctx, blog.UserID)
	if err != nil {
		return entities.Blog{}, err
	}
	blog.User = user
	return blog, nil
}

func (bc *blogConnection) CreateBlog(ctx context.Context, blog entities.Blog) (entities.Blog, error) {
	if tx := bc.connection.Create(&blog).Error; tx != nil {
		return entities.Blog{}, tx
	}
	user, err := bc.userRepository.GetUserByID(ctx, blog.UserID)
	if err != nil {
		return entities.Blog{}, err
	}
	blog.User = user
	return blog, nil
}

func (bc *blogConnection) LikeBlogByBlogID(ctx context.Context, userID uuid.UUID, blogID uuid.UUID) error {
	var like entities.Like
	if tx := bc.connection.Where("user_id = ? AND blog_id = ?", userID, blogID).Find(&like).Error; tx != nil {
		return tx
	}

	like = entities.Like{
		BlogID: blogID,
		UserID: userID,
	}

	if tx := bc.connection.Create(&like).Error; tx != nil {
		return tx
	}

	var blog entities.Blog
	if tx := bc.connection.Where("id = ?", blogID).Find(&blog).Error; tx != nil {
		return tx
	}
	blog.LikeCount++
	bc.UpdateBlog(ctx, blog)

	return nil
}

func (bc *blogConnection) UpdateBlog(ctx context.Context, blog entities.Blog) error {
	if tx := bc.connection.Save(&blog).Error; tx != nil {
		return tx
	}
	return nil
}

func (bc *blogConnection) DeleteBlog(ctx context.Context, blogID uuid.UUID) error {
	// if tx := bc.connection.Where("blog_id = ?", blogID).Delete(&entities.Like{}).Error; tx != nil {
	// 	return tx
	// }
	// if tx := bc.connection.Where("blog_id = ?", blogID).Delete(&entities.Comment{}).Error; tx != nil {
	// 	return tx
	// }
	if tx := bc.connection.Delete(&entities.Blog{}, "id = ?", &blogID).Error; tx != nil {
		return tx
	}
	return nil
}