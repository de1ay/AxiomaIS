package clinics

import (
	"axioma/conf"
	"axioma/src/api/authorization"
	"axioma/src"
)

func UpdateClinic(token string, newClinic src.Unit) *conf.ApiResponse {
	user, err := authorization.ValidateToken(token); if err != nil {
		return err
	}
	if user.Unit.Type != 0 {
		return conf.ERROR_ACCESS_400
	}
	var clinic src.Unit
	dbErr := src.Connection.Connection.Find(&clinic, newClinic.ID).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	clinic.Name = newClinic.Name
	clinic.Type = newClinic.Type
	clinic.Parent = newClinic.Parent
	dbErr = src.Connection.Connection.Save(&clinic).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	return conf.REQUEST_SUCCESS_200
}