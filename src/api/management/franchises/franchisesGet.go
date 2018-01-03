package franchises

import (
	"axioma/conf"
	"axioma/src/api/authorization"
	"axioma/src"
)

func GetFranchises(token string) *conf.ApiResponse {
	user, err := authorization.ValidateToken(token); if err != nil {
		return err
	}
	if user.Role != 0 && user.Role != 1 && user.Role != 6 {
		return conf.ERROR_ACCESS_400
	}
	var franchises []src.Franchise
	dbErr := src.Connection.Connection.Find(&franchises).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	for i, _ := range franchises {
		dbErr = src.Connection.Connection.Model(&franchises[i]).Related(&franchises[i].Users).Error; if dbErr != nil {
			return conf.ERROR_DATABASE_REQUEST_INVALID_101
		}
		for k, _ := range franchises[i].Users {
			franchises[i].Users[k].Franchise = nil
		}
	}
	return &conf.ApiResponse{200, "success", franchises}
}