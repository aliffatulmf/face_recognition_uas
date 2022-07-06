package service

import (
	"absensi_mahasiswa_uas/src/model"
	"absensi_mahasiswa_uas/src/repository"
)

type MahasiswaRepository interface {
	Create(data *model.Dataset) error
	Find(data *[]model.Dataset, conds ...interface{}) error
	FindByID(id string) *repository.DatasetRepository
	FindByIdentifier(id string) *repository.DatasetRepository
	DeleteByID(id string) *repository.DatasetRepository
}

type datasetService struct {
	Repository MahasiswaRepository
	Model      model.Dataset
}

func NewDatasetService(repository MahasiswaRepository) *datasetService {
	return &datasetService{Repository: repository}
}

func (s *datasetService) Create(data model.Dataset) (model.Dataset, error) {
	err := s.Repository.Create(&data)
	if err != nil {
		return model.Dataset{}, err
	}

	return data, nil
}

func (s *datasetService) Find() ([]model.Dataset, error) {
	var data []model.Dataset
	err := s.Repository.Find(&data)
	if err != nil {
		return []model.Dataset{}, err
	}

	return data, nil
}

func (s *datasetService) FindByID(id string) (model.Dataset, error) {
	tr := s.Repository.FindByID(id)
	if tr.Error != nil {
		return tr.Model, tr.Error
	}

	return tr.Model, nil
}

func (s *datasetService) FindByIdentifier(id string) (model.Dataset, error) {
	tr := s.Repository.FindByIdentifier(id)
	if tr.Error != nil {
		return tr.Model, tr.Error
	}

	return tr.Model, nil
}

func (s *datasetService) DeleteByID(id string) error {
	tr := s.Repository.DeleteByID(id)
	if tr.Error != nil {
		return tr.Error
	}

	return nil
}
