package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"taskman4/models"
)

func (h *Handler) getTask(c *gin.Context) {
	id := c.GetInt("id")
	userID := c.GetInt("user_id")
	task, err := h.services.GetTask(c.Request.Context(), id, userID)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *Handler) getAllTasks(c *gin.Context) {
	userID := c.GetInt("user_id")
	tasks, err := h.services.GetAllTasks(c.Request.Context(), userID)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) addTask(c *gin.Context) {
	var t models.Task
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	t.UserID = c.GetInt("user_id")

	id, err := h.services.AddTask(c.Request.Context(), t)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success", "id": id})
}

func (h *Handler) updateTask(c *gin.Context) {
	var t models.Task
	if err := c.BindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	userID := c.GetInt("user_id")
	id := c.GetInt("id")

	err := h.services.UpdateTask(c.Request.Context(), id, userID, t)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (h *Handler) deleteTask(c *gin.Context) {
	id := c.GetInt("id")
	userID := c.GetInt("user_id")
	err := h.services.DeleteTask(c.Request.Context(), id, userID)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (h *Handler) markTask(c *gin.Context) {
	id := c.GetInt("id")
	userID := c.GetInt("user_id")
	err := h.services.MarkTask(c.Request.Context(), id, userID)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}

func (h *Handler) getOverdueTasks(c *gin.Context) {
	userID := c.GetInt("user_id")
	tasks, err := h.services.GetOverdueTasks(c.Request.Context(), userID)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) reassignTask(c *gin.Context) {
	var req struct {
		TaskID  int    `json:"task_id"`
		NewUser string `json:"username"`
	}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.BadRequest)
		return
	}
	userID := c.GetInt("user_id")

	err := h.services.ReassignTask(c.Request.Context(), userID, req.TaskID, req.NewUser)
	if err != nil {
		models.ReplyError(c, err)
		return
	}
	c.JSON(http.StatusOK, models.OK)
}
