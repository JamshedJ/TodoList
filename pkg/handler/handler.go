package handler

import (
	"taskman4/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	auth := router.Group("/auth")
	auth.POST("/signup", h.signUp)
	auth.POST("/signin", h.signIn)

	users := router.Group("/users", h.mwUserAuth, MwGetID)
	users.GET("/me", h.getUser)
	users.GET("/", h.getAllUsers)
	users.PUT("/", h.updateUser)
	users.DELETE("/", h.deleteUser)

	tasks := router.Group("/tasks", h.mwUserAuth, MwGetID)
	tasks.POST("/", h.addTask)
	tasks.GET("/:id", h.getTask)
	tasks.GET("/", h.getAllTasks)
	tasks.GET("/overdue", h.getOverdueTasks)
	tasks.PUT("/:id", h.updateTask)
	tasks.DELETE("/:id", h.deleteTask)
	tasks.PATCH("/:id", h.markTask)
	tasks.PATCH("/", h.reassignTask)

	return router
}
