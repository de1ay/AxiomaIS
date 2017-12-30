package franchises

import (
	"axioma/conf"
	"axioma/src/api/authorization"
	"axioma/src"
)

// TODO: unify methods for clinics and franchises

func UpdateFranchise(token string, newFranchise src.Unit) *conf.ApiResponse {
	user, err := authorization.ValidateToken(token); if err != nil {
		return err
	}
	if user.Unit.Type != 0 {
		return conf.ERROR_ACCESS_400
	}
	var franchise src.Unit
	dbErr := src.Connection.Connection.Find(&franchise, newFranchise.ID).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	franchise.Name = newFranchise.Name
	franchise.Type = newFranchise.Type
	franchise.Parent = newFranchise.Parent
	dbErr = src.Connection.Connection.Save(&franchise).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	return conf.REQUEST_SUCCESS_200
}