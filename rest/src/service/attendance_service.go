package service

import (
	"absensi_mahasiswa_uas_rest/src/model"
	"errors"

	"gorm.io/gorm"
)

type AttendanceService interface {
	Create(data *model.Attendance) error
	Find(data *[]model.Attendance, conds ...interface{}) error
	FindByID(data *model.Attendance, id string) error
}

type attService struct {
	repository AttendanceService
}

func NewAttendanceService(repository AttendanceService) *attService {
	return &attService{repository}
}

func (att attService) Create(data *model.Attendance) error {
	return att.repository.Create(data)
}

func (att attService) Find() ([]model.Attendance, error) {
	var data []model.Attendance

	err := att.repository.Find(&data)
	return data, err
}

func (att attService) FindByID(id string) (model.Attendance, error) {
	var data model.Attendance

	err := att.repository.FindByID(&data, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return data, errors.New("The item does not exist")
	}

	return data, nil
}
