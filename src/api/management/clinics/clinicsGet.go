package clinics

import (
	"axioma/conf"
	"axioma/src/api/authorization"
	"axioma/src"
)

func GetClinics(token string) *conf.ApiResponse {
	user, err := authorization.ValidateToken(token); if err != nil {
		return err
	}
	if user.Role == 7 || user.Role == 8 {
		return conf.ERROR_ACCESS_400
	}
	var clinics []src.Clinic
	if user.Role == 0 || user.Role == 1 || user.Role == 6 {
		dbErr := src.Connection.Connection.Find(&clinics).Error; if dbErr != nil {
			return conf.ERROR_DATABASE_REQUEST_INVALID_101
		}
	} else {
		dbErr := src.Connection.Connection.Where("franchise_id=?", user.FranchiseID).Find(&clinics).Error
		if dbErr != nil {
			return conf.ERROR_DATABASE_REQUEST_INVALID_101
		}
	}
	for i, _ := range clinics {
		dbErr := src.Connection.Connection.Model(&clinics[i]).Related(&clinics[i].Users).Error; if dbErr != nil {
			return conf.ERROR_DATABASE_REQUEST_INVALID_101
		}
		dbErr = src.Connection.Connection.Model(&clinics[i]).Related(&clinics[i].Franchise).Error; if dbErr != nil {
			return conf.ERROR_DATABASE_REQUEST_INVALID_101
		}
		for k, _ := range clinics[i].Users {
			clinics[i].Users[k].Clinic = nil
		}
	}
	return &conf.ApiResponse{200, "success", clinics}
}