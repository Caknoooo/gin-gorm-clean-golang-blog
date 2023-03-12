package services

import (
	"context"

	"github.com/Caknoooo/Golang-BLOG/dto"
	"github.com/Caknoooo/Golang-BLOG/entities"
	"github.com/Caknoooo/Golang-BLOG/repository"
	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type CommentService interface {
	CreateComment(ctx context.Context, commentDTO dto.CommentCreateDTO) (entities.Comment, error)
	GetAllComment(ctx context.Context) ([]entities.Comment, error)
	GetCommentByBlogID(ctx context.Context, blogID uuid.UUID) ([]entities.Blog, error)
	UpdateComment(ctx context.Context, commentDTO dto.CommentUpdateDTO) (error)
	DeleteBlog(ctx context.Context, userID uuid.UUID) (error) 
}

type commentService struct {
	commentRepository repository.CommentRepository
}

func NewCommentService(cr repository.CommentRepository) CommentService {
	return &commentService{
		commentRepository: cr,
	}
}

func (cs *commentService) CreateComment(ctx context.Context, commentDTO dto.CommentCreateDTO) (entities.Comment, error) {
	comment := entities.Comment{}
	if err := smapping.FillStruct(&comment, smapping.MapFields(commentDTO)); err != nil {
		return comment, nil
	}
	return cs.commentRepository.CreateComment(ctx, comment)
}

func (cs *commentService) GetAllComment(ctx context.Context) ([]entities.Comment, error) {
	return cs.commentRepository.GetAllComment(ctx)
}

func (cs *commentService) GetCommentByBlogID(ctx context.Context, blogID uuid.UUID) ([]entities.Blog, error) {
	return cs.commentRepository.GetCommentsByBlogID(ctx, blogID)
}

func (cs *commentService) UpdateComment(ctx context.Context, commentDTO dto.CommentUpdateDTO) (error) {
	comment := entities.Comment{}
	if err := smapping.FillStruct(&comment, smapping.MapFields(&commentDTO)); err != nil {
		return nil
	}
	return cs.commentRepository.UpdateComment(ctx, comment)
}

func (cs *commentService) DeleteBlog(ctx context.Context, userID uuid.UUID) (error) {
	return cs.commentRepository.DeleteComment(ctx, userID)
}