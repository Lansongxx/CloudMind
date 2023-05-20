package model

import (
	"gorm.io/gorm"
)

var _ UserFileModel = (*customUserFileModel)(nil)

type (
	// UserFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserFileModel.
	UserFileModel interface {
		userFileModel
	}

	customUserFileModel struct {
		*defaultUserFileModel
	}
)

// NewUserFileModel returns a model for the database table.
func NewUserFileModel(conn *gorm.DB) UserFileModel {
	return &customUserFileModel{
		defaultUserFileModel: newUserFileModel(conn),
	}
}
