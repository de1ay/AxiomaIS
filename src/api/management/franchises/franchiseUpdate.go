package franchises

import (
	"axioma/conf"
	"axioma/src/api/authorization"
	"axioma/src"
)


func UpdateFranchise(token string, newFranchise src.Franchise) *conf.ApiResponse {
	user, err := authorization.ValidateToken(token); if err != nil {
		return err
	}
	if (user.Role != 0 && user.Role != 1 && user.Role != 6) &&
		(user.Franchise.ID != newFranchise.ID && user.Role != 2 && user.Role != 3 && user.Role != 4 && user.Role != 5){
		return conf.ERROR_ACCESS_400
	}
	var franchise src.Franchise
	dbErr := src.Connection.Connection.Find(&franchise, newFranchise.ID).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	franchise.Name = newFranchise.Name
	franchise.OfficialName = newFranchise.OfficialName
	franchise.FullOfficialName = newFranchise.FullOfficialName
	franchise.License = newFranchise.License
	franchise.City = newFranchise.City
	franchise.LegalAddress = newFranchise.LegalAddress
	franchise.Inn = newFranchise.Inn
	franchise.Kpp = newFranchise.Kpp
	franchise.BankAccount = newFranchise.BankAccount
	franchise.Bank = newFranchise.Bank
	franchise.CorrespondentAccount = newFranchise.CorrespondentAccount
	franchise.Bik = newFranchise.Bik
	franchise.Ogrn = newFranchise.Ogrn
	franchise.Okved = newFranchise.Okved
	franchise.Act = newFranchise.Act
	franchise.CorrespondentAddress = newFranchise.CorrespondentAddress
	franchise.Number = newFranchise.Number
	dbErr = src.Connection.Connection.Save(&franchise).Error; if dbErr != nil {
		return conf.ERROR_DATABASE_REQUEST_INVALID_101
	}
	return conf.REQUEST_SUCCESS_200
}