package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/models"
)

type CommentController struct{}

func getCommentFromId(commentId string) (*models.Comment, error) {
	db := database.GetDatabase()

	var matchedComment *models.Comment
	if err := db.First(&matchedComment, commentId).Error; err != nil {
		return nil, err
	}
	return matchedComment, nil
}

func (comment CommentController) CreateComment(c *gin.Context) {
	db := database.GetDatabase()

	var reqComment models.Comment
	if err := c.BindJSON(&reqComment); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Create(&reqComment).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (comment CommentController) UpdateComment(c *gin.Context) {
	db := database.GetDatabase()

	var updatedComment models.Comment
	if err := c.BindJSON(&updatedComment); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	matchedComment, err := getCommentFromId(c.Param("id"))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if infoIsChanged(updatedComment.Content, matchedComment.Content) {
		matchedComment.Content = updatedComment.Content
	}

	if err := db.Save(&matchedComment).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (comment CommentController) DeleteComment(c *gin.Context) {
	db := database.GetDatabase()

	matchedComment, err := getCommentFromId(c.Param("id"))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Delete(&matchedComment).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}
