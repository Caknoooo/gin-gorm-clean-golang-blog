package services

import (
	"context"

	"github.com/Caknoooo/Golang-BLOG/dto"
	"github.com/Caknoooo/Golang-BLOG/entities"
	"github.com/Caknoooo/Golang-BLOG/repository"
	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type BlogService interface {
	GetAllBlog(ctx context.Context) ([]entities.Blog, error)
	GetBlogByID(ctx context.Context, userID uuid.UUID) (entities.Blog, error)
	CreateBlog(ctx context.Context, blogDTO dto.BlogCreateDTO) (entities.Blog, error)
	LikeBlogByBlogID(ctx context.Context, userID uuid.UUID, blogID uuid.UUID) (error)
	UpdateBlog(ctx context.Context, blogDTO dto.BlogUpdateDTO) error
	DeleteBlog(ctx context.Context, userID uuid.UUID) error
}

type blogService struct {
	blogRepository repository.BlogRepository
}

func NewBlogService(bs repository.BlogRepository) BlogService {
	return &blogService{
		blogRepository: bs,
	}
}

func (bs *blogService) CreateBlog(ctx context.Context, blogDTO dto.BlogCreateDTO) (entities.Blog, error) {
	blog := entities.Blog{}
	if err := smapping.FillStruct(&blog, smapping.MapFields(blogDTO)); err != nil {
		return blog, nil
	}
	return bs.blogRepository.CreateBlog(ctx, blog)
}

func (bs *blogService) LikeBlogByBlogID(ctx context.Context, userID uuid.UUID, blogID uuid.UUID) (error) {
	return bs.blogRepository.LikeBlogByBlogID(ctx, userID, blogID)
}

func (bs *blogService) GetAllBlog(ctx context.Context) ([]entities.Blog, error) {
	return bs.blogRepository.GetAllBlogs(ctx)
}

func (bs *blogService) GetBlogByID(ctx context.Context, userID uuid.UUID) (entities.Blog, error) {
	return bs.blogRepository.GetBlogByID(ctx, userID)
}

func (bs *blogService) UpdateBlog(ctx context.Context, blogDTO dto.BlogUpdateDTO) error {
	blog := entities.Blog{}
	if err := smapping.FillStruct(&blog, smapping.MapFields(blogDTO)); err != nil {
		return nil
	}
	return bs.blogRepository.UpdateBlog(ctx, blog)
}

func (bs *blogService) DeleteBlog(ctx context.Context, userID uuid.UUID) error {
	return bs.blogRepository.DeleteBlog(ctx, userID)
}
