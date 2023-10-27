package controller

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/onainadapdap1/kartu_prakerja/simple_unjuk_ket/prakerja-final/config"
	"github.com/onainadapdap1/kartu_prakerja/simple_unjuk_ket/prakerja-final/helpers"
	"github.com/onainadapdap1/kartu_prakerja/simple_unjuk_ket/prakerja-final/models"
)

func AddUserController(c echo.Context) error {
	var newUser models.User
	c.Bind(&newUser)

	var user models.User
	cekUser := config.DB.Where("nik = ?", &newUser.NIK).First(&user)
	if cekUser.Error == gorm.ErrRecordNotFound {
		result := config.DB.Create(&newUser)
		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, helpers.FailResponse{
				Message: "Terjadi kesalahan server, data user gagal disimpan.",
			})
		}
		return c.JSON(http.StatusCreated, helpers.BaseResponse{
			Message: "Success",
			Data:    newUser,
		})
	}
	return c.JSON(http.StatusBadRequest, helpers.FailResponse{
		Message: "Failed, NIK sudah terdaftar.",
	})
}

func GetUserController(c echo.Context) error {
	var user []models.User
	// result := config.DB.Preload("loker").Find(&user) //can't preload field loker for models.User 
	result := config.DB.Preload("Loker").Find(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, helpers.FailResponse{
			Message: "Terjadi kesalahan server, Gagal mendapatkan data users.",
		})
	}
	return c.JSON(http.StatusOK, helpers.BaseResponse{
		Message: "Success",
		Data:    user,
	})
}