package routers

import (
	"api-backend/config"
	"api-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB, router *gin.Engine) {

	router.POST("api/v1/user/access/post", func(c *gin.Context) {
		var loginReq struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		var user models.Users
		if err := c.BindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Richiesta non valida"})
			return
		}
		if err := db.Where("username = ? AND password = ?", loginReq.Username, loginReq.Password).First(&user).Error; err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": "Accesso negato, un utente con queste credenziali non esiste"},
			)
			return
		}
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": "success",
				"token":  user.Token,
			},
		)

	})

	group := router.Group("test/api/v1/backend-api")

	group.Use(config.AuthMiddleware())

	group.POST("/projects/post", func(c *gin.Context) {
		var project models.ProjectModel
		if err := c.ShouldBindJSON(&project); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&project).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, project)
	})

	group.GET("/projects/get", func(c *gin.Context) {
		var projects []models.ProjectModel
		if err := db.Find(&projects).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, projects)
	})

	group.GET("/projects/get/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
			return
		}

		var project models.ProjectModel
		if err := db.Where("id = ?", id).First(&project).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		c.JSON(http.StatusOK, project)
	})

	group.GET("/projects/:id/tasks", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
			return
		}

		var tasks []models.TaskModel
		if err := db.Where("project_id = ?", id).Find(&tasks).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, tasks)
	})

	group.POST("/tasks/post", func(c *gin.Context) {
		var task models.TaskModel

		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&task).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
	})
}
