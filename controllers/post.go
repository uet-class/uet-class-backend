package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/models"
)

type PostController struct{}

func getPostFromId(postId string) (*models.Post, error) {
	db := database.GetDatabase()

	var matchedPost *models.Post
	if err := db.First(&matchedPost, postId).Error; err != nil {
		return nil, err
	}
	return matchedPost, nil
}

func (post PostController) CreatePost(c *gin.Context) {
	db := database.GetDatabase()

	var reqPost models.Post
	if err := c.BindJSON(&reqPost); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Create(&reqPost).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (post PostController) GetAllPosts(c *gin.Context) {
	db := database.GetDatabase()

	var allPosts []models.Post

	if err := db.Find(&allPosts).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, allPosts)
}

func (post PostController) GetPost(c *gin.Context) {
	db := database.GetDatabase()

	var matchedPost models.Post
	if err := db.Preload("Comment").First(&matchedPost, c.Param("id")).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, matchedPost)
}

func (post PostController) UpdatePost(c *gin.Context) {
	db := database.GetDatabase()

	var updatedPost models.Post
	if err := c.BindJSON(&updatedPost); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	matchedPost, err := getPostFromId(c.Param("id"))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if infoIsChanged(updatedPost.Title, matchedPost.Title) {
		matchedPost.Title = updatedPost.Title
	}

	if infoIsChanged(updatedPost.Content, matchedPost.Content) {
		matchedPost.Content = updatedPost.Content
	}

	if err := db.Save(&matchedPost).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (post PostController) DeletePost(c *gin.Context) {
	db := database.GetDatabase()

	matchedPost, err := getPostFromId(c.Param("id"))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Delete(&matchedPost).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}
