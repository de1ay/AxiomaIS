package clinics

import (
	"axioma/conf"
	"axioma/src/api/authorization"
	"axioma/src"
)

func UpdateClinic(token string, newClinic src.Clinic) *conf.ApiResponse {
	user, err := authorization.ValidateToken(token); if err != nil {
		return err
	}
	if user.Role == 7 || user.Role == 8 || (user.Role != 0 && user.Role != 1 && user.Role != 6 && user.FranchiseID != newClinic.FranchiseID ) {
		return conf.ERROR_ACCESS_400
	}
	var clinic src.Clinic
	dbErr := src.Connection.Connection.Find(&clinic, newClinic.ID).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	clinic.ID = newClinic.ID
	clinic.Name = newClinic.Name
	clinic.City = newClinic.City
	clinic.Address = newClinic.Address
	clinic.IsActive = newClinic.IsActive
	dbErr = src.Connection.Connection.Save(&clinic).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	return conf.REQUEST_SUCCESS_200
}