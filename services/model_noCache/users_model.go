package model_noCache

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		customUsersLogicModel
	}

	customUsersModel struct {
		*defaultUsersModel
	}

	customUsersLogicModel interface {
	}
)

// BeforeCreate hook create time
func (s *Users) BeforeCreate(tx *gorm.DB) error {
	now := time.Now()
	s.CreatedAt = sql.NullTime{Time: now, Valid: true}
	s.UpdatedAt = sql.NullTime{Time: now, Valid: true}
	return nil
}

// BeforeUpdate hook update time
func (s *Users) BeforeUpdate(tx *gorm.DB) error {
	s.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	return nil
}

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn *gorm.DB) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn),
	}
}
