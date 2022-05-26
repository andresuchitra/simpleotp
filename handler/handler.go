package handler

import (
	"net/http"

	"github.com/andresuchitra/simpleotp/models"
	"github.com/andresuchitra/simpleotp/repository"
	"github.com/andresuchitra/simpleotp/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResponseData struct {
	Status uint32 `json:"status"`
	Data   string `json:"data"`
}
type ResponseError struct {
	Status uint32 `json:"status"`
	Error  string `json:"error"`
}

type handler struct {
	repo *repository.OTPRepository
}

func NewHandler(r *gin.Engine, repo *repository.OTPRepository) {
	h := handler{
		repo: repo,
	}

	// setup endpoints
	r.POST("/generate", h.GenerateOTP)
	r.POST("/verify", h.VerifyOTP)
}

func (h *handler) GenerateOTP(c *gin.Context) {
	var newOtp models.OTPItem

	// check the request payload format
	if binderr := c.ShouldBindJSON(&newOtp); binderr != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Error: "Invalid payload format",
		})
	}

	// generate UUID
	uuid := uuid.New()
	newOtp.ID = uuid.String()

	// generate token
	otpManager := service.NewOTPManager(6)
	otpToken, err := otpManager.GenerateOTP()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ResponseError{
			Status: http.StatusUnprocessableEntity,
			Error:  err.Error(),
		})
	}

	// store the otp to current mobile phone request

	c.JSON(http.StatusOK, ResponseData{
		Status: 200,
		Data:   otpToken,
	})
}

func (h *handler) VerifyOTP(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseData{
		Status: 200,
		Data:   "Done VerifyOTP",
	})
}
