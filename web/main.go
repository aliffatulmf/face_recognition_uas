package main

import (
	"absensi_mahasiswa_uas/src"
	"absensi_mahasiswa_uas/src/handler"
	"absensi_mahasiswa_uas/src/model"
	"absensi_mahasiswa_uas/src/repository"
	"absensi_mahasiswa_uas/src/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

type assets struct {
	path     string
	location string
}
type Templates struct {
	Assets    assets
	Templates string
}

func (s *Server) LoadHandlers() {
	repo := repository.NewMahasiswaRepository(s.DB)
	serv := service.NewDatasetService(repo)
	webHandler := handler.NewWebHandler(serv)

	web := s.Router.Group("/web/")
	web.GET("dataset/", webHandler.GetMahasiswaList)
	web.GET("face/recognizing", webHandler.FaceFront)
	web.GET("dataset/list", webHandler.DatasetList)
	web.GET("dataset/training", webHandler.DatasetTraining)
	web.POST("dataset/training", webHandler.DatasetExecutor)
	web.POST("dataset/", webHandler.CreateDataset)

	web.POST("face/recognizing", webHandler.FaceRecognize)
}

func (s *Server) RunDebug(addr string) (err error) {
	err = s.Router.Run(addr)
	return
}

func (s *Server) EnableHTMLTemplates(pattern Templates) {
	s.Router.Static("/dist", "dist/")
	s.Router.LoadHTMLGlob(pattern.Templates)
}

func main() {
	server := &Server{
		DB:     src.NewDBConnection(),
		Router: gin.Default(),
	}

	server.DB.AutoMigrate(model.Dataset{})
	server.EnableHTMLTemplates(Templates{
		Assets: assets{
			path:     "/dist",
			location: "assets/",
		},
		Templates: "templates/*.html",
	})
	server.LoadHandlers()
	server.RunDebug(":8080")
}
