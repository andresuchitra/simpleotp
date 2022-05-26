package handler

import (
	"net/http"

	"github.com/andresuchitra/simpleotp/models"
	"github.com/andresuchitra/simpleotp/service"
	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Status uint32      `json:"status"`
	Data   interface{} `json:"data"`
}
type ResponseError struct {
	Status uint32 `json:"status"`
	Error  string `json:"error"`
}

type handler struct {
	service service.OTPService
}

func NewHandler(r *gin.Engine, service service.OTPService) {
	h := handler{
		service: service,
	}

	// setup endpoints
	r.POST("/generate", h.GenerateOTP)
	r.POST("/verify", h.VerifyOTP)
}

func (h *handler) GenerateOTP(c *gin.Context) {
	var newOtp *models.OTPItem

	// check the request payload format
	if binderr := c.ShouldBindJSON(&newOtp); binderr != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Error: "Invalid payload format",
		})
	}
	ctx := c.Request.Context()

	otp, err := h.service.CreateOTP(&ctx, newOtp.Phone)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, ResponseError{
			Error: err.Error(),
		})
	}

	c.JSON(http.StatusOK, ResponseData{
		Status: 200,
		Data:   otp,
	})
}

func (h *handler) VerifyOTP(c *gin.Context) {
	c.JSON(http.StatusOK, ResponseData{
		Status: 200,
		Data:   "Done VerifyOTP",
	})
}
