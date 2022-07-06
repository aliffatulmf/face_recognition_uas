package handler

import (
	"absensi_mahasiswa_uas/src"
	"absensi_mahasiswa_uas/src/model"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DatasetService interface {
	Create(data model.Dataset) (model.Dataset, error)
	FindByIdentifier(id string) (model.Dataset, error)
	Find() ([]model.Dataset, error)
}

type datasetHandler struct {
	service DatasetService
}

type DatasetFile struct {
	Form   model.Dataset
	Images []*multipart.FileHeader `form:"images"`
}

func NewWebHandler(service DatasetService) datasetHandler {
	return datasetHandler{service}
}

func (web datasetHandler) GetMahasiswaList(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func (web datasetHandler) CreateDataset(ctx *gin.Context) {
	var data DatasetFile

	err := ctx.ShouldBind(&data)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message": "Dataset error",
			"error":   true,
		})
		return
	}

	_, err = web.service.Create(data.Form)
	if err != nil {
		ctx.HTML(http.StatusBadRequest, "index.html", gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	p, _ := filepath.Abs(
		fmt.Sprintf("face_recognition/dataset/%s.%s", data.Form.Name, data.Form.Identifier),
	)

	if _, err := os.Stat(p); errors.Is(err, fs.ErrNotExist) {
		err = os.MkdirAll(p, os.ModeDir)
		if err != nil {
			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"message": err.Error(),
				"error":   true,
			})
			return
		}
	}

	for _, file := range data.Images {
		ext := strings.Split(file.Filename, ".")
		if err := ctx.SaveUploadedFile(file, fmt.Sprintf("%s/%s.%s", p, uuid.NewString(), ext[len(ext)-1])); err != nil {
			panic(err)
		}
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"message": fmt.Sprintf("%d files uploaded", len(data.Images)),
		"success": true,
	})
}

func (web datasetHandler) DatasetList(ctx *gin.Context) {
	res, _ := web.service.Find()

	ctx.HTML(http.StatusOK, "dataset_list.html", gin.H{
		"data": res,
	})
}

func (web datasetHandler) DatasetTraining(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "dataset_training.html", gin.H{})
}

func (web datasetHandler) FaceFront(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "face_recognize.html", gin.H{})
}

func (web datasetHandler) DatasetExecutor(ctx *gin.Context) {
	output := src.TrainingModel()

	ctx.JSON(http.StatusOK, gin.H{
		"data": output,
	})
}

type FaceRecognition struct {
	Name   string                `form:"name"`
	Target *multipart.FileHeader `form:"target"`
}

func (web datasetHandler) FaceRecognize(ctx *gin.Context) {
	var data FaceRecognition

	if err := ctx.ShouldBind(&data); err != nil {
		ctx.HTML(http.StatusOK, "face_recognize.html", gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	uuid := uuid.NewString()
	ext := strings.Split(data.Target.Filename, ".")
	ipath := filepath.Join("face_recognition/temp", fmt.Sprintf("%s.%s", uuid, ext[len(ext)-1]))

	err := ctx.SaveUploadedFile(data.Target, ipath)
	if err != nil {
		ctx.HTML(http.StatusOK, "face_recognize.html", gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	label := src.FaceRecognize(ipath)
	if label != "Unknown" {
		arr := strings.Split(label, ".")
		trimed := strings.TrimSuffix(arr[1], "\r\n")
		fnd, _ := web.service.FindByIdentifier(trimed)

		label = fmt.Sprintf(
			"Found\nName: %s\nIdentifier: %s", fnd.Name, fnd.Identifier,
		)

		// POST to server
		host := "http://103.189.234.30:8081/api/v1/attendance"

		payload := strings.NewReader(
			fmt.Sprintf("-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"name\"\r\n\r\n%s\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"image\"; filename=\"%s\"\r\nContent-Type: image/png\r\n\r\n\r\n-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"identifier\"\r\n\r\n%s\r\n-----011000010111000001101001--\r\n", fnd.Name, ipath, fnd.Identifier),
		)

		req, _ := http.NewRequest("POST", host, payload)

		req.Header.Add("Content-Type", "multipart/form-data; boundary=---011000010111000001101001")

		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(string(body))
	}

	ctx.HTML(http.StatusOK, "face_recognize.html", gin.H{
		"message": label,
		"success": true,
	})
}
