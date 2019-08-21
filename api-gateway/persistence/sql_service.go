package persistence

import (
	"github.com/jinzhu/gorm"
	"log"
)

type SQLService struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewSQLService(db *gorm.DB, l *log.Logger) *SQLService {
	return &SQLService{
		db:     db,
		logger: l,
	}
}

func (ss *SQLService) GetUser(userId string) *User {
	user := User{}
	db := ss.db.First(&user, "userId=?", userId)
	if db.Error != nil {
		return nil
	}
	return &user
}

func (ss *SQLService) SetUser(user *User) error {
	db := ss.db.Create(user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
