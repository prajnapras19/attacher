package user

import (
	"errors"

	"github.com/prajnapras19/attacher/config"
	"gorm.io/gorm"
)

type Repository interface {
	GetUserByUsername(username string) (*User, error)
	UpsertUser(user User) error
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

func (r *repository) UpsertUser(user User) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&user).Where("id = ?", user.ID).First(&User{}).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return tx.Create(&user).Error
			}
			return err
		}
		return tx.Model(&user).Where("id = ?", user.ID).Updates(&user).Error
	})
}
