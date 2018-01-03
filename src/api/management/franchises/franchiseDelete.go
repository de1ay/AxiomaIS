package franchises

import (
	"axioma/conf"
	"axioma/src/api/authorization"
	"axioma/src"
)

func DeleteFranchise(token string, id uint) *conf.ApiResponse {
	user, err := authorization.ValidateToken(token); if err != nil {
		return err
	}
	if user.Role != 0 && user.Role != 1 && user.Role != 6 && user.Franchise.ID != id {
		return conf.ERROR_ACCESS_400
	}
	dbErr := src.Connection.Connection.Where("id=?", id).Delete(&src.Franchise{}).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	return conf.REQUEST_SUCCESS_200
}
