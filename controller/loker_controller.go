package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/onainadapdap1/kartu_prakerja/simple_unjuk_ket/prakerja-final/config"
	"github.com/onainadapdap1/kartu_prakerja/simple_unjuk_ket/prakerja-final/helpers"
	"github.com/onainadapdap1/kartu_prakerja/simple_unjuk_ket/prakerja-final/models"
)

func AddLokerController(c echo.Context) error {
	var newLoker models.Loker
	c.Bind(&newLoker)

	var nasabah models.User
	cekNasabah := config.DB.Where("id = ?", &newLoker.UserId).First(&nasabah)
	if cekNasabah.Error == gorm.ErrRecordNotFound {
		return c.JSON(http.StatusBadRequest, helpers.FailResponse{
			Message: "Gagal membuat loker, user tidak ditemukan.",
		})
	}

	var loker models.Loker
	cekLoker := config.DB.Where("no_loker = ?", &newLoker.NoLoker).First(&loker)
	if cekLoker.Error == gorm.ErrRecordNotFound {
		result := config.DB.Create(&newLoker)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, helpers.FailResponse{
				Message: "Terjadi kesalahan server, loker gagal disimpan.",
			})
		}
		return c.JSON(http.StatusCreated, helpers.BaseResponse{
			Message: "Success",
			Data:    newLoker,
		})
	}
	return c.JSON(http.StatusBadRequest, helpers.FailResponse{
		Message: "Failed, No Loker sudah terdaftar.",
	})
}

func GetLokerController(c echo.Context) error {
	var loker []models.Loker
	result := config.DB.Find(&loker)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailResponse{
			Message: "Terjadi kesalahan server, Gagal mendapatkan data loker.",
		})
	}

	return c.JSON(http.StatusOK, helpers.BaseResponse{
		Message: "Success",
		Data:    loker,
	})
}