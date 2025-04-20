package attachment

import (
	"github.com/prajnapras19/attacher/config"
	"gorm.io/gorm"
)

type Repository interface {
}

type repository struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewRepository(
	cfg *config.Config,
	db *gorm.DB,
) Repository {
	return &repository{
		cfg: cfg,
		db:  db,
	}
}
