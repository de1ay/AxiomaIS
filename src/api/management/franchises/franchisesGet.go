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
	if user.Unit.Type != 0 {
		return conf.ERROR_ACCESS_400
	}
	var franchises []src.Unit
	dbErr := src.Connection.Connection.Where("type=?", 1).Find(&franchises).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	for i, _ := range franchises {
		dbErr = src.Connection.Connection.Model(&franchises[i]).Related(&franchises[i].Users).Error; if dbErr != nil {
			return conf.ERROR_DATABASE_REQUEST_INVALID_101
		}
		for k, _ := range franchises[i].Users {
			franchises[i].Users[k].Unit = src.Unit{
				ID: franchises[i].ID,
				Name: franchises[i].Name,
				Type: franchises[i].Type,
				Parent: franchises[i].Parent,
			}
		}
	}
	return &conf.ApiResponse{200, "success", franchises}
}