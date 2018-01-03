package clinics

import (
	"axioma/conf"
	"axioma/src/api/authorization"
	"axioma/src"
)

func DeleteClinic(token string, clinic src.Clinic) *conf.ApiResponse {
	user, err := authorization.ValidateToken(token); if err != nil {
		return err
	}
	if user.Role == 7 || user.Role == 8 || (user.Role != 0 && user.Role != 1 && user.Role != 6 && user.FranchiseID != clinic.FranchiseID ) {
		return conf.ERROR_ACCESS_400
	}
	dbErr := src.Connection.Connection.Where("id=?", clinic.ID).Delete(&src.Clinic{}).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	return conf.REQUEST_SUCCESS_200
}
