package src

import (
	"axioma/conf"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
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
	dbConnection.Connection.AutoMigrate(&Unit{}, &User{}, &Session{})
}

type User struct {
	ID         uint       `json:"id";gorm:"primary_key"`
	CreatedAt  time.Time  `json:"-"`
	UpdatedAt  time.Time  `json:"-"`
	DeletedAt  *time.Time `json:"-"`
	Login      string     `json:"login";gorm:"size:64"`
	Password   string     `json:"-";gorm:"size:255"`
	Surname    string     `json:"surname";gorm:"size:64"`
	Name       string     `json:"name";gorm:"size:64"`
	Patronymic string     `json:"patronymic";gorm:"size:64"`
	Role       int        `json:"role"`
	Unit       Unit       `json:"unit,omitempty"`
	UnitID     uint       `json:"-"`
}

type Session struct {
	ID        uint       `json:"id";gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	User      User       `json:"user"`
	UserID    uint       `json:"-"`
	Token     string     `json:"token";gorm:"size:255"`
}

type Unit struct {
	ID        uint       `json:"id";gorm:"primary_key"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	Name      string     `json:"name"`
	Type      int        `json:"type"`
	Parent    int        `json:"parent"`
	Users     []User     `json:"users,omitempty";gorm:"ForeignKey:UnitID"`
}
