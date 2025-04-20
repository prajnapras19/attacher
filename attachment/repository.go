package attachment

import (
	"github.com/prajnapras19/attacher/config"
	"gorm.io/gorm"
)

type Repository interface {
	GetAllActiveAttachmentsByUserID(userID uint) ([]*Attachment, error)
	GetActiveAttachmentByUserIDAndSerial(userID uint, serial string) (*Attachment, error)
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

func (r *repository) GetAllActiveAttachmentsByUserID(userID uint) ([]*Attachment, error) {
	var res []*Attachment
	err := r.db.Where("user_id = ? AND is_active", userID).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *repository) GetActiveAttachmentByUserIDAndSerial(userID uint, serial string) (*Attachment, error) {
	var res Attachment
	err := r.db.Where("user_id = ? AND serial = ? AND is_active", userID, serial).First(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}
