package repository_test

import (
	"absensi_mahasiswa_uas/src"
	"absensi_mahasiswa_uas/src/model"
	"absensi_mahasiswa_uas/src/repository"
	"testing"

	"gorm.io/gorm"
)

func TestFindMahasiswaAll(t *testing.T) {
	db := src.NewDBConnection()
	r := repository.NewMahasiswaRepository(db)

	data := []model.Dataset{}
	err := r.Find(&data)
	if err != nil {
		t.Fatal("Test Error:", err.Error())
	}

	t.Log(data)
}
func TestFindMahasiswaByID(t *testing.T) {
	db := src.NewDBConnection()
	r := repository.NewMahasiswaRepository(db)

	res := r.FindByID("7")
	if res.Error == gorm.ErrRecordNotFound {
		t.Fatal("Fatal Error:", res.Error.Error())
	}

	t.Log(res.Model)
}

func TestDeleteMahasiswaByID(t *testing.T) {
	db := src.NewDBConnection()
	r := repository.NewMahasiswaRepository(db)

	res := r.DeleteByID("4")
	if res.Error != nil {
		t.Fatal(res.Error.Error())
	}

	t.Log(res.Model)
}

func TestFindMahasiswaByIdentifier(t *testing.T) {
	db := src.NewDBConnection()
	r := repository.NewMahasiswaRepository(db)

	res := r.FindByIdentifier("se3")
	if res.Error == gorm.ErrRecordNotFound {
		t.Fatal("Fatal Error:", res.Error.Error())
	}

	t.Log(res.Model)
}
