package handler

import (
	"github.com/Ribeiro/gopportunities/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger()
	db = config.GetSQLite()
}
