package storage

import (
	"github.com/zxc7563598/key-heat/internal/model"
	"gorm.io/gorm"
)

func Run(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.KeyEvent{},
	)
}
