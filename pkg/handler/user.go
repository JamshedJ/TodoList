package handler

import (
	"net/http"
	"taskman4/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getUser(c *gin.Context) {
	userID := c.GetInt("user_id")
	user, err := h.services.GetUser(c.Request.Context(), userID)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) getAllUsers(c *gin.Context) {
	users, err := h.services.GetAllUsers(c.Request.Context())
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) updateUser(c *gin.Context) {
	var u models.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	u.ID = c.GetInt("user_id")
	err := h.services.UpdateUser(c.Request.Context(), u)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (h *Handler) deleteUser(c *gin.Context) {
	userID := c.GetInt("user_id")
	err := h.services.DeleteUser(c.Request.Context(), userID)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}
