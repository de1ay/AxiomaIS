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
	if user.Unit.Type == 2 {
		return conf.ERROR_ACCESS_400
	}
	var clinics []src.Unit
	if user.Unit.Type == 0 {
		dbErr := src.Connection.Connection.Where("type=?", 2).Find(&clinics).Error;
		if dbErr != nil {
			return conf.ERROR_DATABASE_REQUEST_INVALID_101
		}
	} else {
		dbErr := src.Connection.Connection.Where("type=? AND parent=?", 2, user.Unit.ID).Find(&clinics).Error;
		if dbErr != nil {
			return conf.ERROR_DATABASE_REQUEST_INVALID_101
		}
	}
	for i, _ := range clinics {
		dbErr := src.Connection.Connection.Model(&clinics[i]).Related(&clinics[i].Users).Error; if dbErr != nil {
			return conf.ERROR_DATABASE_REQUEST_INVALID_101
		}
		for k, _ := range clinics[i].Users {
			clinics[i].Users[k].Unit = src.Unit{
				ID: clinics[i].ID,
				Name: clinics[i].Name,
				Type: clinics[i].Type,
				Parent: clinics[i].Parent,
			}
		}
	}
	return &conf.ApiResponse{200, "success", clinics}
}