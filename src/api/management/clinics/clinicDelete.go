package clinics

import (
	"axioma/conf"
	"axioma/src/api/authorization"
	"axioma/src"
)

func DeleteClinic(token string, id uint) *conf.ApiResponse {
	user, err := authorization.ValidateToken(token); if err != nil {
		return err
	}
	if user.Unit.Type != 0 {
		return conf.ERROR_ACCESS_400
	}
	dbErr := src.Connection.Connection.Where("id=? AND type=?", id, 2).Delete(&src.Unit{}).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	return conf.REQUEST_SUCCESS_200
}
