package clinics

import (
	"axioma/src"
	"axioma/conf"
	"axioma/src/api/authorization"
)

func AddClinic(token string, clinic src.Clinic) *conf.ApiResponse {
	user, err := authorization.ValidateToken(token); if err != nil {
		return err
	}
	if user.Role == 7 || user.Role == 8 || (user.Role != 0 && user.Role != 1 && user.Role != 6 && user.FranchiseID != clinic.FranchiseID ) {
		return conf.ERROR_ACCESS_400
	}
	dbErr := src.Connection.Connection.Create(&clinic).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	return &conf.ApiResponse{200, "success", &clinic}
}
