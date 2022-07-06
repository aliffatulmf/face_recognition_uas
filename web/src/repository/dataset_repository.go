package repository

import (
	"absensi_mahasiswa_uas/src/model"

	"gorm.io/gorm"
)

type DatasetRepository struct {
	DB    *gorm.DB
	Model model.Dataset
	Error error
}

func NewMahasiswaRepository(db *gorm.DB) *DatasetRepository {
	return &DatasetRepository{
		DB:    db,
		Error: nil,
	}
}

func (r DatasetRepository) Create(data *model.Dataset) error {
	res := r.DB.Create(data)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r DatasetRepository) Find(data *[]model.Dataset, conds ...interface{}) error {
	err := r.DB.Find(data, conds).Error
	if err != nil {
		return err
	}
	return nil
}

func (r DatasetRepository) FindByID(id string) *DatasetRepository {
	r.Error = r.DB.First(&r.Model, "id = ?", id).Error

	return &r
}

func (r DatasetRepository) FindByIdentifier(id string) *DatasetRepository {
	r.Error = r.DB.First(&r.Model, "identifier = ?", id).Error

	return &r
}

func (r DatasetRepository) DeleteByID(id string) *DatasetRepository {
	data := r.FindByID(id)
	if data.Error != nil {
		r.Error = data.Error
	} else {
		r.Error = r.DB.Delete(&data.Model, "id = ?", id).Error
	}

	return &r
}
