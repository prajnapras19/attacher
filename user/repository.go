package user

import (
	"github.com/prajnapras19/attacher/config"
	"gorm.io/gorm"
)

type Repository interface {
	GetUserByUsername(username string) (*User, error)
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

func (r *repository) GetUserByUsername(username string) (*User, error) {
	var res User
	err := r.db.Model(&User{}).Where("username = ?", username).First(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}
