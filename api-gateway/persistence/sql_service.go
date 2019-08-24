package persistence

import (
	"api-gateway/model"
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

func (ss *SQLService) GetUser(userId string) *model.User {
	user := model.User{}
	db := ss.db.First(&user, "user_id=?", userId)
	if db.Error != nil {
		return nil
	}
	return &user
}

func (ss *SQLService) SetUser(user *model.User) error {
	db := ss.db.Create(user)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
