package personshare

import (
	"time"

	"gorm.io/plugin/optimisticlock"
)

type PersonModel struct {
	ID        uint
	Name      string `gorm:"notNull;size:256"`
	Email     string `gorm:"notNull;unique;size:256"`
	Password  string `gorm:"notNull;size:256"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   optimisticlock.Version
}
