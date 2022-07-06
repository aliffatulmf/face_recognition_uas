package repository

import (
	"absensi_mahasiswa_uas_rest/src/model"

	"gorm.io/gorm"
)

type AttendanceRepository struct {
	DB    *gorm.DB
	Model model.Attendance
	Error error
}

func NewAttendance(db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{DB: db}
}

func (att AttendanceRepository) Create(data *model.Attendance) error {
	return att.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Select([]string{"Name", "NIM", "Image", "Identifier"}).Create(data).Error
		if err != nil {
			return err
		}
		return nil
	})
}

func (att AttendanceRepository) Find(data *[]model.Attendance, conds ...interface{}) error {
	return att.DB.Find(data, conds...).Error
}

func (att AttendanceRepository) FindByID(data *model.Attendance, id string) error {
	return att.DB.First(data, "id = ?", id).Error
}
