package model

import (
	"gorm.io/gorm"
)

var _ ProfilePictureModel = (*customProfilePictureModel)(nil)

type (
	// ProfilePictureModel is an interface to be customized, add more methods here,
	// and implement the added methods in customProfilePictureModel.
	ProfilePictureModel interface {
		profilePictureModel
	}

	customProfilePictureModel struct {
		*defaultProfilePictureModel
	}
)

// NewProfilePictureModel returns a model for the database table.
func NewProfilePictureModel(conn *gorm.DB) ProfilePictureModel {
	return &customProfilePictureModel{
		defaultProfilePictureModel: newProfilePictureModel(conn),
	}
}
