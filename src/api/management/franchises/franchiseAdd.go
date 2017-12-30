package franchises

import (
	"axioma/src"
	"axioma/conf"
	"axioma/src/api/authorization"
)

func AddFranchise(token string, franchise src.Unit) *conf.ApiResponse {
	user, err := authorization.ValidateToken(token); if err != nil {
		return err
	}
	if user.Unit.Type != 0 {
		return conf.ERROR_ACCESS_400
	}
	dbErr := src.Connection.Connection.Create(&franchise).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	return &conf.ApiResponse{200, "success", &franchise}
}