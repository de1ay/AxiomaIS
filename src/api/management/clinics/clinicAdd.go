package clinics

import (
	"axioma/src"
	"axioma/conf"
	"axioma/src/api/authorization"
)

// TODO: check parent unit

func AddClinic(token string, clinic src.Unit) *conf.ApiResponse {
	user, err := authorization.ValidateToken(token); if err != nil {
		return err
	}
	if user.Unit.Type != 0 {
		return conf.ERROR_ACCESS_400
	}
	dbErr := src.Connection.Connection.Create(&clinic).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	return &conf.ApiResponse{200, "success", &clinic}
}
