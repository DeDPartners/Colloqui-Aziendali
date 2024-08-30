package main

import (
	"log"
	"os"
	"time"

  "api-backend/routers"
  "api-backend/models"
  "api-backend/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := config.ConnectDatabase()

	if err != nil {
		log.Fatal("Impossibile connettersi al database:", err)
	}

	// config.GetBlankDB(db)
	db.AutoMigrate(&models.Users{})
	db.AutoMigrate(&models.TaskModel{})
	db.AutoMigrate(&models.ProjectModel{})

	r := gin.Default()
	r.Use(
		cors.New(
			cors.Config{
				AllowOrigins: []string{"*"},
				AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				AllowHeaders: []string{
					"Origin",
					"Content-Type",
					"Content-Length",
					"Authorization",
					"Accept",
					"Access-Control-Allow-Origin",
					"Access-Control-Allow-Headers",
					"Access-Control-Allow-Methods",
					"Access-Control-Allow-Credentials",
				},
				ExposeHeaders: []string{
					"Content-Length",
					"Authorization",
					"Content-Type",
					"Pdf-Name",
				},
				AllowCredentials: true,
				MaxAge:           6 * time.Hour,
			}))

	r.SetTrustedProxies(nil)
	r.Static("/media", "media")
	routers.Routes(db, r)

	port := os.Getenv("PORT_backend")
	if port == "" {
		log.Fatal("PORT_BACKEND non è impostato.")
	} else {
		err = r.Run(":" + port)
		if err != nil {
			log.Fatal("Errore nell'avvio del server:", err)
		}
	}
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Erore nell'avvio del server:", err)
	}
}
