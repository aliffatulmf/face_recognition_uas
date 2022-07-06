package main

import (
	"absensi_mahasiswa_uas_rest/src"
	"absensi_mahasiswa_uas_rest/src/handler"
	"absensi_mahasiswa_uas_rest/src/model"
	"absensi_mahasiswa_uas_rest/src/repository"
	"absensi_mahasiswa_uas_rest/src/service"

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
	attRepo := repository.NewAttendance(s.DB)
	attServ := service.NewAttendanceService(attRepo)
	attHandler := handler.NewAttendanceHandler(attServ)

	v1 := s.Router.Group("/api/v1")
	v1.GET("attendance/", attHandler.FindAttendance)
	v1.GET("attendance/:id", attHandler.FindAttendanceByID)
	v1.POST("attendance/", attHandler.CreateAttendanceHandler)

	s.Router.GET("list", attHandler.FindAttendanceWeb)
}

func (s *Server) RunDebug(addr string) (err error) {
	err = s.Router.Run(addr)
	return
}

func (s *Server) EnableHTMLTemplates(pattern Templates) {
	s.Router.LoadHTMLGlob(pattern.Templates)
}

func (s *Server) StaticHandler() {
	s.Router.Static("/dist", "dist/")
	s.Router.Static("/media", "media/")
}

func main() {
	server := &Server{
		DB:     src.NewDBConnection(),
		Router: gin.Default(),
	}

	server.DB.AutoMigrate(model.Attendance{})
	server.EnableHTMLTemplates(Templates{
		Assets: assets{
			path:     "/dist",
			location: "assets/",
		},
		Templates: "templates/*.html",
	})
	server.StaticHandler()
	server.LoadHandlers()
	server.RunDebug(":8081")
}
