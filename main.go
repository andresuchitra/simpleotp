package main

import (
	"log"
	"net/http"

	"github.com/andresuchitra/simpleotp/handler"
	"github.com/andresuchitra/simpleotp/repository"
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

	DB := setupDB()

	repo := repository.NewRepository(DB)
	r := setupRouter()
	handler.NewHandler(r, repo)

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
