package controller

import (
	"net/http"

	"github.com/Caknoooo/Golang-BLOG/dto"
	"github.com/Caknoooo/Golang-BLOG/services"
	"github.com/Caknoooo/Golang-BLOG/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BlogController interface {
	CreateBlog(ctx *gin.Context)
	GetAllBlog(ctx *gin.Context)
	GetBlogByID(ctx *gin.Context)
	LikeBlogByBlogID(ctx *gin.Context)
	UpdateBlog(ctx *gin.Context)
	DeleteBlog(ctx *gin.Context)
}

type blogController struct {
	jwtService  services.JWTService
	blogService services.BlogService
}

func NewBlogController(bs services.BlogService, jwt services.JWTService) BlogController {
	return &blogController{
		jwtService:  jwt,
		blogService: bs,
	}
}

func (bc *blogController) CreateBlog(ctx *gin.Context) {
	var blogDTO dto.BlogCreateDTO
	if err := ctx.ShouldBindJSON(&blogDTO); err != nil {
		res := utils.BuildResponseFailed("Error", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	blog, err := bc.blogService.CreateBlog(ctx, blogDTO)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Menambahkan Blog", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Menambahkan Blog", blog)
	ctx.JSON(http.StatusOK, res)
}

func (bc *blogController) GetAllBlog(ctx *gin.Context) {
	blogs, err := bc.blogService.GetAllBlog(ctx)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan List Blog", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan List Blog", blogs)
	ctx.JSON(http.StatusOK, res)
}

func (bc *blogController) GetBlogByID(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	// fmt.Println()
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := bc.blogService.GetBlogByID(ctx, uuid)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Blog", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mendapatkan Blog", result)
	ctx.JSON(http.StatusOK, res)
}

func (bc *blogController) LikeBlogByBlogID(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	blog_id := ctx.Param("blog_id")
	user_uuid, err := uuid.Parse(user_id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	blog_uuid, err := uuid.Parse(blog_id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err = bc.blogService.LikeBlogByBlogID(ctx, user_uuid, blog_uuid); err != nil {
		res := utils.BuildResponseFailed("Gagal Like Blog", "Blog Tidak Ditemukan", utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Like Blog", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}

func (bc *blogController) UpdateBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	var blogDTO dto.BlogUpdateDTO
	if err := ctx.ShouldBindJSON(&blogDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mendapatkan Request Dari Body", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	blogDTO.ID = uuid
	if err = bc.blogService.UpdateBlog(ctx, blogDTO); err != nil {
		res := utils.BuildResponseFailed("Gagal Mengupdate Blog", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Mengupdate Blog", blogDTO)
	ctx.JSON(http.StatusOK, res)
}

func (bc *blogController) DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	uuid, err := uuid.Parse(id)
	if err != nil {
		res := utils.BuildResponseFailed("Gagal Parse Id", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err = bc.blogService.DeleteBlog(ctx, uuid); err != nil {
		res := utils.BuildResponseFailed("Gagal Delete Blog", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponseSuccess("Berhasil Delete Blog", utils.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
