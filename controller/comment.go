package controller

import (
	"net/http"

	"github.com/Caknoooo/Golang-BLOG/dto"
	"github.com/Caknoooo/Golang-BLOG/services"
	"github.com/Caknoooo/Golang-BLOG/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CommentController interface {
	CreateComment(ctx *gin.Context)
	GetAllComment(ctx *gin.Context)
	GetCommentByBlogID(ctx *gin.Context)
	UpdateComment(ctx *gin.Context)
	DeleteComment(ctx *gin.Context)
}

type commentController struct {
	jwtService     services.JWTService
	commentService services.CommentService
}

func NewCommentController(cs services.CommentService, jwt services.JWTService) CommentController {
	return &commentController{
		jwtService:     jwt,
		commentService: cs,
	}
}

func (cc *commentController) CreateComment(ctx *gin.Context) {
	var commentDTO dto.CommentCreateDTO
	if err := ctx.ShouldBindJSON(&commentDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	comment, err := cc.commentService.CreateComment(ctx, commentDTO)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Comment", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan Komentar", comment)
	ctx.JSON(http.StatusOK, res)
}

func (cc *commentController) GetAllComment(ctx *gin.Context) {
	comments, err := cc.commentService.GetAllComment(ctx)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Comment", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Comment", comments)
	ctx.JSON(http.StatusOK, res)
}

func (cc *commentController) GetCommentByBlogID(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	// fmt.Println()
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := cc.commentService.GetCommentByBlogID(ctx, uuid)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Komentar Dari Blog", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Komentar Dari Blog", result)
	ctx.JSON(http.StatusOK, res)
}

func (cc *commentController) UpdateComment(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var commentDTO dto.CommentUpdateDTO
	if err := ctx.ShouldBindJSON(&commentDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	commentDTO.ID = uuid
	if err = cc.commentService.UpdateComment(ctx, commentDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mengupdate Comment", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mengupdate Comment", commentDTO)
	ctx.JSON(http.StatusOK, res)
}

func (cc *commentController) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err = cc.commentService.DeleteBlog(ctx, uuid); err != nil {
		res := utils.BuildResponseFailed("Gagal Delete Comment", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Delete Comment", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
