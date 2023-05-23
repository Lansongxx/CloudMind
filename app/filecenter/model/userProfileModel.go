package model

import (
	"gorm.io/gorm"
)

var _ UserProfileModel = (*customUserProfileModel)(nil)

type (
	// UserProfileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserProfileModel.
	UserProfileModel interface {
		userProfileModel
	}

	customUserProfileModel struct {
		*defaultUserProfileModel
	}
)

// NewUserProfileModel returns a model for the database table.
func NewUserProfileModel(conn *gorm.DB) UserProfileModel {
	return &customUserProfileModel{
		defaultUserProfileModel: newUserProfileModel(conn),
	}
}
