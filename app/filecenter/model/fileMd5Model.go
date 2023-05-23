package model

import (
	"gorm.io/gorm"
)

var _ FileMd5Model = (*customFileMd5Model)(nil)

type (
	// FileMd5Model is an interface to be customized, add more methods here,
	// and implement the added methods in customFileMd5Model.
	FileMd5Model interface {
		fileMd5Model
	}

	customFileMd5Model struct {
		*defaultFileMd5Model
	}
)

// NewFileMd5Model returns a model for the database table.
func NewFileMd5Model(conn *gorm.DB) FileMd5Model {
	return &customFileMd5Model{
		defaultFileMd5Model: newFileMd5Model(conn),
	}
}
