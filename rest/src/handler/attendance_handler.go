package handler

import (
	"absensi_mahasiswa_uas_rest/src/model"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AttendanceService interface {
	Create(data *model.Attendance) error
	Find() ([]model.Attendance, error)
	FindByID(id string) (model.Attendance, error)
}

type attHandler struct {
	service AttendanceService
}

func NewAttendanceHandler(service AttendanceService) *attHandler {
	return &attHandler{service}
}

var MediaLocation = "media/"

type AttendanceRequest struct {
	Name       string                `form:"name" json:"name" binding:"required"`
	Image      *multipart.FileHeader `form:"image" json:"image" binding:"required"`
	Identifier string                `form:"identifier" json:"identifier" binding:"required"`
}

func (att attHandler) CreateAttendanceHandler(c *gin.Context) {
	var bind AttendanceRequest
	var uuid string = uuid.NewString()

	if err := c.ShouldBind(&bind); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	filename := fmt.Sprint(MediaLocation, uuid, ".", strings.Split(bind.Image.Filename, ".")[1])
	err := c.SaveUploadedFile(bind.Image, filename)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	data := model.Attendance{
		Name:       bind.Name,
		Image:      filename,
		Identifier: bind.Identifier,
	}
	if err := att.service.Create(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data":    data,
		"message": "The item was crated successfully",
	})
}

func (att attHandler) FindAttendance(c *gin.Context) {
	data, err := att.service.Find()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (att attHandler) FindAttendanceWeb(c *gin.Context) {
	data, err := att.service.Find()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	c.HTML(http.StatusOK, "attendance_list.html", gin.H{
		"data": data,
	})
}

func (att attHandler) FindAttendanceByID(c *gin.Context) {
	id := c.Param("id")
	data, err := att.service.FindByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
