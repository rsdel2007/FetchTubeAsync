package job

import (
	"github.com/rsdel2007/proj/contract"
	"gorm.io/gorm"
)

type DatabaseInitializer struct {
	DB *gorm.DB
}

func NewDatabaseInitializer(db *gorm.DB) *DatabaseInitializer {
	return &DatabaseInitializer{
		DB: db,
	}
}

func (di *DatabaseInitializer) Initialize() {
	di.DB.AutoMigrate(&contract.Video{})
}
