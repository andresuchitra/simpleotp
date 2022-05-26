package main

import (
	"log"
	"net/http"

	"github.com/andresuchitra/simpleotp/handler"
	"github.com/andresuchitra/simpleotp/repository"
	"github.com/andresuchitra/simpleotp/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := setupRouter()
	DB := setupDB()

	repo := repository.NewRepository(DB)
	service := service.NewOTPService(repo)
	handler.NewHandler(r, service)

	log.Fatal(r.Run(":3333"))
}

func setupDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   "Hello",
		})
	})

	return r
}
