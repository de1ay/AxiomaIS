package src

import (
	"axioma/conf"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"regexp"
)

var Connection DatabaseConnection

type DatabaseConnection struct {
	Connection *gorm.DB
}

func (dbConnection *DatabaseConnection) getSource() string {
	source := conf.MYSQL_LOGIN + ":"
	source += conf.MYSQL_PASSWORD + "@tcp("
	source += conf.MYSQL_SERVER_ADDR + conf.MYSQL_SERVER_PORT + ")/"
	source += conf.MYSQL_DB_NAME
	source += "?charset=utf8&parseTime=True"
	return source
}

func (dbConnection *DatabaseConnection) Connect() {
	var err error
	source := dbConnection.getSource()
	dbConnection.Connection, err = gorm.Open("mysql", source);
	if err != nil {
		log.Fatal(err)
	}
}

func (dbConnection *DatabaseConnection) Sync() {
	dbConnection.Connection.AutoMigrate(&Franchise{}, &Clinic{}, &User{}, &Session{})
}

type User struct {
	ID          uint       `json:"id";gorm:"primary_key"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	Login       string     `json:"login";gorm:"size:64"`
	Password    string     `json:"-";gorm:"size:255"`
	Surname     string     `json:"surname";gorm:"size:64"`
	Name        string     `json:"name";gorm:"size:64"`
	Patronymic  string     `json:"patronymic";gorm:"size:64"`
	Number      string     `json:"number"`
	Birthday    time.Time  `json:"birthday"`
	Role        int        `json:"role"`
	IsActive    bool       `json:"is_active"`
	Franchise   *Franchise `json:"franchise,omitempty"`
	FranchiseID uint       `json:"-"`
	Clinic      *Clinic    `json:"clinic,omitempty"`
	ClinicID    uint       `json:"-"`
}

type Session struct {
	ID        uint       `json:"id";gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Staff     *User      `json:"users"`
	UserID    uint       `json:"-"`
	Token     string     `json:"token";gorm:"size:255"`
}

type Franchise struct {
	ID                   uint       `json:"id";gorm:"primary_key"`
	CreatedAt            time.Time  `json:"-"`
	UpdatedAt            time.Time  `json:"-"`
	DeletedAt            *time.Time `json:"-"`
	Users                []User     `json:"users,omitempty"`
	Name                 string     `json:"name"`
	OfficialName         string     `json:"official_name"`
	FullOfficialName     string     `json:"full_official_name"`
	License              string     `json:"license";gorm:"size:255"`
	City                 string     `json:"city"`
	LegalAddress         string     `json:"legal_address"`
	Inn                  string     `json:"inn"`
	Kpp                  string     `json:"kpp"`
	BankAccount          string     `json:"bank_account"`
	Bank                 string     `json:"bank"`
	CorrespondentAccount string     `json:"correspondent_account"`
	Bik                  string     `json:"bik"`
	Ogrn                 string     `json:"ogrn"`
	Okved                string     `json:"okved"`
	Principal            User      `json:"principal,omitempty"`
	PrincipalID          uint       `json:"-"`
	Act                  string     `json:"act"`
	CorrespondentAddress string     `json:"correspondent_address"`
	Number               string     `json:"number"`
}

// Using for unmarshal request body
type FranchiseRaw struct {
	ID                   float64    `json:"id"`
	CreatedAt            time.Time  `json:"-"`
	UpdatedAt            time.Time  `json:"-"`
	DeletedAt            *time.Time `json:"-"`
	Users                []User     `json:"-"`
	Name                 string     `json:"name"`
	OfficialName         string     `json:"official_name"`
	FullOfficialName     string     `json:"full_official_name"`
	License              string     `json:"license"`
	City                 string     `json:"city"`
	LegalAddress         string     `json:"legal_address"`
	Inn                  string     `json:"inn"`
	Kpp                  string     `json:"kpp"`
	BankAccount          string     `json:"bank_account"`
	Bank                 string     `json:"bank"`
	CorrespondentAccount string     `json:"correspondent_account"`
	Bik                  string     `json:"bik"`
	Ogrn                 string     `json:"ogrn"`
	Okved                string     `json:"okved"`
	Principal            *User      `json:"-"`
	PrincipalID          uint       `json:"-"`
	Act                  string     `json:"act"`
	CorrespondentAddress string     `json:"correspondent_address"`
	Number               string     `json:"number"`
}

func (rawFranchise *FranchiseRaw) ToFranchise(adding bool) (Franchise, *conf.ApiResponse) {
	var franchise Franchise
	if rawFranchise.ID <= 0 && !adding {
		return franchise, conf.ERROR_DATA_FORMAT_INVALID_500
	}
	if len(rawFranchise.Name) < 1 {
		return franchise, conf.InvalidField("Название")
	}
	if len(rawFranchise.OfficialName) < 1 {
		return franchise, conf.InvalidField("Официальное наименование")
	}
	if len(rawFranchise.FullOfficialName) < 1 {
		return franchise, conf.InvalidField("Полное официальное наименование")
	}
	if len(rawFranchise.License) < 1 {
		return franchise, conf.InvalidField("Лицензия")
	}
	if len(rawFranchise.City) < 1 {
		return franchise, conf.InvalidField("Город")
	}
	if len(rawFranchise.LegalAddress) < 1 {
		return franchise, conf.InvalidField("Юридический адрес")
	}
	if !regexp.MustCompile("[0-9]{10}").MatchString(rawFranchise.Inn) {
		return franchise, conf.InvalidField("ИНН")
	}
	if !regexp.MustCompile("[0-9]{9}").MatchString(rawFranchise.Kpp) {
		return franchise, conf.InvalidField("КПП")
	}
	if !regexp.MustCompile("[0-9]{20}").MatchString(rawFranchise.BankAccount) {
		return franchise, conf.InvalidField("Р/C")
	}
	if len(rawFranchise.Bank) < 1 {
		return franchise, conf.InvalidField("Банк")
	}
	if !regexp.MustCompile("[0-9]{19}").MatchString(rawFranchise.CorrespondentAccount) {
		return franchise, conf.InvalidField("Кор.счёт")
	}
	if !regexp.MustCompile("[0-9]{9}").MatchString(rawFranchise.Bik) {
		return franchise, conf.InvalidField("БИК")
	}
	if !regexp.MustCompile("[0-9]{13}").MatchString(rawFranchise.Ogrn) {
		return franchise, conf.InvalidField("ОГРН")
	}
	if len(rawFranchise.Okved) < 1 {
		return franchise, conf.InvalidField("ОКВЭД")
	}
	if len(rawFranchise.Act) < 1 {
		return franchise, conf.InvalidField("Действует")
	}
	if len(rawFranchise.CorrespondentAddress) < 1 {
		return franchise, conf.InvalidField("Адрес для корреспонденции")
	}
	if !regexp.MustCompile("[0-9]{10}").MatchString(rawFranchise.Number) {
		return franchise, conf.InvalidField("Телефон")
	}
	franchise.ID = uint(int(rawFranchise.ID))
	franchise.Name = rawFranchise.Name
	franchise.OfficialName = rawFranchise.OfficialName
	franchise.FullOfficialName = rawFranchise.FullOfficialName
	franchise.License = rawFranchise.License
	franchise.City = rawFranchise.City
	franchise.LegalAddress = rawFranchise.LegalAddress
	franchise.Inn = rawFranchise.Inn
	franchise.Kpp = rawFranchise.Kpp
	franchise.BankAccount = rawFranchise.BankAccount
	franchise.Bank = rawFranchise.Bank
	franchise.CorrespondentAccount = rawFranchise.CorrespondentAccount
	franchise.Bik = rawFranchise.Bik
	franchise.Ogrn = rawFranchise.Ogrn
	franchise.Okved = rawFranchise.Okved
	franchise.Act = rawFranchise.Act
	franchise.CorrespondentAddress = rawFranchise.CorrespondentAddress
	franchise.Number = rawFranchise.Number
	return franchise, nil
}

type Clinic struct {
	ID          uint       `json:"id";gorm:"primary_key"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	Name        string     `json:"name"`
	City        string     `json:"city"`
	Address     string     `json:"address"`
	IsActive    bool       `json:"is_active"`
	Franchise   Franchise `json:"franchise,omitempty"`
	FranchiseID uint       `json:"-"`
	Users       []User     `json:"users,omitempty"`
}

// Using for unmarshal request body
type ClinicRaw struct {
	ID          float64    `json:"id"`
	CreatedAt   time.Time  `json:"-"`
	UpdatedAt   time.Time  `json:"-"`
	DeletedAt   *time.Time `json:"-"`
	Name        string     `json:"name"`
	City        string     `json:"city"`
	Address     string     `json:"address"`
	IsActive    bool       `json:"is_active"`
	Franchise   *Franchise `json:"-"`
	FranchiseID float64    `json:"franchise_id"`
	Users       []User     `json:"-"`
}

func (rawClinic *ClinicRaw) ToClinic(adding bool) (Clinic, *conf.ApiResponse) {
	var clinic Clinic
	if rawClinic.ID <= 0 && !adding {
		return clinic, conf.ERROR_DATA_FORMAT_INVALID_500
	}
	if rawClinic.FranchiseID <= 0 {
		return clinic, conf.ERROR_DATA_FORMAT_INVALID_500
	}
	if len(rawClinic.Name) < 1 {
		return clinic, conf.InvalidField("Название")
	}
	if len(rawClinic.City) < 1 {
		return clinic, conf.InvalidField("Город")
	}
	if len(rawClinic.Address) < 1 {
		return clinic, conf.InvalidField("Адрес")
	}
	clinic.ID = uint(int(rawClinic.ID))
	clinic.Name = rawClinic.Name
	clinic.City = rawClinic.City
	clinic.Address = rawClinic.Address
	clinic.IsActive = rawClinic.IsActive
	clinic.FranchiseID = uint(int(rawClinic.FranchiseID))
	return clinic, nil
}
